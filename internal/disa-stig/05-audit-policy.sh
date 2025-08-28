#!/bin/bash
set -euo pipefail

# Usage: ./05-audit-policy.sh <directory> <filename>
# Example: ./05-audit-policy.sh /etc/kubernetes audit-policy.yaml

if [ "$#" -ne 2 ]; then
		echo "Usage: $0 <directory> <filename>"
		exit 1
fi

DIR="$1"
NAME="$2"

if [ -z "$DIR" ] || [ -z "$NAME" ]; then
		echo "Invalid audit policy path or name"
		exit 1
fi

AUDIT_POLICY_PATH="$DIR/$NAME"

AUDIT_POLICY_CONTENT="# Log all requests at the Metadata level.
apiVersion: audit.k8s.io/v1
kind: Policy
rules:
	- level: Metadata
"

echo "$AUDIT_POLICY_CONTENT" | sudo tee "$AUDIT_POLICY_PATH" > /dev/null
sudo chmod 644 "$AUDIT_POLICY_PATH"

echo "Audit policy written to $AUDIT_POLICY_PATH"