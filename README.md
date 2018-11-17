# Vigne
A Discord bot written in Go. Used by the [Gabriel Dropout](http://discord.gg/e2Svd88) server.

## Install
```shell
$ go get https://github.com/bela333/Vigne
```

## Database structure

The **Bold** ones are required.

| name | Hash information | description |
| - | - | - |
**vigne:hasConfig** | | Should always be 1
**vigne:config** | token, commandRegex | *Hash*. Simple configuration
**vigne:mods** | | *Set*. User IDs of moderators
vigne:roles | Key is the lowercase name of the role. Value is the ID of the role. | *Hash*. Used by the --role command
vigne:welcomer:main | | ID of the channel where welcome messages should go
vigne:welcomer:secret | | ID of the channel where messages of leaves and joins should go
vigne:welcomer:text:before | | Message, that should be displayed when a user joins. Example: `Welcome %s! Have fun!`
vigne:welcomer:text:after | | After a short period of time, the original message gets replaced by this one. Example: `Welcome %s!`

## Default commands

Every command starts with a prefix. With the default config there are multiple prefixes. The *official* one is `--`

| Command | Description |
| - | - |
ping | Pong!
role | Gives the user a role according to vigne:roles. A user can only have one role from vigne:roles at a time.
help | Lists all available, not hidden commands.
roles | Hidden. Can only be used by moderators. Lists all roles available on the server and their IDs.
replace | Hidden. Sends `Hello` then after a short while replaces the text with `World!`.

## Roadmap
- [ ] Documentation
- [ ] Errors that are sent to the user
- [x] Menu system with Reactions
- [ ] Music Bot
