package drone_info

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
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

func TestMockDroneInfoDroneSystemRefs(t *testing.T) {
	// mock MockDroneInfoDroneSystemRefs
	type args struct {
		droneProto    string
		droneHost     string
		droneHostName string
		status        string
		refs          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "proto error", // testdata/TestMockDroneInfoDroneSystemRefs/sample.golden
			args: args{
				droneProto:    "git",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/heads/main",
			},
			wantErr: fmt.Errorf("droneProto only support http or https, now is: git"),
		},
		{
			name: "droneHost error",
			args: args{
				droneProto:    "http",
				droneHost:     "",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/heads/main",
			},
			wantErr: fmt.Errorf("droneHost not support nil"),
		},
		{
			name: "droneHostName error",
			args: args{
				droneProto:    "http",
				droneHost:     "drone.company.com",
				droneHostName: "",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/heads/main",
			},
			wantErr: fmt.Errorf("droneHostName not support nil"),
		},
		{
			name: "refs empty",
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
			},
			wantErr: fmt.Errorf("refs not support nil"),
		},
		{
			name: "refs not has refs/",
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "heads/main",
			},
			wantErr: fmt.Errorf("refs not has prefix refs/ , now is: heads/main"),
		},
		{
			name: "refs path error",
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/main",
			},
			wantErr: fmt.Errorf("refs parase error, now is: refs/main"),
		},
		{
			name: "refs type error",
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/foo/main",
			},
			wantErr: fmt.Errorf("not support refsType by refs: refs/foo/main"),
		},
		{
			name: "default", // testdata/TestMockDroneInfoDroneSystemRefs/default.golden
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				refs:          "refs/heads/main",
			},
		},
		{
			name: "sample", // testdata/TestMockDroneInfoDroneSystemRefs/sample.golden
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/heads/main",
			},
		},
		{
			name: "sample failure", // testdata/TestMockDroneInfoDroneSystemRefs/sample_failure.golden
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusFailure,
				refs:          "refs/heads/main",
			},
		},
		{
			name: "sample tag", // testdata/TestMockDroneInfoDroneSystemRefs/sample_tag.golden
			args: args{
				droneProto:    "https",
				droneHost:     "drone.company.com",
				droneHostName: "drone.company.com",
				status:        DroneBuildStatusSuccess,
				refs:          "refs/tags/v1.5.0",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do MockDroneInfoDroneSystemRefs
			gotResult, gotErr := MockDroneInfoDroneSystemRefs(
				tc.args.droneProto,
				tc.args.droneHost,
				tc.args.droneHostName,
				tc.args.status,
				tc.args.refs,
			)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			// verify MockDroneInfoDroneSystemRefs

			// fix different platform path run error
			gotResult.Build.WorkSpace = ""
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
