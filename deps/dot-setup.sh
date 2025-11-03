#!/usr/bin/env bash
set -euo pipefail

REPO_URL="https://github.com/chihabderghal/dotfiles.git"
DOTFILES_DIR="$HOME/.dotfiles"
CONFIG_DIR="$HOME/.config"

log() { echo -e "[\e[32mINFO\e[0m] $*"; }
warn() { echo -e "[\e[33mWARN\e[0m] $*"; }

if ! command -v git &> /dev/null; then
    warn "git could not be found. Please install git first."
    exit 1
fi

if [ -d "$DOTFILES_DIR/.git" ]; then
    log "Dotfiles repository already exists. Pulling latest changes..."
    git -C "$DOTFILES_DIR" pull --rebase
else
    log "Cloning dotfiles repository..."
    git clone "$REPO_URL" "$DOTFILES_DIR"
fi

for dir in hypr i3 i3status tmux alacritty waybar; do
    src="$DOTFILES_DIR/$dir"
    dest="$CONFIG_DIR/$dir"
    if [ -d "$src" ]; then
        log "Syncing $dir..."
        rsync -a --delete "$src/" "$dest/"
    else
        warn "Missing $src, skipping."
    fi
done

cat "$DOTFILES_DIR/bash/dot-bashrc" >> "$HOME/.bashrc"
cat "$DOTFILES_DIR/bash/dot-bash_aliases" >> "$HOME/.bash_aliases"

log "[ok] Dotfiles setup complete."
