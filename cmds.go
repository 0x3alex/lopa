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

func searchArtistCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	//send error
	if len(args) < 2 {
		sendErrorEmbed("Wrong usage!",
			bot.GetBot().Prefix+"SearchArtist <artist>", s, m)
		return
	}
	artist := strings.Join(args[1:], " ")
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
		Image: &discordgo.MessageEmbedImage{
			URL: result.Image,
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func searchAlbumCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	//send error
	if len(args) < 2 {
		sendErrorEmbed("Wrong usage!",
			bot.GetBot().Prefix+"SearchAlbum <artist>", s, m)
		return
	}

	album := strings.Join(args[1:], " ")
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
				Name:   "Total Tracks",
				Value:  result[0].TotalTracks,
				Inline: true,
			},
			{
				Name:   "Released",
				Value:  result[0].Release,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
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
