# Dotsync CLI

`dotsync` is a custom command-line utility to **set up your Arch Linux development environment**, including dotfiles, fonts, AUR packages, and desktop environment tools. It automates repetitive setup tasks to save time and keep your system consistent.

---

## Prerequisites

Before installing `dotsync`, make sure you have:

- **Golang** (for building the CLI)
- **Git** (for cloning the repository)
```bash
sudo pacman -S --needed --noconfirm go
sudo pacman -S --needed --noconfirm git 
```

## Installation
1. Clone the repository:
```bash
cd /tmp/
git clone https://github.com/chihabderghal/dotsync.git
```
2. Ensure Go modules are tidy:
```bash
go mod tidy
```
3. Build the CLI:
```bash
make build-default # to create a binary without a version suffix 
```
- This will produce the `dotsync` binary
4. Move the binary to a directory in your system PATH:
```bash
sudo mv ./bin/dotsync /usr/local/bin/dotsync
```

## Usage
Check if `dotsync` is installed successfully:
```bash
dotsync version
```
Run commands: 
```bash
dotsync fonts       # Install common fonts
dotsync aur         # Install AUR packages
dotsync dev         # Set up development environment
dotsync dot-setup   # Sync dotfiles and configurations
dotsync env         # Install environment and desktop packages
dotsync all         # Run all setup steps
```

**NOTE** : `Currently, this command is designed only for my personal setup. Future versions will support multiple OSes and various configurations.`