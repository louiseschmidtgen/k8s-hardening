#!/bin/bash
set -euo pipefail


# Allow kubernetes ports and services
# API-server
sudo ufw allow 6443/tcp
# kubelet
sudo ufw allow 10250/tcp

# kubernetes daemon
sudo ufw allow 6400/tcp

# CNI cilium
sudo ufw allow 4240/tcp
sudo ufw allow 8472/udp