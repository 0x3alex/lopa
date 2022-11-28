package discord

import "github.com/0x3alex/lopa/apis"

type (
	Bot struct {
		Name          string `json:"name"`
		Token         string `json:"token"`
		Prefix        string `json:"prefix"`
		Status        string `json:"status"`
		SpotifyID     string `json:"spotifyid"`
		SpotifySecret string `json:"spotifysecret"`
		Spotify       *apis.Spotify
	}
	Command struct {
		Name string
		Desc string
	}
)
