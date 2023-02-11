package drone_info

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMockDroneInfoFull(t *testing.T) {
	// mock MockDroneInfoEnvFull
	MockDroneInfoEnvFull(true)
	t.Logf("~> mock MockDroneInfoEnvFull")
	// do MockDroneInfoEnvFull
	t.Logf("~> do MockDroneInfoEnvFull")
	MockEnvDebugPrint()
	// verify MockDroneInfoEnvFull

	assert.Equal(t, mockEnvDroneRepoOwner, os.Getenv(EnvDroneRepoOwner))
}

func TestMockDroneInfo(t *testing.T) {
	// mock MockDroneInfo

	droneInfo := MockDroneInfo(mockEnvDroneBuildStatusSuccess)
	t.Logf("~> mock MockDroneInfo")
	// do MockDroneInfo
	t.Logf("~> do MockDroneInfo")
	// verify MockDroneInfo
	assert.Equal(t, mockEnvDroneRepoOwner, droneInfo.Repo.OwnerName)
	assert.Equal(t, mockEnvDroneRepoOwner, droneInfo.Repo.GroupName)
	assert.Equal(t, mockEnvDroneRepo, droneInfo.Repo.ShortName)
	assert.Equal(t, mockEnvDroneBuildStatusSuccess, droneInfo.Build.Status)

	droneInfoFail := MockDroneInfo(mockEnvDroneBuildStatusFailure)
	assert.Equal(t, droneInfoFail.Build.Status, mockEnvDroneBuildStatusFailure)
}

func TestMockDroneInfoRefs(t *testing.T) {
	// mock MockDroneInfoRefs

	t.Logf("~> mock MockDroneInfoRefs")
	droneInfoRefs, err := MockDroneInfoRefs(DroneBuildStatusSuccess, fmt.Sprintf("refs/head/%s", mockEnvDroneBranch))
	if err == nil {
		t.Fatalf("not check head to support")
	}
	droneInfoRefs, err = MockDroneInfoRefs(DroneBuildStatusSuccess, fmt.Sprintf("refs/heads/%s", mockEnvDroneBranch))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, mockEnvDroneRepoOwner, droneInfoRefs.Repo.OwnerName)
	assert.Equal(t, mockEnvDroneBranch, droneInfoRefs.Commit.Branch)
	assert.Equal(t, mockEnvDroneBranch, droneInfoRefs.Build.Branch)

	droneInfoRefs, err = MockDroneInfoRefs(DroneBuildStatusSuccess, fmt.Sprintf("refs/heads/%s", "features/baz"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, mockEnvDroneRepoOwner, droneInfoRefs.Repo.OwnerName)
	assert.Equal(t, "features/baz", droneInfoRefs.Commit.Branch)
	assert.Equal(t, "features/baz", droneInfoRefs.Build.Branch)

	droneInfoRefs, err = MockDroneInfoRefs(DroneBuildStatusSuccess, fmt.Sprintf("refs/tags/%s", "v1.2.3"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, mockEnvDroneRepoOwner, droneInfoRefs.Repo.OwnerName)
	assert.Equal(t, "", droneInfoRefs.Commit.Branch)
	assert.Equal(t, "", droneInfoRefs.Build.Branch)
	assert.Equal(t, "1.2.3", droneInfoRefs.Build.Tag)

}
