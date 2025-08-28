#!/bin/bash
set -euo pipefail

if dpkg -l | grep -qw openssh-server; then
	echo "Removing openssh-server..."
	sudo apt-get remove -y openssh-server
else
	echo "openssh-server is not installed. Skipping removal."
fi