#!/bin/bash
set -euo pipefail


# Usage: sudo bash 01-fips.sh [pro_token]

PRO_TOKEN=""
if [ "$#" -gt 1 ]; then
	echo "Usage: $0 [pro_token]"
	exit 1
elif [ "$#" -eq 1 ]; then
	PRO_TOKEN="$1"
fi


sudo apt updates
sudo apt install -y ubuntu-advantage-tools


echo "Checking Ubuntu Pro status..."
if pro status | grep -q "The system is not attached"; then
	if [ -z "$PRO_TOKEN" ]; then
		echo "ERROR: Ubuntu Pro is not enabled and no token was provided. Please pass your Pro token as an argument."
		exit 1
	fi
	echo "Attaching Ubuntu Pro..."
	sudo pro attach "$PRO_TOKEN" --no-auto-enable
else
	echo "Ubuntu Pro is already enabled."
fi

sudo pro enable fips-updates

sudo snap install k8s --channel=1.34-classic/stable --classic

# TODO can we run this before reboot?
echo "Installing core22 snap with FIPS channel..."
# Install the FIPS core base snap
if snap list | grep -q '^core22 '; then
	echo "core22 is already installed. Refreshing to fips-updates/stable channel."
	sudo snap refresh core22 --channel=fips-updates/stable
else
	echo "Installing core22 with FIPS-certified libraries."
	sudo snap install core22 --channel=fips-updates/stable
fi


if [ -f /proc/sys/crypto/fips_enabled ] && [ "$(cat /proc/sys/crypto/fips_enabled)" -eq 1 ]; then
	echo "FIPS is already enabled. Skipping reboot."
else
	echo "Rebooting to apply FIPS kernel..."
	sudo reboot
fi