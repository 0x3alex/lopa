package etc

import "github.com/bwmarrin/discordgo"

type tCommands map[Command]func(*discordgo.Session, *discordgo.MessageCreate)

var Commands = make(tCommands)

func RegisterCommand(cmd Command,
	//typical function signature for bot commands
	fn func(*discordgo.Session, *discordgo.MessageCreate)) {
	Commands[cmd] = fn
}

func PrintCommands() []string {
	result := make([]string, 0)
	for i := range Commands {
		result = append(result, GetBot().Prefix+i.Name+";"+i.Desc)
	}
	return result
}
