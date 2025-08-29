#!/bin/bash
set -euo pipefail

# Install usg tool
sudo apt update
sudo apt upgrade -y 
sudo apt install -y usg

# Generate a DISA STIG compliance audit report
# usg audit disa_stig

# Automatically apply recommended hardening changes
sudo pro enable usg
sudo usg fix disa_stig

echo "DISA STIG host compliance steps completed."
echo "Run 'sudo usg audit disa_stig' to check the host's compliance status."
