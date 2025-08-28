#!/bin/bash
set -euo pipefail

# Set Artifactory Credentials (must be exported before running)
if [[ -z "${ARTIFACTORY_USER:-}" || -z "${ARTIFACTORY_PASS:-}" ]]; then
  echo "Please export ARTIFACTORY_USER and ARTIFACTORY_PASS before running this script."
  exit 1
fi

# Download Artifactory Signing Keys
curl -u "$ARTIFACTORY_USER:$ARTIFACTORY_PASS" \
  "https://canonical.jfrog.io/artifactory/api/security/keypair/soss-deb-stable-local-signing-keys/public" \
  | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/soss-deb-stable.gpg > /dev/null

# Pin-Priority for the Artifactory Repository
cat <<EOF | sudo tee /etc/apt/preferences.d/ubuntu-fips
Package: *
Pin: origin "canonical.jfrog.io"
Pin-Priority: 1001
EOF

# Enable Artifactory Authentication
cat <<EOF | sudo tee /etc/apt/auth.conf.d/soss-deb-stable.conf
machine https://canonical.jfrog.io
login $ARTIFACTORY_USER
password $ARTIFACTORY_PASS
EOF
sudo chmod 600 /etc/apt/auth.conf.d/soss-deb-stable.conf

# Enable the soss-stig-fips Noble repository
REPO=soss-stig-fips
RELEASE=noble
echo "deb https://canonical.jfrog.io/artifactory/$REPO/ $RELEASE main" | sudo tee /etc/apt/sources.list.d/$REPO.list

# Update the APT cache and install the packages
sudo apt update
sudo apt install ubuntu-fips openssh-server --allow-downgrades
sudo apt upgrade -y

# Reboot the system
sudo reboot

# Check the running kernel, installed packages and proc filesystem
uname -a
dpkg -l | grep -i fips
cat /proc/sys/crypto/fips_enabled
