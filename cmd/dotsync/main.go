package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "undefined"

func printHelp() {
	fmt.Println(`dotsync - A simple tool to sync your dotfiles and system configs

Usage:
  dotsync [command] [flags]

Commands:
  all 	      Set up the entire system (development environment, fonts, AUR packages)
  dev      	  Set up development environment using the provided Bash script
  fonts       Install common fonts using the provided Bash script
  aur         Install AUR packages using the provided Bash script
  help        Show this help message
  version     Show version information

Examples:
  dotsync fonts
  dotsync --help`)
}

func runScript(scriptPath string) error {
	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		return fmt.Errorf("could not get absolute path: %w", err)
	}

	cmd := exec.Command("bash", absPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("[*] Running %s...\n", absPath)
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "-h", "--help", "help":
		printHelp()
	case "-v", "--version", "version":
		fmt.Println("dotsync version", version)
	case "all":
		scriptPaths := []string{
			"./deps/fonts.sh", 
			"./deps/dev.sh", 
			"./deps/aur.sh",
			"./deps/env.sh",
		}
		for _, scriptPath := range scriptPaths {
			if err := runScript(scriptPath); err != nil {
				fmt.Printf("[!] Error executing script %s: %v\n", scriptPath, err)
				os.Exit(1)
			}
		}
	case "env":
		scriptPath := "./deps/env.sh"
		if err := runScript(scriptPath); err != nil {
			fmt.Printf("[!] Error executing script: %v\n", err)
			os.Exit(1)
		}
	case "fonts":
		scriptPath := "./deps/fonts.sh" 
		if err := runScript(scriptPath); err != nil {
			fmt.Printf("[!] Error executing script: %v\n", err)
			os.Exit(1)
		}
	case "dev":
		scriptPath := "./deps/dev.sh" 
		if err := runScript(scriptPath); err != nil {
			fmt.Printf("[!] Error executing script: %v\n", err)
			os.Exit(1)
		}
	case "aur":
		scriptPath := "./deps/aur.sh"
		if err := runScript(scriptPath); err != nil {
			fmt.Printf("[!] Error executing script: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printHelp()
	}
}
