package database

import (
	"github.com/bela333/Vigne/errors"
	"time"
)

type MusicDatabase struct {
	d *Database
}

func (d *Database) Music() (*MusicDatabase, error) {
	if !d.Redis.HExists(d.Decorate("config"), "musicChannel").Val() {
		return nil, errors.NoMusic
	}
	if !d.Redis.HExists(d.Decorate("config"), "musicVoiceChannel").Val() {
		return nil, errors.NoMusic
	}
	return &MusicDatabase{d}, nil
}

func (d MusicDatabase) GetChannel() string {
	return d.d.Redis.HGet(d.d.Decorate("config"), "musicChannel").Val()
}

func (d MusicDatabase) PopNext() string {
	return d.d.Redis.LPop(d.d.Decorate("musicQueue")).Val()
}

func (d MusicDatabase) AddSong(data []byte) error {
	return d.d.Redis.RPush(d.d.Decorate("musicQueue"), string(data)).Err()
}

func (d MusicDatabase) IsValidExtractor(extractor string) bool {
	if d.d.Redis.Exists(d.d.Decorate("validExtractors")).Val() == 0 {
		//If validExtractors doesn't exist, we accept the extractor anyway
		return true
	}
	return d.d.Redis.SIsMember(d.d.Decorate("validExtractors"), extractor).Val()
}

func (d MusicDatabase) GetVoiceChannel() string {
	return d.d.Redis.HGet(d.d.Decorate("config"), "musicVoiceChannel").Val()
}

func (d MusicDatabase) CanPlay(duration time.Duration) bool {
	if !d.d.Redis.HExists(d.d.Decorate("config"), "maxMusicDuration").Val() {
		//No maxMusicDuration is set
		return true
	}
	max, err := d.d.Redis.HGet(d.d.Decorate("config"), "maxMusicDuration").Int()
	if err != nil {
		//Couldn't get maxMusicDuration
		return true
	}
	if time.Duration(max)*time.Second < duration {
		//music duration is larger than maxMusicDuration. Don't play it.
		return false
	}
	return true

}

func (d MusicDatabase) CanPlayLive() bool {
	if d.d.Redis.Exists(d.d.Decorate("canPlayLive")).Val() == 0{
		return true
	}
	val, err := d.d.Redis.Get(d.d.Decorate("canPlayLive")).Int()
	if err != nil {
		return  true
	}
	if val != 0 {
		return true
	}

	return false
}