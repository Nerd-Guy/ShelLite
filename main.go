package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

var variables map[string]string = make(map[string]string)

func LoadEnv() {}

func InitCommands() {
	NewCommand("dev", "test", func(arguments []string) error {
		fmt.Println("Dev Testing")
		fmt.Println("Arguments~")
		for i, v := range arguments {
			fmt.Println(i, v)
		}
		return nil
	})
	NewCommand("help", "get help on shellite", func(arguments []string) error {
		fmt.Println("Available commands:")

		HelpTable := NewTable()
		HelpTable.AddColumn("Name")
		HelpTable.AddColumn("Description")

		for _, v := range commands {
			HelpTable.AddRow(v.Name, v.Desc)
		}

		HelpTable.Print()
		return nil
	})
	NewCommand("cd", "change directory", func(arguments []string) error {
		if len(arguments) == 0 {
			return nil
		}
		err := os.Chdir(arguments[0])
		if err != nil {
			return err
		}
		return nil
	})
	NewCommand("dir", "list directory contents", func(arguments []string) error {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			return err
		}
		FileTable := NewTable()
		FileTable.AddColumn("Name")
		FileTable.AddColumn("Size (bytes)")
		FileTable.AddColumn("Date Modified")
		for _, f := range files {
			FileTable.AddRow(f.Name(), fmt.Sprint(f.Size()), fmt.Sprint(f.ModTime().String()))
		}
		FileTable.Print()
		return nil
	})
	NewCommand("cls", "clear screen", func(arguments []string) error {
		screen.Clear()
		screen.MoveTopLeft()
		return nil
	})
	NewCommand("var", "Create environment variable", func(arguments []string) error {
		if len(arguments) < 2 {
			return errors.New("Insufficient arguments")
		}
		variables[arguments[0]] = arguments[1]
		return nil
	})
	NewCommand("print", "Print arguments to screen", func(arguments []string) error {
		for _, v := range arguments {
			fmt.Print(v)
		}
		fmt.Print("\n")
		return nil
	})
}

type Config struct {
	StartupPath string
	DebugMode   bool
}

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("Shellite is only supported on Windows")
		fmt.Println("Press any key to exit...")
		keyboard.GetSingleKey()
		return
	}
	InitCommands()
	if len(os.Args) > 1 {
		ExecProg(strings.Join(os.Args[1:], " "))
		os.Exit(0)
	}
	startup := "Shellite REPL.\nUse the 'help' or 'docs' command if you are stuck!\n"
	fmt.Print(startup)
	err := RunRepl()
	if err != nil {
		fmt.Println("Shellite Runtime Error!")
		fmt.Println(strings.Repeat("-", len(err.Error())))
		fmt.Println(err)
		fmt.Println("Press any key to exit...")
		keyboard.GetSingleKey()
		os.Exit(1)
	}
}

func RunRepl() error {
	for {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Printf("%s$ ", cwd)
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)
		err = HandleLine(line)
		if err != nil {
			return err
		}
	}
}

func HandleLine(line string) error {
	if line == "fail" {
		return errors.New("Test Error")
	}
	if line == "" {
		return nil
	}
	tokens := strings.Split(line, " ")
	tokens = ReplaceWithValues(tokens)
	cmd := tokens[0]
	var args []string = []string{}
	if len(tokens) > 1 {
		args = strings.Split(strings.Join(tokens[1:], " "), ",")
	}
	Eval(cmd, args)
	return nil
}

func ReplaceWithValues(tokens []string) []string {
	// Replace variables with values
	new_tokens := tokens
	for i, v := range tokens {
		if strings.HasPrefix(v, "$") {
			new_tokens[i] = variables[v]
		}
	}
	return new_tokens
}

func CommandLookup(command string) (*Command, error) {
	for _, v := range commands {
		if v.Name == command {
			return &v, nil
		}
	}
	return nil, errors.New(
		fmt.Sprintf("Command not found: '%s'", command),
	)
}

func Eval(command string, arguments []string) {
	cmd, err := CommandLookup(command)
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Func(arguments)
}

func ExecProg(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
	}
	contents := string(data)
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		err = HandleLine(line)
		if err != nil {
			fmt.Println(err)
		}
	}
}
