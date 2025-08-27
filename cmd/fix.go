package cmd

import (
	"github.com/louiseschmidtgen/k8s-hardening/internal/hardening"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Apply automated hardening rules",
	Run: func(cmd *cobra.Command, args []string) {
		if err := hardening.ApplyFix(baseline, tailoringFile, nodeRole); err != nil {
			cmd.PrintErrf("Error: failed to apply hardening rules.\n\n The error was %v", err)
			return
		}
		cmd.Printf("Automatic rules for the strict DISA STIG profile were applied. Further manual steps are required for full compliance.\n\nPlease consult https://documentation.ubuntu.com/canonical-kubernetes/latest/snap/howto/security/disa-stig-assessment for guidance.\n")
	},
}

var baseline string
var tailoringFile string
var nodeRole string

func init() {
	fixCmd.Flags().StringVar(&baseline, "baseline", "disa-stig", "Baseline profile to apply")
	fixCmd.Flags().StringVar(&tailoringFile, "tailoring-file", "", "Override configuration file")
	fixCmd.Flags().StringVar(&nodeRole, "node-role", "", "Node role: control-plane or worker")
}
