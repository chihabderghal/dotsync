#!/usr/bin/env bash

set -euo pipefail

sudo pacman -S --needed --noconfirm \
    hyprland \
    hyprpaper \
    hyprlock \
    waybar \
    wofi \
    grim \
    slurp \
    swaybg \
    swaylock \
    brightnessctl \
    wl-clipboard \
    rofi \
    thunar \
    dolphin

echo "[ok] Environment packages installed"

if !command vicinae version &> /dev/null; then
    echo "[info] Installing vicinae..."
    
    # Download and run the installation script
    curl -fsSL https://vicinae.com/install.sh | bash

    echo "[ok] Vicinae installed"
fi

echo "[ok] Vicinae launcher is already installed"
