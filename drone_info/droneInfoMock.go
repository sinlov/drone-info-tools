package drone_info

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	mockEnvDroneRepo              = "drone-file-browser-plugin"
	mockEnvDroneCommitAuthorEmail = "sinlovgmppt@gmail.com"
	mockEnvDroneRepoOwner         = "sinlov"
	mockEnvDroneBranch            = "main"
	mockEnvDroneRemoteUrlBase     = "https://github.com"

	mockEnvDroneSystemVersion  = "1.0.0"
	mockEnvDroneSystemHost     = "drone.xxx.com"
	mockEnvDroneSystemHostName = "drone.xxx.com"
	mockEnvDroneSystemProto    = "https"

	mockEnvDroneUrlBase = "https://drone.xxx.com"
	mockEnvDroneRepoScm = "git"

	mockEnvDroneStageStarted  uint64 = 1674531206
	mockEnvDroneStageFinished uint64 = 1674532106
	mockEnvDroneBuildStarted  uint64 = 1674531206
	mockEnvDroneBuildFinished uint64 = 1674532206
	mockEnvDroneBuildNumber   uint64 = 1
	mockEnvDroneStageMachine  string = "CI-machine"
	mockEnvDroneStageOs       string = "linux"
	mockEnvDroneStageArch     string = "amd64"
	mockEnvDroneStageVariant  string = ""
	mockEnvDroneStageType     string = "docker"
	mockEnvDroneStageKind     string = "pipeline"
	mockEnvDroneStageName     string = "build"

	mockEnvDroneTag                = ""
	mockEnvDroneTargetBranch       = ""
	mockEnvDroneBuildEvent         = "push"
	mockEnvDroneBuildStatusSuccess = "success"
	mockEnvDroneBuildStatusFailure = "failure"

	mockEnvDroneCommitMessage = "mock message commit\nmore line\nand more line\r\n"
	mockEnvDroneCommitSha     = "68e3d62dd69f06077a243a1db1460109377add64"

	mockEnvFailedStages = "build,test"
	mockEnvFailedSteps  = "backend,frontend"
)

// MockDroneInfo
// status DroneBuildStatusSuccess DroneBuildStatusFailure
// this mock will use mockEnvDroneBranch as heads by MockDroneInfoRefs
func MockDroneInfo(status string) *Drone {
	droneInfoRefs, _ := MockDroneInfoRefs(status, fmt.Sprintf("refs/heads/%s", mockEnvDroneBranch))
	return droneInfoRefs
}

// MockDroneInfoRefs
// status DroneBuildStatusSuccess DroneBuildStatusFailure
// refs by: git show-ref --head --dereference
// @doc https://git-scm.com/docs/git-show-ref
// like refs/heads/master refs/remotes/* refs/pull/* refs/tags/v1.0.0
func MockDroneInfoRefs(status string, refs string) (*Drone, error) {

	if refs == "" {
		return nil, fmt.Errorf("refs not support nil")
	}
	if !strings.HasPrefix(refs, "refs/") {
		return nil, fmt.Errorf("refs not has prefix refs/ , now is: %s", refs)
	}
	refsPath := strings.Replace(refs, "refs/", "", -1)
	refsSplit := strings.SplitN(refsPath, "/", 2)
	if len(refsSplit) < 2 {
		return nil, fmt.Errorf("refs parase error, now is: %s", refs)
	}

	if status == "" {
		status = mockEnvDroneBuildStatusSuccess
	}

	workspace, _ := getCurrentFolderPath()

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	repoUrl := fmt.Sprintf("%s/%s/%s", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoHttpUrl := fmt.Sprintf("%s/%s/%s.git", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoSshUrl := fmt.Sprintf("git@%s:%s/%s.git", owner, repoName, "github.com")
	repoHost := ""
	repoHostName := ""
	parse, err := url.Parse(repoHttpUrl)
	if err == nil {
		repoHost = parse.Host
		repoHostName = parse.Hostname()
	}
	stageStartT := mockEnvDroneStageStarted
	stageStartTime := time.Unix(int64(stageStartT), 0).Format(DroneTimeFormatDefault)
	stageFinishedT := mockEnvDroneStageFinished
	stageFinishedTime := time.Unix(int64(stageStartT), 0).Format(DroneTimeFormatDefault)
	commitSHA := mockEnvDroneCommitSha
	droneBaseUrl := mockEnvDroneUrlBase
	buildNumber := mockEnvDroneBuildNumber

	var drone = Drone{
		//  repo info
		Repo: Repo{
			Link:      repoUrl,
			ShortName: repoName,
			GroupName: mockEnvDroneRepoOwner,
			FullName:  fmt.Sprintf("%s/%s", owner, repoName),
			OwnerName: mockEnvDroneRepoOwner,
			Scm:       mockEnvDroneRepoScm,
			RemoteURL: repoHttpUrl,
			HttpUrl:   repoHttpUrl,
			SshUrl:    repoSshUrl,
			Host:      repoHost,
			HostName:  repoHostName,
		},
		//  build info
		Build: Build{
			WorkSpace:    workspace,
			Status:       status,
			Number:       buildNumber,
			Tag:          mockEnvDroneTag,
			TargetBranch: mockEnvDroneTargetBranch,
			Link:         fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, buildNumber),
			Event:        mockEnvDroneBuildEvent,
			StartAt:      mockEnvDroneBuildStarted,
			FinishedAt:   mockEnvDroneBuildFinished,
			PR:           "",
			DeployTo:     "",
			FailedStages: mockEnvFailedStages,
			FailedSteps:  mockEnvFailedSteps,
		},
		Commit: Commit{
			Link:    fmt.Sprintf("%s/commit/%s", repoHttpUrl, commitSHA),
			Message: mockEnvDroneCommitMessage,
			Sha:     commitSHA,
			Ref:     refs,
			Author: CommitAuthor{
				Username: owner,
				Email:    email,
				Name:     owner,
				Avatar:   "",
			},
		},
		Stage: Stage{
			StartedAt:    stageStartT,
			StartedTime:  stageStartTime,
			FinishedAt:   stageFinishedT,
			FinishedTime: stageFinishedTime,
			Machine:      mockEnvDroneStageMachine,
			Os:           mockEnvDroneStageOs,
			Arch:         mockEnvDroneStageArch,
			Variant:      mockEnvDroneStageVariant,
			Type:         mockEnvDroneStageType,
			Kind:         mockEnvDroneStageKind,
			Name:         mockEnvDroneStageName,
		},
		DroneSystem: DroneSystem{
			Version:  mockEnvDroneSystemVersion,
			Host:     mockEnvDroneSystemHost,
			HostName: mockEnvDroneSystemHostName,
			Proto:    mockEnvDroneSystemProto,
		},
	}

	refsType := refsSplit[0]
	refsContent := refsSplit[1]
	switch refsType {
	case "tags":
		drone.Build.Tag = strings.Replace(refsContent, "v", "", 1)
		drone.Commit.Branch = ""
		drone.Build.Branch = ""
	case "heads":
		fallthrough
	case "remotes":
		fallthrough
	case "pull":
		drone.Commit.Branch = refsContent
		drone.Build.Branch = refsContent
	default:
		return nil, fmt.Errorf("not support refsType by refs: %s", refs)

	}

	//branch := mockEnvDroneCommitBranch

	return &drone, nil
}

func MockDroneInfoEnvFull(debug bool) {
	setEnvBool("PLUGIN_DEBUG", debug)

	workspace, _ := getCurrentFolderPath()

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	repoUrl := fmt.Sprintf("%s/%s/%s", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoHttpUrl := fmt.Sprintf("%s/%s/%s.git", mockEnvDroneRemoteUrlBase, owner, repoName)
	repoSshUrl := fmt.Sprintf("git@%s:%s/%s.git", owner, repoName, "github.com")

	commitSHA := mockEnvDroneCommitSha
	branch := mockEnvDroneBranch
	droneBaseUrl := mockEnvDroneUrlBase
	buildNumber := mockEnvDroneBuildNumber

	setEnvStr(EnvDroneRepoLink, repoUrl)
	setEnvStr(EnvDroneRepo, fmt.Sprintf("%s/%s", owner, repoName))
	setEnvStr(EnvDroneRepoName, repoName)
	setEnvStr(EnvDroneRepoNamespace, owner)
	setEnvStr(EnvDroneRemoteUrl, repoHttpUrl)
	setEnvStr(EnvDroneRepoOwner, owner)
	setEnvStr(EnvDroneGitHttpUrl, repoHttpUrl)
	setEnvStr(EnvDroneGitSshUrl, repoSshUrl)

	setEnvStr(EnvDroneBuildWorkSpace, workspace)
	setEnvStr(EnvDroneBuildStatus, mockEnvDroneBuildStatusSuccess)
	setEnvU64(EnvDroneBuildNumber, mockEnvDroneBuildNumber)
	setEnvStr(EnvDroneTag, mockEnvDroneTag)
	setEnvStr(EnvDroneTargetBranch, mockEnvDroneTargetBranch)
	setEnvStr(EnvDroneBuildLink, fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, buildNumber))
	setEnvStr(EnvDroneBuildEvent, mockEnvDroneBuildEvent)
	setEnvU64(EnvDroneBuildStarted, mockEnvDroneBuildStarted)
	setEnvU64(EnvDroneBuildFinished, mockEnvDroneBuildFinished)
	setEnvStr(EnvDroneFailedStages, "")
	setEnvStr(EnvDroneFailedSteps, "")

	setEnvStr(EnvDroneCommitAuthor, owner)
	setEnvStr(EnvDroneCommitAuthorName, owner)
	setEnvStr(EnvDroneCommitAuthorAvatar, "")
	setEnvStr(EnvDroneCommitAuthorEmail, email)
	setEnvStr(EnvDroneCommitLink, fmt.Sprintf("%s/commit/%s", repoUrl, commitSHA))
	setEnvStr(EnvDroneBranch, branch)
	setEnvStr(EnvDroneCommitBranch, branch)
	setEnvStr(EnvDroneCommitMessage, mockEnvDroneCommitMessage)
	setEnvStr(EnvDroneCommitSha, commitSHA)
	setEnvStr(EnvDroneCommitRef, fmt.Sprintf("refs/heads/%s", branch))

	setEnvU64(EnvDroneStageStarted, mockEnvDroneStageStarted)
	setEnvU64(EnvDroneStageFinished, mockEnvDroneStageFinished)
	setEnvStr(EnvDroneStageMachine, mockEnvDroneStageMachine)
	setEnvStr(EnvDroneStageOs, mockEnvDroneStageOs)
	setEnvStr(EnvDroneStageArch, mockEnvDroneStageArch)
	setEnvStr(EnvDroneStageVariant, mockEnvDroneStageVariant)
	setEnvStr(EnvDroneStageType, mockEnvDroneStageType)
	setEnvStr(EnvDroneStageKind, mockEnvDroneStageKind)
	setEnvStr(EnvDroneStageName, mockEnvDroneStageName)

	setEnvStr(EnvDroneSystemVersion, mockEnvDroneSystemVersion)
	setEnvStr(EnvDroneSystemHost, mockEnvDroneSystemHost)
	setEnvStr(EnvDroneSystemHostName, mockEnvDroneSystemHostName)
	setEnvStr(EnvDroneSystemProto, mockEnvDroneSystemProto)

}

func MockEnvDebugPrint() {
	envDebug, find := os.LookupEnv("PLUGIN_DEBUG")
	if find && envDebug == "true" {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}
}

func setEnvStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}

func setEnvBool(key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		log.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

func setEnvU64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// getCurrentFolderPath can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return filepath.Dir(file), nil
}
