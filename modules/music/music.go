package music

import (
	"encoding/json"
	"fmt"
	"github.com/bela333/Vigne/errors"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"io"
	"os/exec"
	"time"
)

type MusicPlayer struct {
	server    *server.Server
	IsPlaying bool
	skip      bool
	CurrentlyPlaying *Music
}

//TODO: It might be better if this is moved into another thread that gets notified via a channel
func (p *MusicPlayer) Pump() {
	if p.IsPlaying {
		return
	}
	p.IsPlaying = true
	//Get voice channel
	d, err := p.server.Database.Music()
	if err != nil {
		return
	}
	channel, err := p.server.Session.State.Channel(d.GetVoiceChannel())
	if err != nil {
		return
	}
	//Join voice channel
	voice, err := p.server.Session.ChannelVoiceJoin(channel.GuildID, channel.ID, false, true)
	if err != nil {
		return
	}
	//Get first song
	url, err := p.PopNext()
	if err != nil {
		url = nil
	}
	//Get next songs until we run out of them
	for url != nil {
		p.play(url, voice)
		url, err = p.PopNext()
		if err != nil {
			url = nil
		}
	}
	p.IsPlaying = false
	//No more songs in the queue. Disconnect from voice chat
	voice.Disconnect()

}

func (p *MusicPlayer) Skip()  {
	p.skip = true
}

func (p *MusicPlayer) play(m *Music, voice *discordgo.VoiceConnection) {
	p.CurrentlyPlaying = m
	//Send message in music channel
	//Get music channel
	musicDb, err := p.server.Database.Music()
	if err == nil {
		//Get messaging module
		msgi, err := p.server.GetModuleByName("messages")
		if err == nil {
			msg := msgi.(*messages.MessagesModule)
			//Create new message creator
			creator := msg.NewMessageCreator(musicDb.GetChannel())
			//Create new message
			message := creator.NewMessage()
			//Set content of message to mention of requester
			message.SetContent(fmt.Sprintf("<@%s>", m.RequesterID))
			embed := message.GetEmbedBuilder()
			//Get requester
			user, err := p.server.Session.User(m.RequesterID)
			if err == nil {
				//Use the same embed as the one we used for adding to the queue
				EmbedGenerator(embed, m, user, "Now playing", 0xfed330)
			}

			creator.Send()
		}
	}
	//1. Set up Youtube-DL
	//2. Set up FFmpeg
	//3. Set up thread to copy output of ytdl into ffmpeg
	//4. Start both
	//5. Copy song URL to ytdl STDIN
	//6. Copy STDOUT of FFmpeg and channel it to discordgo
	//7. After copying is done, wait for both processes
	ytdl := exec.Command("youtube-dl", "-o", "-", "-a", "-")

	//2 channel opus with a rate of 48000 and constant bitrate
	//We are using the lowest quality settings
	ffmpeg := exec.Command("ffmpeg", "-i", "-", "-f", "s8", "-ar", "48000", "-ac", "2", "-c:a", "libopus", "-vbr", "off", "-compression_level", "0", "-application", "voip", "-packet_loss", "25", "-") //Must be these values
	//Get stdio
	ytdlIn, err := ytdl.StdinPipe()
	if err != nil {
		return
	}
	ytdlOut, err := ytdl.StdoutPipe()
	if err != nil {
		return
	}
	ffmpegIn, err := ffmpeg.StdinPipe()
	if err != nil {
		return
	}
	ffmpegOut, err := ffmpeg.StdoutPipe()
	if err != nil {
		return
	}
	//Copy from ytdl to ffmpeg
	go func() {
		io.Copy(ffmpegIn, ytdlOut)
		//Copying has finished. This usually means that a pipe have been closed.
		ytdlOut.Close()
		ffmpegIn.Close()
		ffmpegOut.Close()

	}()
	//Start both
	err = ffmpeg.Start()
	if err != nil {
		return
	}
	err = ytdl.Start()
	if err != nil {
		return
	}
	ytdlIn.Write([]byte(m.URL))
	ytdlIn.Close()
	//Copy opus data until skipped
	opusData := make([]byte, 240) //240 is the framesize for Discordgo. It should probably be increased somehow
	p.skip = false
	for !p.skip {
		_, err = io.ReadAtLeast(ffmpegOut, opusData, 240)
		if err != nil {
			break
		}
		voice.OpusSend <- opusData
	}
	//io.ReadAtLeast returned an error. This usually means that the stream has finished.
	//Kill both processes
	ytdl.Process.Kill()
	ffmpeg.Process.Kill()
	//Wait for them so they won't become zombie processes
	ytdl.Wait()
	ffmpeg.Wait()
}

func (p *MusicPlayer) AddMusic(url string, user *discordgo.User) (*Music, error) {
	//Get info from music
	m, err := GetInfo(url)
	if err != nil {
		return nil, err
	}
	m.RequesterName = user.Username
	m.RequesterID = user.ID
	//Check if the extractor is correct
	if !p.IsValidExtractor(m) {
		return nil, errors.InvalidExtractor
	}
	//Add music database object
	d, err := p.server.Database.Music()
	if err != nil {
		return nil, errors.NoMusic
	}
	//Check if the song is too long
	if !d.CanPlay(time.Duration(m.Duration) * time.Second) {
		return m, errors.MusicTooLong
	}
	//Make sure that it isn't live
	if m.IsLive && !d.CanPlayLive() {
		return m, errors.MusicIsLive
	}
	//Convert music data to JSON
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	//And song to queue
	err = d.AddSong(jsonData)
	if err != nil {
		return nil, err
	}
	//Start pumping
	go p.Pump()
	return m, nil
}

func (p *MusicPlayer) IsValidExtractor(m *Music) bool {
	d, err := p.server.Database.Music()
	if err != nil {
		return false
	}
	return d.IsValidExtractor(m.Extractor)
}

func (p *MusicPlayer) PopNext() (*Music, error) {
	d, err := p.server.Database.Music()
	if err != nil {
		return nil, err
	}
	jsonData := d.PopNext()
	var m Music
	err = json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}