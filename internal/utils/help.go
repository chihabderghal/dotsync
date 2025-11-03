package utils

import "fmt"

type Command struct {
	Name        string
	Description string
}

type Flag struct {
	Name        string
	Description string
}

var commands = []Command{
	{"all", "Set up the entire system (development environment, fonts, AUR packages)"},
	{"dev", "Set up development environment using the provided Bash script"},
	{"fonts", "Install common fonts using the provided Bash script"},
	{"aur", "Install AUR packages using the provided Bash script"},
	{"env", "Set up environment packages using the provided Bash script"},
	{"dot-setup", "Set up dotfiles and configurations using the provided Bash script"},
}

var flags = []Flag{
	{"help", "Show this help message"},
	{"version", "Show version information"},
}

func PrintHelp() {
	fmt.Println("dotsync - A simple tool to sync your dotfiles and system configs")
	fmt.Print("Usage:\n  dotsync [command] [flags]\n")

	fmt.Println("Commands:")
	for _, cmd := range commands {
		fmt.Printf("  %-10s %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println("\nFlags:")
	for _, flag := range flags {
		fmt.Printf("  %-10s %s\n", flag.Name, flag.Description)
	}

	fmt.Println("\nExamples:")
	fmt.Println("  dotsync fonts")
	fmt.Println("  dotsync --help")
}
