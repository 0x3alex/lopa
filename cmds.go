package main

import (
	"fmt"
	bot "github.com/0x3alex/lopa/etc"
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

func searchCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	if len(args) < 3 {
		sendErrorEmbed("Wrong usage!",
			bot.GetBot().Prefix+"search <artist|album> <name>", s, m)
		return
	}
	mode := strings.ToLower(args[1])
	if mode != "artist" && mode != "album" {
		sendErrorEmbed("Wrong usage!",
			bot.GetBot().Prefix+"search <artist|album> <name>", s, m)
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
	result := bot.GetBot().Spotify.SearchArtist(artist)
	if result.Name == "" {
		sendErrorEmbed("No artist found!", "", s, m)
		return
	}
	embed := &discordgo.MessageEmbed{
		Color: 0x63dc3c,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(s.State.User.Avatar),
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
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: result.Image,
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func searchAlbumCommand(album string, s *discordgo.Session, m *discordgo.MessageCreate) {
	count := 1
	result := bot.GetBot().Spotify.SearchAlbum(album, count)
	if len(result) < 1 {
		sendErrorEmbed("No album found!", "", s, m)
		return
	}
	embed := &discordgo.MessageEmbed{
		Color: 0x63dc3c,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(s.State.User.Avatar),
		},
		Title:       result[0].AlbumName,
		Description: "- " + result[0].AristName[:len(result[0].AristName)-2],
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Released",
				Value:  result[0].Release,
				Inline: false,
			},
			{
				Name:   "Tracks - " + fmt.Sprintf("%d", result[0].TotalTracks),
				Value:  result[0].Tracks,
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: result[0].Image,
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err.Error())
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
	p := bot.PrintCommands()
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
