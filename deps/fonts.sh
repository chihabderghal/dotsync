#!/usr/bin/env bash

set -euo pipefail

# Check if pacman is available
if ! command -v pacman &> /dev/null; then
    echo "pacman could not be found, skipping font installation."
    exit 0
fi

# Fonts installation
sudo pacman -S --needed --noconfirm \
    ttf-bitstream-vera \
    ttf-dejavu \
    gnu-free-fonts \
    ttf-liberation \
    noto-fonts \
    ttf-roboto \
    ttf-anonymous-pro \
    ttf-cascadia-code \
    noto-fonts-emoji \
    ttf-ubuntu-mono-nerd \
    ttf-nerd-fonts-symbols \
    ttf-nerd-fonts-symbols-common \
    ttf-jetbrains-mono-nerd \
    ttf-jetbrains-mono

echo "[ok] Fonts installed"

# Refresh font cache
fc-cache -fv

echo "[ok] Fonts cache refreshed"
