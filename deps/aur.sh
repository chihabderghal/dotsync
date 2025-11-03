#!/usr/bin/env bash

set -euo pipefail

if ! command -v paru &> /dev/null; then
    echo "[info] Installing paru..."
    
    # Clone paru AUR repository
    cd /tmp
    git clone https://aur.archlinux.org/paru.git
    cd paru

    # Build and install paru 
    makepkg -si --noconfirm 
    
    echo "[ok] Paru installed"
fi

echo "[ok] Paru is already installed"