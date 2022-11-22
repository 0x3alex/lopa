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
	bot.GetBot().Spotify.Auth()

	//register the commands before the bot launch
	bot.RegisterCommand(bot.Command{
		Name: "SearchArtist",
		Desc: "Search for an artist on spotify",
	}, searchArtistCommand)

	bot.RegisterCommand(bot.Command{
		Name: "SearchAlbum",
		Desc: "Search for albums on spotify",
	}, searchAlbumCommand)

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
