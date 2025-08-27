package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-hardening",
	Short: "Kubernetes hardening tool for DISA-STIG compliance",
	Long:  "Automates DISA-STIG hardening for Canonical Kubernetes clusters.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(fixCmd)
	rootCmd.AddCommand(auditCmd)
}
