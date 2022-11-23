package main

import (
	"github.com/0x3alex/lopa/apis"
	"github.com/0x3alex/lopa/discord"
)

func main() {
	discord.GetBot().ParseConfig("config.json")

	//create spotify instance
	discord.GetBot().Spotify = apis.NewSpotify(
		discord.GetBot().SpotifyID,
		discord.GetBot().SpotifySecret)

	//auth instance
	res := discord.GetBot().Spotify.Auth()
	//register the commands before the bot launch
	if res {
		discord.RegisterCommand(discord.Command{
			Name: "Search",
			Desc: "Search for an artist/album on spotify",
		}, searchCommand)
	}

	discord.RegisterCommand(discord.Command{
		Name: "Urban",
		Desc: "Get a random word from the urban dictionary\n" +
			"or search for one by using **" + discord.GetBot().Prefix + "Urban <term>**",
	}, urbanCommand)

	discord.RegisterCommand(discord.Command{
		Name: "Help",
		Desc: "Typical help command",
	}, helpCommand)

	discord.RegisterCommand(discord.Command{
		Name: "Pong",
		Desc: "Ping Pong",
	}, pongCommand)

	discord.GetBot().Start()
}
