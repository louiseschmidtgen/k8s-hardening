package cmd

import (
	"github.com/louiseschmidtgen/k8s-hardening/internal/hardening"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit hardening rules and configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := hardening.Audit(baseline, tailoringFile, nodeRole); err != nil {
			cmd.PrintErrf("Error: failed to audit hardening rules.\n\n The error was %v", err)
			return
		}
		cmd.Printf("Audit completed successfully. To audit the host machine please run:\n\n  sudo usg audit disa_stig.\n")
	},
}

func init() {
	auditCmd.Flags().StringVar(&baseline, "baseline", "disa-stig", "Baseline profile to audit")
	auditCmd.Flags().StringVar(&tailoringFile, "tailoring-file", "", "Override configuration file")
	auditCmd.Flags().StringVar(&nodeRole, "node-role", "", "Node role: control-plane or worker")
}
