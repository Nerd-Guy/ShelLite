package shellite

type Command struct {
	Name string
	Desc string
	Func func(arguments []string) error
}

var commands = []Command{}

func NewCommand(Name string, Desc string, Func func(arguments []string) error) *Command {
	var com Command = Command{Name, Desc, Func}
	commands = append(commands, com)
	return &com
}
