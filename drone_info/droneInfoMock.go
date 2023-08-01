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
	mockEnvDroneRemoteProto       = "https"
	mockEnvDroneRemoteHost        = "github.com"

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
	mockEnvDroneSourceBranch       = ""
	mockEnvDroneTargetBranch       = ""
	mockEnvDroneBuildEvent         = "push"
	mockEnvDroneBuildStatusSuccess = "success"
	mockEnvDroneBuildStatusFailure = "failure"

	mockEnvDroneCommitMessage = "mock message commit\nmore line\nand more line\r\n"
	mockEnvDroneCommitSha     = "f5513d4e409730c28ee51ba85d01031d1e3a4eba"

	mockEnvFailedStages = "build,test"
	mockEnvFailedSteps  = "backend,frontend"
)

// MockDroneInfo
// status DroneBuildStatusSuccess DroneBuildStatusFailure
// this mock will use mockEnvDroneBranch as heads by MockDroneInfoRefs
//
// notes: this method mock time use time.UTC
func MockDroneInfo(status string) *Drone {
	droneInfoRefs, _ := MockDroneInfoRefs(status, fmt.Sprintf("refs/heads/%s", mockEnvDroneBranch))
	return droneInfoRefs
}

// MockDroneInfoRefs
//
//	host use mockEnvDroneSystemHost
//	status DroneBuildStatusSuccess DroneBuildStatusFailure
//	refs by: git show-ref --head --dereference
//
// notes: this method mock time use time.UTC
//
// @doc https://git-scm.com/docs/git-show-ref
// like refs/heads/master refs/remotes/* refs/pull/* refs/tags/v1.0.0
func MockDroneInfoRefs(status string, refs string) (*Drone, error) {
	return MockDroneInfoDroneSystemRefs(mockEnvDroneSystemProto, mockEnvDroneSystemHost, mockEnvDroneSystemHostName, status, refs)
}

// MockDroneInfoDroneSystemRefs
//
// droneHost like drone.xxx.com:80
//
// droneHostName like drone.xxx.com
//
// status DroneBuildStatusSuccess DroneBuildStatusFailure and default is DroneBuildStatusSuccess
//
// refs by: git show-ref --head --dereference
//
// notes: this method mock time use time.UTC
//
// @doc https://git-scm.com/docs/git-show-ref
// like refs/heads/master refs/remotes/* refs/pull/* refs/tags/v1.0.0
func MockDroneInfoDroneSystemRefs(
	droneProto,
	droneHost,
	droneHostName,
	status string,
	refs string,
) (*Drone, error) {
	return MockDroneInfoByRefsAndNumber(
		droneProto,
		droneHost,
		droneHostName,
		mockEnvDroneRemoteProto,
		mockEnvDroneRemoteHost,
		mockEnvDroneRepoOwner,
		mockEnvDroneRepo,
		status,
		refs,
		mockEnvDroneBuildNumber,
	)
}

// MockDroneInfoByRefsAndNumber
//
// droneHost like drone.xxx.com:80
//
// droneHostName like drone.xxx.com
//
// droneRemoteProto like http or https
//
// droneRemoteHost like github.com
//
// owner like sinlov
//
// repoName like drone-file-browser-plugin
//
// status DroneBuildStatusSuccess DroneBuildStatusFailure and default is DroneBuildStatusSuccess
//
// refs by: git show-ref --head --dereference
//
// notes: this method mock time use time.UTC
//
// @doc https://git-scm.com/docs/git-show-ref
// like refs/heads/master refs/remotes/* refs/pull/* refs/tags/v1.0.0
func MockDroneInfoByRefsAndNumber(
	droneProto,
	droneHost,
	droneHostName string,
	droneRemoteProto,
	droneRemoteHost,
	owner,
	repoName string,
	status string,
	refs string,
	buildNumber uint64,
) (*Drone, error) {

	if droneProto != "http" && droneProto != "https" {
		return nil, fmt.Errorf("droneProto only support http or https, now is: %s", droneProto)
	}

	if droneRemoteProto != "http" && droneRemoteProto != "https" {
		return nil, fmt.Errorf("droneRemoteProto only support http or https, now is: %s", droneRemoteProto)
	}

	if droneHost == "" {
		return nil, fmt.Errorf("droneHost not support nil")
	}
	if droneHostName == "" {
		return nil, fmt.Errorf("droneHostName not support nil")
	}

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

	email := mockEnvDroneCommitAuthorEmail
	repoUrl := fmt.Sprintf("%s://%s/%s/%s", droneRemoteProto, droneRemoteHost, owner, repoName)
	repoHttpUrl := fmt.Sprintf("%s://%s/%s/%s.git", droneRemoteProto, droneRemoteHost, owner, repoName)
	repoSshUrl := fmt.Sprintf("git@%s:%s/%s.git", droneRemoteHost, owner, repoName)
	repoHost := ""
	repoHostName := ""
	parse, err := url.Parse(repoHttpUrl)
	if err == nil {
		repoHost = parse.Host
		repoHostName = parse.Hostname()
	}
	stageStartT := mockEnvDroneStageStarted
	stageStartTime := time.Unix(int64(stageStartT), 0).In(time.UTC).Format(DroneTimeFormatDefault)
	stageFinishedT := mockEnvDroneStageFinished
	stageFinishedTime := time.Unix(int64(stageStartT), 0).In(time.UTC).Format(DroneTimeFormatDefault)
	commitSHA := mockEnvDroneCommitSha
	droneBaseUrl := fmt.Sprintf("%s://%s", droneProto, droneHost)

	var drone = Drone{
		//  repo info
		Repo: Repo{
			Link:      repoUrl,
			ShortName: repoName,
			GroupName: owner,
			FullName:  fmt.Sprintf("%s/%s", owner, repoName),
			OwnerName: owner,
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
			SourceBranch: mockEnvDroneSourceBranch,
			Link:         fmt.Sprintf("%s/%s/%s/%d", droneBaseUrl, owner, repoName, buildNumber),
			Event:        mockEnvDroneBuildEvent,
			StartAt:      mockEnvDroneBuildStarted,
			FinishedAt:   mockEnvDroneBuildFinished,
			PR:           "",
			DeployTo:     "",
		},
		Commit: Commit{
			Link:    fmt.Sprintf("%s/commit/%s", repoUrl, commitSHA),
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
			Host:     droneHost,
			HostName: droneHostName,
			Proto:    droneProto,
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

	if drone.Build.Status != DroneBuildStatusSuccess {
		drone.Build.FailedStages = mockEnvFailedStages
		drone.Build.FailedSteps = mockEnvFailedSteps
	}

	return &drone, nil
}

func MockDroneInfoEnvFull(debug bool) {
	setEnvBool(EnvKeyPluginDebug, debug)

	workspace, _ := getCurrentFolderPath()

	owner := mockEnvDroneRepoOwner
	email := mockEnvDroneCommitAuthorEmail
	repoName := mockEnvDroneRepo
	repoUrl := fmt.Sprintf("%s://%s/%s/%s", mockEnvDroneRemoteProto, mockEnvDroneRemoteHost, owner, repoName)
	repoHttpUrl := fmt.Sprintf("%s://%s/%s/%s.git", mockEnvDroneRemoteProto, mockEnvDroneRemoteHost, owner, repoName)
	repoSshUrl := fmt.Sprintf("git@%s:%s/%s.git", mockEnvDroneRemoteHost, owner, repoName)

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
	setEnvStr(EnvDroneSourceBranch, mockEnvDroneSourceBranch)
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
	envDebug, find := os.LookupEnv(EnvKeyPluginDebug)
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
