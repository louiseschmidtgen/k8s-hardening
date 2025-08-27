package hardening

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type USG struct {
	Apply *bool `json:"apply,omitempty" yaml:"apply,omitempty"`
}

type UFW struct {
	Apply *bool `json:"apply,omitempty" yaml:"apply,omitempty"`
}

type AuditLogging struct {
	Apply           *bool   `json:"apply,omitempty" yaml:"apply,omitempty"`
	AuditPolicyPath *string `json:"audit-policy-path,omitempty" yaml:"audit-policy-path,omitempty"`
	AuditPolicyName *string `json:"audit-policy-name,omitempty" yaml:"audit-policy-name,omitempty"`
}

type SystemTuning struct {
	Apply            *bool              `json:"apply,omitempty" yaml:"apply,omitempty"`
	SystemParameters map[string]*string `json:"system-parameters,omitempty" yaml:"system-parameters,omitempty"`
}

type SSHDRemoval struct {
	Apply *bool `json:"apply,omitempty" yaml:"apply,omitempty"`
}

type HardeningConfig struct {
	USG          USG          `json:"usg,omitempty" yaml:"usg,omitempty"`
	UFW          UFW          `json:"ufw,omitempty" yaml:"ufw,omitempty"`
	AuditLogging AuditLogging `json:"audit-logging,omitempty" yaml:"audit-logging,omitempty"`
	SystemTuning SystemTuning `json:"system-tuning,omitempty" yaml:"system-tuning,omitempty"`
	SSHDRemoval  SSHDRemoval  `json:"sshd-removal,omitempty" yaml:"sshd-removal,omitempty"`
}

func LoadConfig(baseline, tailoringFile string) (*HardeningConfig, error) {
	// Load baseline config

	baselineData, err := os.ReadFile(fmt.Sprintf("profiles/%s.yaml", baseline))
	if err != nil {
		return nil, fmt.Errorf("failed to read baseline config file: %w", err)
	}
	var baselineConfig HardeningConfig
	if err := yaml.UnmarshalStrict(baselineData, &baselineConfig); err != nil {
		return nil, fmt.Errorf("failed to parse YAML config file: %w", err)
	}

	// Load tailoring config if provided
	if tailoringFile != "" {
		tailorData, err := os.ReadFile(tailoringFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read tailoring config file: %w", err)
		}
		var tailorConfig HardeningConfig
		if err := yaml.UnmarshalStrict(tailorData, &tailorConfig); err != nil {
			return nil, fmt.Errorf("failed to parse YAML config file: %w", err)
		}

		// Merge tailoring config into baselineConfig (tailoring takes precedence)
		mergeConfigs(&baselineConfig, &tailorConfig)
	}

	return &baselineConfig, nil
}

func mergeConfigs(baseConfig *HardeningConfig, tailorConfig *HardeningConfig) {

	if tailorConfig.USG.Apply != nil {
		baseConfig.USG.Apply = tailorConfig.USG.Apply
	}
	if tailorConfig.UFW.Apply != nil {
		baseConfig.UFW.Apply = tailorConfig.UFW.Apply
	}
	if tailorConfig.AuditLogging.Apply != nil {
		baseConfig.AuditLogging.Apply = tailorConfig.AuditLogging.Apply
	}
	if tailorConfig.AuditLogging.AuditPolicyPath != nil {
		baseConfig.AuditLogging.AuditPolicyPath = tailorConfig.AuditLogging.AuditPolicyPath
	}
	if tailorConfig.AuditLogging.AuditPolicyName != nil {
		baseConfig.AuditLogging.AuditPolicyName = tailorConfig.AuditLogging.AuditPolicyName
	}
	if tailorConfig.SystemTuning.Apply != nil {
		baseConfig.SystemTuning.Apply = tailorConfig.SystemTuning.Apply
	}
	if tailorConfig.SystemTuning.SystemParameters != nil {
		baseConfig.SystemTuning.SystemParameters = tailorConfig.SystemTuning.SystemParameters
	}
	if tailorConfig.SSHDRemoval.Apply != nil {
		baseConfig.SSHDRemoval.Apply = tailorConfig.SSHDRemoval.Apply
	}
}
