package main

import (
	"fmt"
	"github.com/0x3alex/lopa/apis"
	"github.com/0x3alex/lopa/discord"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func sendErrorEmbed(title, desc string, s *discordgo.Session,
	m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Color: 0xee421f,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(s.State.User.Avatar),
		},
		Title:       title,
		Description: desc,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}

func urbanCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	word := ""
	args := strings.Split(m.Content, " ")
	if len(args) > 1 {
		word = strings.Join(args[1:], " ")
	}
	result := apis.UrbanGetRandom(word)
	votingStr := fmt.Sprintf("%d :thumbsup: and %d :thumbsdown:",
		result.ThumbsUp, result.ThumbsDown)
	embed := &discordgo.MessageEmbed{
		Color: 0x3cbcdc,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Urban Random Word",
			IconURL: "https://slack-files2.s3-us-west-2.amazonaws.com/" +
				"avatars/2018-01-11/297387706245_85899a44216ce1604c93_512.jpg",
		},
		Title:       result.Word + " by " + result.Author,
		Description: result.Link,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Definition",
				Value:  result.Definition,
				Inline: false,
			},
			{
				Name:   "Written On",
				Value:  result.Date,
				Inline: true,
			},
			{
				Name:   "Voting",
				Value:  votingStr,
				Inline: true,
			},
			{
				Name:   "Example",
				Value:  result.Example,
				Inline: false,
			},
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		sendErrorEmbed("Oops!",
			"An Error occurred while processing the Spotify Web API", s, m)
	}
}

func searchCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	if len(args) < 3 {
		sendErrorEmbed("Wrong usage!",
			discord.GetBot().Prefix+"search <artist|album> <name>", s, m)
		return
	}
	mode := strings.ToLower(args[1])
	if mode != "artist" && mode != "album" {
		sendErrorEmbed("Wrong usage!",
			discord.GetBot().Prefix+"search <artist|album> <name>", s, m)
		return
	}
	switch mode {
	case "artist":
		searchArtistCommand(strings.Join(args[2:], " "), s, m)
	case "album":
		searchAlbumCommand(strings.Join(args[2:], " "), s, m)
	}
}

func searchArtistCommand(artist string, s *discordgo.Session, m *discordgo.MessageCreate) {
	result := discord.GetBot().Spotify.SearchArtist(artist)
	if result.Name == "" {
		sendErrorEmbed("No artist found!", "", s, m)
		return
	}

	embed := &discordgo.MessageEmbed{
		Color: 0x63dc3c,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Spotify Lookup",
			IconURL: "https://cdn-icons-png.flaticon.com/512/174/174872.png",
		},
		Title:       result.Name,
		Description: fmt.Sprintf("%d", result.Follower) + " Followers",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Genres",
				Value:  result.Genres,
				Inline: true,
			},
			{
				Name:   "Popularity",
				Value:  fmt.Sprintf("%d", result.Popularity),
				Inline: true,
			},
			{
				Name:   "Albums",
				Value:  result.Albums,
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: result.Image,
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		sendErrorEmbed("Oops!",
			"An Error occurred while processing the Spotify Web API", s, m)
	}
}

func searchAlbumCommand(album string, s *discordgo.Session, m *discordgo.MessageCreate) {
	count := 1 //the number of results, one is the best
	result := discord.GetBot().Spotify.SearchAlbum(album, count)
	if len(result) < 1 {
		sendErrorEmbed("No album found!", "", s, m)
		return
	}
	for _, e := range result {
		embed := &discordgo.MessageEmbed{
			Color: 0x63dc3c,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Spotify Lookup",
				IconURL: "https://cdn-icons-png.flaticon.com/512/174/174872.png",
			},
			Title:       e.AlbumName,
			Description: "- " + e.AristName[:len(e.AristName)-2],
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Released",
					Value:  e.Release,
					Inline: false,
				},
				{
					Name:   "Tracks - " + fmt.Sprintf("%d", result[0].TotalTracks),
					Value:  e.Tracks,
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: e.Image,
			},
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			sendErrorEmbed("Oops!",
				"An Error occurred while processing the Spotify Web API", s, m)
		}
	}
}

func pongCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(s.State.User.Avatar),
		},
		Description: "Ping the pong!",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.ytimg.com/vi/CABqlL02I28/maxresdefault.jpg",
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	p := discord.PrintCommands()
	embedFields := make([]*discordgo.MessageEmbedField, 0)
	for _, e := range p {
		args := strings.Split(e, ";")
		f := discordgo.MessageEmbedField{
			Name:   args[0],
			Value:  args[1],
			Inline: false,
		}
		embedFields = append(embedFields, &f)
	}
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(s.State.User.Avatar),
		},
		Description: "Help is on the way!",
		Fields:      embedFields,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
