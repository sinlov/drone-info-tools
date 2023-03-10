package drone_info

const (
	DroneBuildStatusSuccess = "success"
	DroneBuildStatusFailure = "failure"
	DroneBuildStatusError   = "error"
	DroneBuildStatusKilled  = "killed"

	// DroneTimeFormatDefault
	// default time format for Stage.StartedTime and Stage.FinishedTime
	DroneTimeFormatDefault = "2006-01-02-03-04-05"

	// EnvDroneCommitAuthorName
	// Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-name/
	EnvDroneCommitAuthorName = "DRONE_COMMIT_AUTHOR_NAME"
	// EnvDroneCommitAuthorEmail
	// Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-email/
	EnvDroneCommitAuthorEmail = "DRONE_COMMIT_AUTHOR_EMAIL"
	// EnvDroneCommitAuthor
	// Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username).
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author/
	EnvDroneCommitAuthor = "DRONE_COMMIT_AUTHOR"
	// EnvDroneCommitAuthorAvatar
	// Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub).
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-avatar/
	EnvDroneCommitAuthorAvatar = "DRONE_COMMIT_AUTHOR_AVATAR"
	EnvDroneCommitBranch       = "DRONE_COMMIT_BRANCH"
	EnvDroneCommitLink         = "DRONE_COMMIT_LINK"
	EnvDroneCommitMessage      = "DRONE_COMMIT_MESSAGE"
	EnvDroneCommitSha          = "DRONE_COMMIT_SHA"
	// EnvDroneCommitRef
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-ref/
	EnvDroneCommitRef = "DRONE_COMMIT_REF"
	// EnvDroneRepo
	// most is EnvDroneRepoNamespace / EnvDroneRepoName
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo/
	EnvDroneRepo     = "DRONE_REPO"
	EnvDroneRepoName = "DRONE_REPO_NAME"
	// EnvDroneRepoLink
	// by env:DRONE_REPO_LINK
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-link/
	EnvDroneRepoLink = "DRONE_REPO_LINK"
	// EnvDroneRepoNamespace
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-namespace/
	EnvDroneRepoNamespace = "DRONE_REPO_NAMESPACE"
	// EnvDroneGitHttpUrl
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-git-http-url/
	EnvDroneGitHttpUrl = "DRONE_GIT_HTTP_URL"
	// EnvDroneGitSshUrl
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-git-ssh-url/
	EnvDroneGitSshUrl = "DRONE_GIT_SSH_URL"
	EnvDroneRemoteUrl = "DRONE_REMOTE_URL"
	EnvDroneRepoOwner = "DRONE_REPO_OWNER"
	// EnvDroneRepoScm
	// must is: git hg
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-scm/
	EnvDroneRepoScm = "DRONE_REPO_SCM"

	// EnvDroneBuildWorkSpace
	// drone???s working directory for a pipeline
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-workspace/
	EnvDroneBuildWorkSpace = "DRONE_WORKSPACE"
	EnvDroneBuildStatus    = "DRONE_BUILD_STATUS"
	EnvDroneBuildNumber    = "DRONE_BUILD_NUMBER"
	EnvDroneBuildLink      = "DRONE_BUILD_LINK"
	EnvDroneBuildEvent     = "DRONE_BUILD_EVENT"
	EnvDroneBuildStarted   = "DRONE_BUILD_STARTED"
	EnvDroneBuildFinished  = "DRONE_BUILD_FINISHED"

	// EnvDroneTag
	// by env: DRONE_TAG
	// Provides the tag for the current running build. This value is only populated for tag events and promotion events that are derived from tags.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-tag/
	EnvDroneTag = "DRONE_TAG"
	// EnvDroneBranch
	// by env: DRONE_BRANCH
	// Provides the target branch for the push or pull request. This value may be empty for tag events.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-branch/
	EnvDroneBranch = "DRONE_BRANCH"
	// EnvDroneTargetBranch
	// This environment variable can be used in conjunction with the source branch variable to get the pull request base and head branch.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-target-branch/
	EnvDroneTargetBranch = "DRONE_TARGET_BRANCH"
	EnvDronePR           = "DRONE_PULL_REQUEST"
	EnvDroneDeployTo     = "DRONE_DEPLOY_TO"

	// EnvDroneStageStarted
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-started/
	EnvDroneStageStarted = "DRONE_STAGE_STARTED"
	// EnvDroneStageFinished
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-finished/
	EnvDroneStageFinished = "DRONE_STAGE_FINISHED"

	// EnvDroneStageMachine
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-machine/
	EnvDroneStageMachine = "DRONE_STAGE_MACHINE"
	// EnvDroneStageOs
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-os/
	EnvDroneStageOs = "DRONE_STAGE_OS"
	// EnvDroneStageArch
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-arch/
	EnvDroneStageArch = "DRONE_STAGE_ARCH"
	// EnvDroneStageVariant
	// Provides the target operating architecture variable for the current build stage. This variable is optional and is only available for arm architectures.
	// most is: "", or v7
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-variant/
	EnvDroneStageVariant = "DRONE_STAGE_VARIANT"
	// EnvDroneStageType
	// most use: docker
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-type/
	EnvDroneStageType = "DRONE_STAGE_TYPE"
	// EnvDroneStageKind
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-kind/
	EnvDroneStageKind = "DRONE_STAGE_KIND"
	// EnvDroneStageName
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-name/
	EnvDroneStageName = "DRONE_STAGE_NAME"

	// EnvDroneFailedStages
	// by env:DRONE_FAILED_STAGES
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-stages/
	EnvDroneFailedStages = "DRONE_FAILED_STAGES"
	// EnvDroneFailedSteps
	// by env:DRONE_FAILED_STEPS
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-steps/
	EnvDroneFailedSteps = "DRONE_FAILED_STEPS"

	// EnvDroneSystemVersion
	// Provides the version of the Drone server.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-version/
	EnvDroneSystemVersion = "DRONE_SYSTEM_VERSION"
	// EnvDroneSystemHost
	// Provides the hostname used by the Drone server. This can be combined with the protocol to construct to the server url.
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-host/
	EnvDroneSystemHost = "DRONE_SYSTEM_HOST"
	// EnvDroneSystemHostName
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-hostname/
	EnvDroneSystemHostName = "DRONE_SYSTEM_HOSTNAME"
	// EnvDroneSystemProto
	// https or http
	// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-proto/
	EnvDroneSystemProto = "DRONE_SYSTEM_PROTO"
)

var (
	// droneBuildStatusStatusOptSupport
	droneBuildStatusStatusOptSupport []string
)

func DroneBuildStatusStatusOptSupport() []string {
	if droneBuildStatusStatusOptSupport == nil {
		droneBuildStatusStatusOptSupport = []string{
			"",
			DroneBuildStatusSuccess,
			DroneBuildStatusFailure,
			DroneBuildStatusError,
			DroneBuildStatusKilled,
		}
	}
	return droneBuildStatusStatusOptSupport
}

type (
	// Repo repo base info
	Repo struct {
		// Link
		// by env:DRONE_REPO_LINK
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-link/
		Link string
		// ShortName
		// by env:DRONE_REPO_NAME
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-name/
		ShortName string //  short name
		// GroupName
		// by env:DRONE_REPO_NAMESPACE
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-namespace/
		GroupName string //  group name
		// FullName
		// by env:DRONE_REPO
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo/
		FullName string //  repository full name
		// OwnerName
		// by env:DRONE_REPO_OWNER
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-owner/
		OwnerName string //  repo owner
		// Scm
		// by env:DRONE_REPO_SCM
		// must is: git hg
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-scm/
		Scm string
		// RemoteURL
		// by env:DRONE_REMOTE_URL
		// Provides the git+https url that should be used to clone the repository
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-remote-url/
		RemoteURL string //  repo remote url
		// HttpUrl
		// by env:DRONE_GIT_HTTP_URL
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-git-http-url/
		HttpUrl string
		// SshUrl
		// by env:DRONE_GIT_SSH_URL
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-git-ssh-url/
		SshUrl string
		// Host
		// this from HttpUrl host
		Host string
		// HostName
		// this from HttpUrl hostname
		HostName string
	}

	// Build info
	Build struct {
		// WorkSpace
		// drone???s working directory for a pipeline
		// by env:DRONE_WORKSPACE
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-workspace/
		WorkSpace string
		Status    string //  providers the current build status
		Number    uint64 //  providers the current build number
		// Branch
		// by env: DRONE_TAG
		// Provides the tag for the current running build. This value is only populated for tag events and promotion events that are derived from tags.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-tag/
		Tag string
		// Branch
		// by env: DRONE_BRANCH
		// Provides the target branch for the push or pull request. This value may be empty for tag events.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-branch/
		Branch string
		// TargetBranch
		// by env:DRONE_TARGET_BRANCH
		// This environment variable can be used in conjunction with the source branch variable to get the pull request base and head branch.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-target-branch/
		TargetBranch string
		// Link
		// by env:DRONE_BUILD_LINK
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-build-link/
		Link       string //  providers the current build link
		Event      string //  trigger event
		StartAt    uint64 //  build start at ( unix timestamp )
		FinishedAt uint64 //  build finish at ( unix timestamp )
		PR         string //  build pull request
		DeployTo   string //  build deploy to
		// FailedStages
		// by env:DRONE_FAILED_STAGES
		// Provides a comma-separate list of failed pipeline stages for the current running build.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-stages/
		FailedStages string
		// FailedSteps
		// by env:DRONE_FAILED_STEPS
		// Provides a comma-separate list of failed pipeline steps.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-steps/
		FailedSteps string
	}

	// Commit info
	Commit struct {
		// Link
		// by env:DRONE_COMMIT_LINK
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-link/
		Link    string //  providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		Branch  string //  providers the branch for the current commit
		Message string //  providers the commit message for the current build
		// Sha
		// by env:DRONE_COMMIT_SHA
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-sha/
		Sha string //  providers the commit sha for the current build
		// Ref
		// by env:DRONE_COMMIT_REF
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-ref/
		Ref    string //  commit ref
		Author CommitAuthor
	}

	// Stage drone stage env
	Stage struct {
		// StartedAt
		// by env:DRONE_STAGE_STARTED
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-started/
		StartedAt uint64
		// FinishedAt
		// by env:DRONE_STAGE_FINISHED
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-finished/
		// FinishedTime
		// form StartedAt
		StartedTime string
		FinishedAt  uint64
		// FinishedTime
		// form FinishedAt
		FinishedTime string
		// Machine
		// by env:DRONE_STAGE_MACHINE
		// Provides the name of the host machine on which the pipeline is currently running.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-machine/
		Machine string
		// Os
		// by env:DRONE_STAGE_OS
		// Provides the target operating system for the current build stag
		// List of all possible values: darwin dragonfly freebsd linux netbsd openbsd solaris windows
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-os/
		Os string
		// Arch
		// by env:DRONE_STAGE_ARCH
		// Provides the platform architecture for the current build stage.
		// List of all possible values: 386 amd64 arm64 arm
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-arch/
		Arch string
		// Variant
		// by env:DRONE_STAGE_VARIANT
		// Provides the target operating architecture variable for the current build stage. This variable is optional and is only available for arm architectures.
		// most is: "", or v7
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-variant/
		Variant string
		// Type
		// by env:DRONE_STAGE_TYPE
		// This value is sourced from the type attribute in the yaml configuration file.
		// most of is: docker
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-type/
		Type string
		// Kind
		// by env:DRONE_STAGE_KIND
		// most of is: pipeline
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-kind/
		Kind string
		// Name
		// by env:DRONE_STAGE_NAME
		// most of is build
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-name/
		Name string
	}

	// CommitAuthor commit author info
	CommitAuthor struct {
		// Username
		// by env:DRONE_COMMIT_AUTHOR_NAME
		// Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-name/
		Username string //  the author username for the current commit
		// Email
		// by env:DRONE_COMMIT_AUTHOR_EMAIL
		// Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-email/
		Email string //  providers the author email for the current commit
		// Name
		// by env:DRONE_COMMIT_AUTHOR
		// Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username).
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author/
		Name string //  providers the author name for the current commit
		// Avatar
		// by env:DRONE_COMMIT_AUTHOR_AVATAR
		// Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub).
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-avatar/
		Avatar string //  providers the author avatar for the current commit
	}

	DroneSystem struct {
		// Version
		// by env: DRONE_SYSTEM_VERSION
		// Provides the version of the Drone server.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-version/
		Version string
		// Host
		// by env:DRONE_SYSTEM_HOST
		// Provides the host used by the Drone server. This can be combined with the protocol to construct to the server url.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-host/
		Host string
		// HostName
		// by env:DRONE_SYSTEM_HOSTNAME
		// Provides the hostname used by the Drone server. This can be combined with the protocol to construct to the server url.
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-hostname/
		HostName string
		// Proto
		// by env:DRONE_SYSTEM_PROTO
		// Provides the protocol used by the Drone server. This can be combined with the hostname to construct to the server url.
		// https or http
		// @doc https://docs.drone.io/pipeline/environment/reference/drone-system-proto/
		Proto string
	}

	// Drone drone info
	Drone struct {
		Repo        Repo
		Build       Build
		Commit      Commit
		Stage       Stage
		DroneSystem DroneSystem
	}
)
