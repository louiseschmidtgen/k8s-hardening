# k8s-hardening

This hardening tool will reduce the manual steps required for users to achieve their desired level of DISA STIG compliance for Canonical Kubernetes.

## Usage instructions

Build the tool by running:

```
go build -o k8s-hardening .
```

Run the disa-stig apply command:

```
sudo ./k8s-hardening fix --baseline=disa-stig --node-role=control-plane
```
