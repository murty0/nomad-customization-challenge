package acl

import (
	"testing"

	"github.com/hashicorp/nomad/ci"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCapabilitySet(t *testing.T) {
	ci.Parallel(t)

	var cs capabilitySet = make(map[string]struct{})

	// Check no capabilities by default
	if cs.Check(PolicyDeny) {
		t.Fatalf("unexpected check")
	}

	// Do a set and check
	cs.Set(PolicyDeny)
	if !cs.Check(PolicyDeny) {
		t.Fatalf("missing check")
	}

	// Clear and check
	cs.Clear()
	if cs.Check(PolicyDeny) {
		t.Fatalf("unexpected check")
	}
}

func TestMaxPrivilege(t *testing.T) {
	ci.Parallel(t)

	type tcase struct {
		Privilege      string
		PrecedenceOver []string
	}
	tcases := []tcase{
		{
			PolicyDeny,
			[]string{PolicyDeny, PolicyWrite, PolicyRead, ""},
		},
		{
			PolicyWrite,
			[]string{PolicyWrite, PolicyRead, ""},
		},
		{
			PolicyRead,
			[]string{PolicyRead, ""},
		},
	}

	for idx1, tc := range tcases {
		for idx2, po := range tc.PrecedenceOver {
			if maxPrivilege(tc.Privilege, po) != tc.Privilege {
				t.Fatalf("failed %d %d", idx1, idx2)
			}
			if maxPrivilege(po, tc.Privilege) != tc.Privilege {
				t.Fatalf("failed %d %d", idx1, idx2)
			}
		}
	}
}

func TestACLManagement(t *testing.T) {
	ci.Parallel(t)

	assert := assert.New(t)

	// Create management ACL
	acl, err := NewACL(true, nil)
	assert.Nil(err)

	// Check default namespace rights
	assert.True(acl.AllowNamespaceOperation("default", NamespaceCapabilityListJobs))
	assert.True(acl.AllowNamespaceOperation("default", NamespaceCapabilitySubmitJob))
	assert.True(acl.AllowNamespace("default"))

	// Check non-specified namespace
	assert.True(acl.AllowNamespaceOperation("foo", NamespaceCapabilityListJobs))
	assert.True(acl.AllowNamespace("foo"))

	// Check the other simpler operations
	assert.True(acl.IsManagement())
	assert.True(acl.AllowAgentRead())
	assert.True(acl.AllowAgentWrite())
	assert.True(acl.AllowNodeRead())
	assert.True(acl.AllowNodeWrite())
	assert.True(acl.AllowOperatorRead())
	assert.True(acl.AllowOperatorWrite())
	assert.True(acl.AllowQuotaRead())
	assert.True(acl.AllowQuotaWrite())
}

func TestACLMerge(t *testing.T) {
	ci.Parallel(t)

	assert := assert.New(t)

	// Merge read + write policy
	p1, err := Parse(readAll)
	assert.Nil(err)
	p2, err := Parse(writeAll)
	assert.Nil(err)
	acl, err := NewACL(false, []*Policy{p1, p2})
	assert.Nil(err)

	// Check default namespace rights
	assert.True(acl.AllowNamespaceOperation("default", NamespaceCapabilityListJobs))
	assert.True(acl.AllowNamespaceOperation("default", NamespaceCapabilitySubmitJob))
	assert.True(acl.AllowNamespace("default"))

	// Check non-specified namespace
	assert.False(acl.AllowNamespaceOperation("foo", NamespaceCapabilityListJobs))
	assert.False(acl.AllowNamespace("foo"))

	// Check the other simpler operations
	assert.False(acl.IsManagement())
	assert.True(acl.AllowAgentRead())
	assert.True(acl.AllowAgentWrite())
	assert.True(acl.AllowNodeRead())
	assert.True(acl.AllowNodeWrite())
	assert.True(acl.AllowOperatorRead())
	assert.True(acl.AllowOperatorWrite())
	assert.True(acl.AllowQuotaRead())
	assert.True(acl.AllowQuotaWrite())

	// Merge read + blank
	p3, err := Parse("")
	assert.Nil(err)
	acl, err = NewACL(false, []*Policy{p1, p3})
	assert.Nil(err)

	// Check default namespace rights
	assert.True(acl.AllowNamespaceOperation("default", NamespaceCapabilityListJobs))
	assert.False(acl.AllowNamespaceOperation("default", NamespaceCapabilitySubmitJob))

	// Check non-specified namespace
	assert.False(acl.AllowNamespaceOperation("foo", NamespaceCapabilityListJobs))

	// Check the other simpler operations
	assert.False(acl.IsManagement())
	assert.True(acl.AllowAgentRead())
	assert.False(acl.AllowAgentWrite())
	assert.True(acl.AllowNodeRead())
	assert.False(acl.AllowNodeWrite())
	assert.True(acl.AllowOperatorRead())
	assert.False(acl.AllowOperatorWrite())
	assert.True(acl.AllowQuotaRead())
	assert.False(acl.AllowQuotaWrite())

	// Merge read + deny
	p4, err := Parse(denyAll)
	assert.Nil(err)
	acl, err = NewACL(false, []*Policy{p1, p4})
	assert.Nil(err)

	// Check default namespace rights
	assert.False(acl.AllowNamespaceOperation("default", NamespaceCapabilityListJobs))
	assert.False(acl.AllowNamespaceOperation("default", NamespaceCapabilitySubmitJob))

	// Check non-specified namespace
	assert.False(acl.AllowNamespaceOperation("foo", NamespaceCapabilityListJobs))

	// Check the other simpler operations
	assert.False(acl.IsManagement())
	assert.False(acl.AllowAgentRead())
	assert.False(acl.AllowAgentWrite())
	assert.False(acl.AllowNodeRead())
	assert.False(acl.AllowNodeWrite())
	assert.False(acl.AllowOperatorRead())
	assert.False(acl.AllowOperatorWrite())
	assert.False(acl.AllowQuotaRead())
	assert.False(acl.AllowQuotaWrite())
}

var readAll = `
namespace "default" {
	policy = "read"
}
agent {
	policy = "read"
}
node {
	policy = "read"
}
operator {
	policy = "read"
}
quota {
	policy = "read"
}
`

var writeAll = `
namespace "default" {
	policy = "write"
}
agent {
	policy = "write"
}
node {
	policy = "write"
}
operator {
	policy = "write"
}
quota {
	policy = "write"
}
`

var denyAll = `
namespace "default" {
	policy = "deny"
}
agent {
	policy = "deny"
}
node {
	policy = "deny"
}
operator {
	policy = "deny"
}
quota {
	policy = "deny"
}
`

func TestAllowNamespace(t *testing.T) {
	ci.Parallel(t)

	tests := []struct {
		name      string
		policy    string
		allow     bool
		namespace string
	}{
		{
			name:      "foo namespace - no capabilities",
			policy:    `namespace "foo" {}`,
			allow:     false,
			namespace: "foo",
		},
		{
			name:      "foo namespace - deny policy",
			policy:    `namespace "foo" { policy = "deny" }`,
			allow:     false,
			namespace: "foo",
		},
		{
			name:      "foo namespace - deny capability",
			policy:    `namespace "foo" { capabilities = ["deny"] }`,
			allow:     false,
			namespace: "foo",
		},
		{
			name:      "foo namespace - with capability",
			policy:    `namespace "foo" { capabilities = ["list-jobs"] }`,
			allow:     true,
			namespace: "foo",
		},
		{
			name:      "foo namespace - with policy",
			policy:    `namespace "foo" { policy = "read" }`,
			allow:     true,
			namespace: "foo",
		},
		{
			name:      "wildcard namespace - no capabilities",
			policy:    `namespace "foo" {}`,
			allow:     false,
			namespace: "*",
		},
		{
			name:      "wildcard namespace - deny policy",
			policy:    `namespace "foo" { policy = "deny" }`,
			allow:     false,
			namespace: "*",
		},
		{
			name:      "wildcard namespace - deny capability",
			policy:    `namespace "foo" { capabilities = ["deny"] }`,
			allow:     false,
			namespace: "*",
		},
		{
			name:      "wildcard namespace - with capability",
			policy:    `namespace "foo" { capabilities = ["list-jobs"] }`,
			allow:     true,
			namespace: "*",
		},
		{
			name:      "wildcard namespace - with policy",
			policy:    `namespace "foo" { policy = "read" }`,
			allow:     true,
			namespace: "*",
		},
		{
			name:      "wildcard namespace - no namespace rule",
			policy:    `agent { policy = "read" }`,
			allow:     false,
			namespace: "*",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			policy, err := Parse(tc.policy)
			require.NoError(t, err)

			acl, err := NewACL(false, []*Policy{policy})
			require.NoError(t, err)

			got := acl.AllowNamespace(tc.namespace)
			require.Equal(t, tc.allow, got)
		})
	}
}

func TestWildcardNamespaceMatching(t *testing.T) {
	ci.Parallel(t)

	tests := []struct {
		name      string
		policy    string
		allow     bool
		namespace string
	}{
		{
			name:      "wildcard matches",
			policy:    `namespace "prod-api-*" { policy = "write" }`,
			allow:     true,
			namespace: "prod-api-services",
		},
		{
			name:      "non globbed namespaces are not wildcards",
			policy:    `namespace "prod-api" { policy = "write" }`,
			allow:     false,
			namespace: "prod-api-services",
		},
		{
			name: "concrete matches take precedence",
			policy: `namespace "prod-api-services" { policy = "deny" }
			         namespace "prod-api-*" { policy = "write" }`,
			allow:     false,
			namespace: "prod-api-services",
		},
		{
			name: "glob match",
			policy: `namespace "prod-api-*" { policy = "deny" }
			         namespace "prod-api-services" { policy = "write" }`,
			allow:     true,
			namespace: "prod-api-services",
		},
		{
			name: "closest character match wins - suffix",
			policy: `namespace "*-api-services" { policy = "deny" }
			         namespace "prod-api-*" { policy = "write" }`, // 4 vs 8 chars
			allow:     false,
			namespace: "prod-api-services",
		},
		{
			name: "closest character match wins - prefix",
			policy: `namespace "prod-api-*" { policy = "write" }
               namespace "*-api-services" { policy = "deny" }`, // 4 vs 8 chars
			allow:     false,
			namespace: "prod-api-services",
		},
		{
			name: "wildcard namespace with glob match",
			policy: `namespace "prod-api-*" { policy = "deny" }
			         namespace "prod-api-services" { policy = "write" }`,
			allow:     true,
			namespace: "*",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			policy, err := Parse(tc.policy)
			require.NoError(t, err)
			require.NotNil(t, policy.Namespaces)

			acl, err := NewACL(false, []*Policy{policy})
			require.NoError(t, err)

			got := acl.AllowNamespace(tc.namespace)
			require.Equal(t, tc.allow, got)
		})
	}
}

func TestWildcardHostVolumeMatching(t *testing.T) {
	ci.Parallel(t)

	tests := []struct {
		Policy string
		Allow  bool
	}{
		{ // Wildcard matches
			Policy: `host_volume "prod-api-*" { policy = "write" }`,
			Allow:  true,
		},
		{ // Non globbed volumes are not wildcards
			Policy: `host_volume "prod-api" { policy = "write" }`,
			Allow:  false,
		},
		{ // Concrete matches take precedence
			Policy: `host_volume "prod-api-services" { policy = "deny" }
			         host_volume "prod-api-*" { policy = "write" }`,
			Allow: false,
		},
		{
			Policy: `host_volume "prod-api-*" { policy = "deny" }
			         host_volume "prod-api-services" { policy = "write" }`,
			Allow: true,
		},
		{ // The closest character match wins
			Policy: `host_volume "*-api-services" { policy = "deny" }
			         host_volume "prod-api-*" { policy = "write" }`, // 4 vs 8 chars
			Allow: false,
		},
		{
			Policy: `host_volume "prod-api-*" { policy = "write" }
               host_volume "*-api-services" { policy = "deny" }`, // 4 vs 8 chars
			Allow: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Policy, func(t *testing.T) {
			assert := assert.New(t)

			policy, err := Parse(tc.Policy)
			assert.NoError(err)
			assert.NotNil(policy.HostVolumes)

			acl, err := NewACL(false, []*Policy{policy})
			assert.Nil(err)

			assert.Equal(tc.Allow, acl.AllowHostVolume("prod-api-services"))
		})
	}
}
func TestACL_matchingCapabilitySet_returnsAllMatches(t *testing.T) {
	ci.Parallel(t)

	tests := []struct {
		Policy        string
		NS            string
		MatchingGlobs []string
	}{
		{
			Policy:        `namespace "production-*" { policy = "write" }`,
			NS:            "production-api",
			MatchingGlobs: []string{"production-*"},
		},
		{
			Policy:        `namespace "prod-*" { policy = "write" }`,
			NS:            "production-api",
			MatchingGlobs: nil,
		},
		{
			Policy: `namespace "production-*" { policy = "write" }
							 namespace "production-*-api" { policy = "deny" }`,

			NS:            "production-admin-api",
			MatchingGlobs: []string{"production-*", "production-*-api"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Policy, func(t *testing.T) {
			assert := assert.New(t)

			policy, err := Parse(tc.Policy)
			assert.NoError(err)
			assert.NotNil(policy.Namespaces)

			acl, err := NewACL(false, []*Policy{policy})
			assert.Nil(err)

			var namespaces []string
			for _, cs := range findAllMatchingWildcards(acl.wildcardNamespaces, tc.NS) {
				namespaces = append(namespaces, cs.name)
			}

			assert.Equal(tc.MatchingGlobs, namespaces)
		})
	}
}

func TestACL_matchingCapabilitySet_difference(t *testing.T) {
	ci.Parallel(t)

	tests := []struct {
		Policy     string
		NS         string
		Difference int
	}{
		{
			Policy:     `namespace "production-*" { policy = "write" }`,
			NS:         "production-api",
			Difference: 3,
		},
		{
			Policy:     `namespace "production-*" { policy = "write" }`,
			NS:         "production-admin-api",
			Difference: 9,
		},
		{
			Policy:     `namespace "production-**" { policy = "write" }`,
			NS:         "production-admin-api",
			Difference: 9,
		},
		{
			Policy:     `namespace "*" { policy = "write" }`,
			NS:         "production-admin-api",
			Difference: 20,
		},
		{
			Policy:     `namespace "*admin*" { policy = "write" }`,
			NS:         "production-admin-api",
			Difference: 15,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Policy, func(t *testing.T) {
			assert := assert.New(t)

			policy, err := Parse(tc.Policy)
			assert.NoError(err)
			assert.NotNil(policy.Namespaces)

			acl, err := NewACL(false, []*Policy{policy})
			assert.Nil(err)

			matches := findAllMatchingWildcards(acl.wildcardNamespaces, tc.NS)
			assert.Equal(tc.Difference, matches[0].difference)
		})
	}

}
