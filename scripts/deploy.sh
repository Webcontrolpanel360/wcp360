#!/bin/bash
# scripts/deploy.sh - WCP360 Auto-Installer
# Use: sudo bash scripts/deploy.sh

echo "🚀 Starting WCP360 Deployment..."

# Update system
sudo apt update && sudo apt upgrade -y

# Install Core Dependencies
sudo apt install -y curl git tar build-essential sqlite3 libsqlite3-dev

# Install Caddy (Modern Web Server)
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update && sudo apt install caddy -y

# Install PHP Multi-versions (Common for Web Hosting)
sudo apt install -y php8.1-fpm php8.2-fpm php8.3-fpm

# Build WCP360 Binary
echo "🔨 Compiling WCP360..."
go build -o wcp360 ./cmd/wcp360/main.go

# Setup Production Folders
sudo mkdir -p /var/www/wcp360/data/www
sudo mkdir -p /var/www/wcp360/backups
sudo mkdir -p /var/www/wcp360/configs/dns
sudo cp wcp360 /var/www/wcp360/
sudo cp -r ui /var/www/wcp360/

# Create Systemd Service
cat <<SERVICE | sudo tee /etc/systemd/system/wcp360.service
[Unit]
Description=WebControlPanel360 Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/var/www/wcp360
ExecStart=/var/www/wcp360/wcp360
Restart=always

[Install]
WantedBy=multi-user.target
SERVICE

# Start Service
sudo systemctl daemon-reload
sudo systemctl enable --now wcp360

echo "✅ WCP360 is now installed and running as a system service!"
echo "📍 Access your panel at http://YOUR_SERVER_IP:8080"
echo "🔐 Credentials: admin / wcp360-secure"
