#!/bin/bash
# scripts/install.sh
# This script installs Caddy and prepares the Linux environment for WCP360.

echo "🌐 Installing WCP360 Dependencies..."

# Install Caddy (Debian/Ubuntu)
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy -y

# Setup directories
sudo mkdir -p /var/www/wcp360
sudo chown -R $USER:$USER /var/www/wcp360

echo "✅ System is ready for WCP360!"
