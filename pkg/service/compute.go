package service

import (
	"context"
)

type ComputeAPI interface {
	// GetVirtualMachinesByLabel returns all virtual machine for given filter input.
	GetVirtualMachinesByLabel(ctx context.Context, zones []string, label *Label) ([]*VirtualMachine, error)

	// GetFirewallRulesForRegionalPolicy returns all firewall rules for a specific policy.
	GetFirewallRulesForRegionalPolicy(ctx context.Context, project, region, name string) ([]*FirewallRule, error)
}

type ComputeService interface {
	// GetAllowedFirewallTags returns all firewall tags available.
	GetAllowedFirewallTags(ctx context.Context) ([]string, error)
}

type VirtualMachine struct {
	Name               string
	ID                 uint64
	Zone               string
	FullyQualifiedZone string
}

type FirewallRule struct {
	Name        string
	SecureTags  []string
	Description string
}

type Label struct {
	Key   string
	Value string
}
