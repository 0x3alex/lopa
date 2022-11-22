package main

import (
	bot "github.com/0x3alex/lopa/etc"
)

func main() {
	bot.GetBot().ParseConfig("config.json")

	//create spotify instance
	bot.GetBot().Spotify = bot.NewSpotify(
		bot.GetBot().SpotifyID,
		bot.GetBot().SpotifySecret)

	//auth instance
	res := bot.GetBot().Spotify.Auth()

	//register the commands before the bot launch
	if res {
		bot.RegisterCommand(bot.Command{
			Name: "Search",
			Desc: "Search for an artist/album on spotify",
		}, searchCommand)
	}

	bot.RegisterCommand(bot.Command{
		Name: "Help",
		Desc: "Typical help command",
	}, helpCommand)

	bot.RegisterCommand(bot.Command{
		Name: "Pong",
		Desc: "Ping Pong",
	}, pongCommand)

	bot.GetBot().Start()
}
