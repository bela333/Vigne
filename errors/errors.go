package errors

import "errors"

var NoConfig = errors.New("couldn't find configuration in the database")
var NoRoles = errors.New("couldn't find role command configuration in the database")
var CreatedConfig = errors.New("couldn't find configuration in the database. Created default one")
var NoModule = errors.New("couldn't find registered module")
var MessageNotSent = errors.New("couldn't replace message, since it doesn't exist")
var NoWelcomer = errors.New("couldn't find welcomer:main, welcomer:text:after or welcomer:text:before in the database")
var InvalidExtractor = errors.New("music can't be played from this site")

//Public errors
var NoCommand = New("", "Couldn't find this command")
var NoMusic = New("couldn't find musicChannel and musicVoiceChannel in the config. Music bot functionality will be unavailable", "The Music Bot is unavailable")
var NotPlaying = New("", "There is nothing playing currently...")
var NotRequester = New("", "You can't skip this music")
var MusicTooLong = New("", "The selected song is too long")
var MusicIsLive = New("", "Sorry, but I can't play live streams")