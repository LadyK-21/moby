//go:build freebsd

package daemon

import (
	"testing"

	"gotest.tools/v3/assert"
)

func cleanupNetworkNamespace(_ testing.TB, _ *Daemon) {}

// CgroupNamespace returns the cgroup namespace the daemon is running in
func (d *Daemon) CgroupNamespace(t testing.TB) string {
	assert.Assert(t, false, "cgroup namespaces are not supported on FreeBSD")
	return ""
}
