package daemon

import (
	"testing"

	"github.com/docker/docker/api/types/container"
)

func TestToContainerdResources_Defaults(t *testing.T) {
	checkResourcesAreUnset(t, toContainerdResources(container.Resources{}))
}
