#!/bin/bash
set -ex

# service install
useradd -d /opt/tomodev -m -U tomo
mkdir -p /opt/tomodev/todos/
chmod -R 755 /opt/tomodev

go mod download
go build

cp todos production.toml .env todos.sh /opt/tomodev/todos/
cp todos.service /etc/systemd/system/
cd /opt/tomodev/todos/

chmod 755 todos.sh
systemctl daemon-reload
systemctl enable todos
systemctl start todos
