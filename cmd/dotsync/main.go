package main

import (
	"fmt"
	utils "github.com/chihabderghal/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "1.1.0"

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
		utils.PrintHelp()
		return
	}

	switch os.Args[1] {
	case "-h", "--help", "help":
		utils.PrintHelp()
	case "-v", "--version", "version":
		fmt.Println("dotsync version", version)
	case "all":
		scriptPaths := []string{
			"./deps/fonts.sh",
			"./deps/dev.sh",
			"./deps/aur.sh",
			"./deps/env.sh",
			"./deps/dot-setup.sh",
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
	case "dot-setup":
		scriptPath := "./deps/dot-setup.sh"
		if err := runScript(scriptPath); err != nil {
			fmt.Printf("[!] Error executing script: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		utils.PrintHelp()
	}
}
