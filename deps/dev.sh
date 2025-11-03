#!/usr/bin/env bash
set -euo pipefail

# Check if pacman is available
if ! command -v pacman &> /dev/null; then
    echo "pacman could not be found, skipping installation."
    exit 0
fi

# Install basic utility packages
sudo pacman -S --needed --noconfirm \
    base-devel \
    htop \
    wget \
    curl \
    unzip \
    fzf \
    ranger \
    tmux

echo "[ok] Utility packages installed"

# Install development tools
sudo pacman -S --needed --noconfirm \
    go \
    nodejs \
    npm \
    python \
    python-pip \
    docker \
    docker-buildx \
    docker-compose \
    kubectl \
    ansible \
    neovim \
    firefox \
    github-cli \
    telegram-desktop \
    vlc \
    copyq

echo "[ok] Development packages installed"

# Start and enable Docker service
sudo systemctl enable --now docker

# Add current user to the Docker group
sudo usermod -aG docker "$USER"

# Install Bun (if not already installed)
if ! command -v bun &> /dev/null; then
    curl -fsSL https://bun.sh/install | bash
    export PATH="$HOME/.bun/bin:$PATH"
fi

echo "[ok] Bun installed"
echo "[ok] Setup complete!"