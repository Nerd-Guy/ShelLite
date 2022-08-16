package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/inancgumus/screen"
)

func InitCommands() {
	NewCommand("cat", "Print contents of file", func(arguments []string) error {
		if len(arguments) == 0 {
			return errors.New("Insufficient arguments")
		}
		file, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			return err
		}
		fmt.Print(string(file))
		fmt.Println()
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
	NewCommand("cls", "clear screen", func(arguments []string) error {
		screen.Clear()
		screen.MoveTopLeft()
		return nil
	})
	NewCommand("dev", "test", func(arguments []string) error {
		fmt.Println("Dev Testing")
		fmt.Println("Arguments~")
		for i, v := range arguments {
			fmt.Println(i, v)
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
			var SizeContent string
			if f.IsDir() {
				SizeContent = "<DIR>"
			} else {
				SizeContent = fmt.Sprintf("%d", f.Size())
			}
			FileTable.AddRow(f.Name(), SizeContent, fmt.Sprint(f.ModTime().String()))
		}
		FileTable.Print()
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
	NewCommand("print", "Print arguments to screen", func(arguments []string) error {
		for _, v := range arguments {
			fmt.Print(v)
		}
		fmt.Print("\n")
		return nil
	})
	NewCommand("restart", "Restart Shellite", func(arguments []string) error {

		os.Exit(0)
		return nil
	})
	NewCommand("var", "Create environment variable", func(arguments []string) error {
		if len(arguments) < 2 {
			return errors.New("Insufficient arguments")
		}
		variables[arguments[0]] = arguments[1]
		return nil
	})
}
