package music

import (
	"encoding/json"
	"os/exec"
	"strconv"
)

func GetInfo(url string) (*Music, error) {
	var m Music
	//Prepare Youtube-Dl
	ytdl := exec.Command("youtube-dl", "-J", "-a", "-", "--playlist-end", "1")
	stdin, err := ytdl.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := ytdl.StdoutPipe()
	if err != nil {
		return nil, err
	}
	//Start youtube-dl
	ytdl.Start()
	//Pass url to youtube-dl
	stdin.Write([]byte(url))
	stdin.Close()
	//Start decoding the JSON
	decoder := json.NewDecoder(stdout)
	decoder.Decode(&m)
	stdout.Close()
	ytdl.Wait()
	if m.Type == "playlist" {
		if len(m.Entries) > 0 {
			m = m.Entries[0]
		}
	}
	return &m, nil
}

func FormatTime(duration int) string {
	hours := duration / (60*60)
	duration %= 60*60
	minutes := duration/60
	seconds := duration%60
	return strconv.Itoa(hours) + ":" + strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
}