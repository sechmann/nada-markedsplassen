package computeengine_test

import (
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"github.com/navikt/nada-backend/pkg/computeengine"
	"github.com/navikt/nada-backend/pkg/computeengine/emulator"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"testing"
)

func strPtr(s string) *string {
	return &s
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}

func TestNewInstancesClient(t *testing.T) {
	testCases := []struct {
		name      string
		project   string
		zones     []string
		filter    string
		instances map[string][]*computepb.Instance
		expect    []*computeengine.VirtualMachine
	}{
		{
			name:      "no instances",
			project:   "test",
			zones:     []string{"europe-north1-b", "europe-north1-c"},
			filter:    "",
			instances: map[string][]*computepb.Instance{},
			expect:    nil,
		},
		{
			name:    "one instance, in one zone",
			project: "test",
			zones:   []string{"europe-north1-b"},
			filter:  "",
			instances: map[string][]*computepb.Instance{
				"europe-north1-b": {
					{
						Name: strPtr("test-instance"),
						Id:   uint64Ptr(123),
						Zone: strPtr("https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-b"),
					},
				},
			},
			expect: []*computeengine.VirtualMachine{
				{
					Name:               "test-instance",
					ID:                 123,
					FullyQualifiedZone: "https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-b",
					Zone:               "europe-north1-b",
				},
			},
		},
		{
			name:    "instance in multiple zones",
			project: "test",
			zones:   []string{"europe-north1-b", "europe-north1-c"},
			filter:  "",
			instances: map[string][]*computepb.Instance{
				"europe-north1-b": {
					{
						Name: strPtr("test-instance-1"),
						Id:   uint64Ptr(123),
						Zone: strPtr("https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-b"),
					},
				},
				"europe-north1-c": {
					{
						Name: strPtr("test-instance-2"),
						Id:   uint64Ptr(1231234),
						Zone: strPtr("https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-c"),
					},
				},
			},
			expect: []*computeengine.VirtualMachine{
				{
					Name:               "test-instance-1",
					ID:                 123,
					FullyQualifiedZone: "https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-b",
					Zone:               "europe-north1-b",
				},
				{
					Name:               "test-instance-2",
					ID:                 1231234,
					FullyQualifiedZone: "https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-c",
					Zone:               "europe-north1-c",
				},
			},
		},
		{
			name:    "one instance from multiple zones",
			project: "test",
			zones:   []string{"europe-north1-b"},
			filter:  "",
			instances: map[string][]*computepb.Instance{
				"europe-north1-b": {
					{
						Name: strPtr("test-instance-1"),
						Id:   uint64Ptr(123),
						Zone: strPtr("https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-b"),
					},
				},
				"europe-north1-c": {
					{
						Name: strPtr("test-instance-2"),
						Id:   uint64Ptr(1231234),
						Zone: strPtr("https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-c"),
					},
				},
			},
			expect: []*computeengine.VirtualMachine{
				{
					Name:               "test-instance-1",
					ID:                 123,
					FullyQualifiedZone: "https://www.googleapis.com/compute/v1/projects/knada-dev/zones/europe-north1-b",
					Zone:               "europe-north1-b",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			log := zerolog.New(zerolog.NewConsoleWriter())
			e := emulator.New(log)
			e.SetInstances(tc.instances)
			url := e.Run()

			c := computeengine.NewClient(url, true)

			got, err := c.ListVirtualMachines(context.Background(), tc.project, tc.zones, tc.filter)
			require.NoError(t, err)
			require.Equal(t, tc.expect, got)
		})
	}
}
func TestNewFirewallPoliciesClient(t *testing.T) {
	testCases := []struct {
		name               string
		firewallPolicyName string
		firewallPolicies   map[string][]*computepb.FirewallPolicy
		expect             []computeengine.FirewallRule
	}{
		{
			name:               "no firewall rules",
			firewallPolicyName: "finnes ikke",
			firewallPolicies:   map[string][]*computepb.FirewallPolicy{},
			expect:             nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			log := zerolog.New(zerolog.NewConsoleWriter())
			e := emulator.New(log)
			e.SetFirewallPolicies(tc.firewallPolicies)
			url := e.Run()

			c := computeengine.NewClient(url, true)

			got, err := c.GetFirewallRulesForPolicy(context.Background(), tc.firewallPolicyName)
			require.NoError(t, err)
			require.Equal(t, tc.expect, got)
		})
	}
}
