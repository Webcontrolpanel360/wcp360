#!/bin/bash

# Couleurs pour le terminal
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}=======================================${NC}"
echo -e "${GREEN}    WCP360 - Auto-Installer v2.2.0    ${NC}"
echo -e "${BLUE}=======================================${NC}"

# 1. Mise à jour du système
echo -e "${BLUE}[1/5]${NC} Mise à jour des paquets..."
apt-get update && apt-get upgrade -y

# 2. Installation des dépendances système
echo -e "${BLUE}[2/5]${NC} Installation de Caddy, PHP 8.3, MySQL et Quotas..."
apt-get install -y debian-keyring debian-archive-keyring apt-transport-https curl
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | tee /etc/apt/sources.list.d/caddy-stable.list
apt-get update
apt-get install -y caddy php8.3-fpm php8.3-mysql mariadb-server quota git golang-go

# 3. Configuration de l'arborescence WCP360
echo -e "${BLUE}[3/5]${NC} Configuration des dossiers..."
mkdir -p /var/www/wcp360/data/{users,vhosts,backups}
mkdir -p /etc/caddy/vhosts

# 4. Compilation du Panel
echo -e "${BLUE}[4/5]${NC} Compilation du binaire Go..."
go build -o wcp360-bin cmd/wcp360/main.go

# 5. Création du service Systemd (Auto-start)
echo -e "${BLUE}[5/5]${NC} Configuration du service système..."
cat <<SERVICE > /etc/systemd/system/wcp360.service
[Unit]
Description=WCP360 Control Panel
After=network.target mysql.service caddy.service

[Service]
Type=simple
User=root
WorkingDirectory=$(pwd)
ExecStart=$(pwd)/wcp360-bin
Restart=always

[Install]
WantedBy=multi-user.target
SERVICE

systemctl daemon-reload
systemctl enable wcp360
systemctl start wcp360

echo -e "${GREEN}=======================================${NC}"
echo -e " INSTALLATION TERMINÉE ! "
echo -e " Accès : http://$(curl -s ifconfig.me):8080"
echo -e " Identifiants : admin / wcp360-secure"
echo -e "${GREEN}=======================================${NC}"
