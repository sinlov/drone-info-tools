package drone_info

const (
	NameCliStepsDebug = "build.debug"

	// EnvKeyPluginDebug
	//  Provides the plugin debug flag. This value is true when the plugin is open debug mode
	EnvKeyPluginDebug = "PLUGIN_DEBUG"

	DroneBuildStatusSuccess = "success"
	DroneBuildStatusFailure = "failure"
	DroneBuildStatusError   = "error"
	DroneBuildStatusKilled  = "killed"

	// DroneTimeFormatDefault
	//  default time format for Stage.StartedTime and Stage.FinishedTime
	DroneTimeFormatDefault = "2006-01-02-03-04-05"

	// EnvDroneCommitAuthorName
	//  Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-name/

	EnvDroneCommitAuthorName = "DRONE_COMMIT_AUTHOR_NAME"
	// EnvDroneCommitAuthorEmail
	//  Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-email/
	EnvDroneCommitAuthorEmail = "DRONE_COMMIT_AUTHOR_EMAIL"

	// EnvDroneCommitAuthor
	//  Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username).
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author/
	EnvDroneCommitAuthor = "DRONE_COMMIT_AUTHOR"

	// EnvDroneCommitAuthorAvatar
	//  Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub).
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-avatar/
	EnvDroneCommitAuthorAvatar = "DRONE_COMMIT_AUTHOR_AVATAR"

	// EnvDroneCommitBranch
	//
	//  Provides the target branch for the push or pull request. This value may be empty for tag events.
	//
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-branch/
	EnvDroneCommitBranch  = "DRONE_COMMIT_BRANCH"
	EnvDroneCommitLink    = "DRONE_COMMIT_LINK"
	EnvDroneCommitMessage = "DRONE_COMMIT_MESSAGE"
	EnvDroneCommitSha     = "DRONE_COMMIT_SHA"

	// EnvDroneCommitRef
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-ref/
	EnvDroneCommitRef = "DRONE_COMMIT_REF"

	// EnvDroneRepo
	//  most is EnvDroneRepoNamespace / EnvDroneRepoName
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo/
	EnvDroneRepo     = "DRONE_REPO"
	EnvDroneRepoName = "DRONE_REPO_NAME"

	// EnvDroneRepoLink
	//  by env:DRONE_REPO_LINK
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-link/
	EnvDroneRepoLink = "DRONE_REPO_LINK"

	// EnvDroneRepoNamespace
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-namespace/
	EnvDroneRepoNamespace = "DRONE_REPO_NAMESPACE"

	// EnvDroneGitHttpUrl
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-git-http-url/
	EnvDroneGitHttpUrl = "DRONE_GIT_HTTP_URL"

	// EnvDroneGitSshUrl
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-git-ssh-url/
	EnvDroneGitSshUrl = "DRONE_GIT_SSH_URL"
	EnvDroneRemoteUrl = "DRONE_REMOTE_URL"
	EnvDroneRepoOwner = "DRONE_REPO_OWNER"

	// EnvDroneRepoScm
	//  must is: git hg
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-scm/
	EnvDroneRepoScm = "DRONE_REPO_SCM"

	// EnvDroneBuildWorkSpace
	//  drone’s working directory for a pipeline
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-workspace/
	EnvDroneBuildWorkSpace = "DRONE_WORKSPACE"

	// EnvDroneBuildDebug
	//	by env: DRONE_BUILD_DEBUG
	//	Provides the build debug flag. This value is true when the build is open debug mode
	//	this not has official doc, just use build steps debug will change this env to true
	EnvDroneBuildDebug = "DRONE_BUILD_DEBUG"

	// EnvDroneBuildTrigger
	//	by env: DRONE_BUILD_TRIGGER
	//	webhook (@hook)
	//	The source of the trigger for the build.
	//	DRONE_BUILD_TRIGGER=root
	//	DRONE_BUILD_TRIGGER=f.bar
	//	DRONE_BUILD_TRIGGER=@hook
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-trigger/
	EnvDroneBuildTrigger = "DRONE_BUILD_TRIGGER"

	// EnvDroneBuildStatus
	//	by env: DRONE_BUILD_STATUS
	//	success or failure
	//	Provides the status for the current running build. If build pipelines and build steps are passing, the build status defaults to success.
	//
	//	Please note this is point in time snapshot. This value may not accurately reflect the overall build status when multiple pipelines are running in parallel.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-status/
	EnvDroneBuildStatus = "DRONE_BUILD_STATUS"

	// EnvDroneBuildNumber
	//  by env: DRONE_BUILD_NUMBER
	//  Provides the build number for the current running build. Please note that this variable is not available for substitution or template injection.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-number/
	EnvDroneBuildNumber = "DRONE_BUILD_NUMBER"

	// EnvDroneTag
	//  by env: DRONE_TAG
	//  Provides the tag for the current running build. This value is only populated for tag events and promotion events that are derived from tags.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-tag/
	EnvDroneTag = "DRONE_TAG"

	// EnvDroneBuildCreated
	//  by env: DRONE_BUILD_CREATED
	//  Provides the unix timestamp for when the build object was created by the system.
	//  build created at ( unix timestamp )
	//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-created/
	EnvDroneBuildCreated = "DRONE_BUILD_CREATED"

	// EnvDroneBuildStarted
	//  by env: DRONE_BUILD_STARTED
	//  Provides the unix timestamp for when the build was started by the system.
	//  build created at ( unix timestamp )
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-started/
	EnvDroneBuildStarted = "DRONE_BUILD_STARTED"

	// EnvDroneBuildFinished
	//  by env: DRONE_BUILD_FINISHED
	//  Provides the unix timestamp for when the build is finished. A running build cannot have a finish timestamp, therefore, the system always sets this value to the current timestamp.
	//  build finished at ( unix timestamp )
	//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-finished/
	EnvDroneBuildFinished = "DRONE_BUILD_FINISHED"

	// EnvDroneBranch
	//  by env: DRONE_BRANCH
	//  Provides the target branch for the push or pull request. This value may be empty for tag events.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-branch/
	EnvDroneBranch = "DRONE_BRANCH"

	// EnvDroneRepoBranch
	//  Provides the default repository branch for the current running build
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-branch/
	EnvDroneRepoBranch = "DRONE_REPO_BRANCH"

	// EnvDroneSourceBranch
	//  Provides the source branch for the pull request. This value may be empty for certain source control management providers.
	//
	//  This environment variable can be used in conjunction with the target branch variable to get the pull request base and head branch.
	//
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-source-branch/
	EnvDroneSourceBranch = "DRONE_SOURCE_BRANCH"

	// EnvDroneTargetBranch
	//
	//  Provides the target branch for the push or pull request. This value may be empty for tag events.
	//
	//  This environment variable can be used in conjunction with the source branch variable to get the pull request base and head branch.
	//
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-target-branch/
	EnvDroneTargetBranch = "DRONE_TARGET_BRANCH"

	// EnvDroneBuildLink
	//  by env: DRONE_BUILD_LINK
	//  Provides a deep link the Drone build results. Please note that this variable is not available for substitution or template injection.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-link/
	EnvDroneBuildLink = "DRONE_BUILD_LINK"

	// EnvDroneBuildEvent
	//  by env:DRONE_BUILD_EVENT
	//
	//	DRONE_BUILD_EVENT=push
	//	DRONE_BUILD_EVENT=pull_request
	//	DRONE_BUILD_EVENT=promote
	//	DRONE_BUILD_EVENT=rollback
	//	DRONE_BUILD_EVENT=tag
	//	DRONE_BUILD_EVENT=cron
	//	DRONE_BUILD_EVENT=custom
	//
	//  Provides the event that triggered the pipeline execution.
	//
	//  This value will be push, pull_request, tag, promote, cron, custom.
	//
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-event/
	EnvDroneBuildEvent = "DRONE_BUILD_EVENT"

	// EnvDroneStageStarted
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-started/
	EnvDroneStageStarted = "DRONE_STAGE_STARTED"

	// EnvDroneStageFinished
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-finished/
	EnvDroneStageFinished = "DRONE_STAGE_FINISHED"

	// EnvDronePR
	//	by env: DRONE_PULL_REQUEST
	//  Provides the pull request number for the current running build. If the build is not a pull request the variable is empty.
	//
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-pull-request/
	EnvDronePR = "DRONE_PULL_REQUEST"

	// EnvDroneBuildAction
	//	by env: DRONE_BUILD_ACTION
	//	DRONE_BUILD_ACTION=sync
	//	DRONE_BUILD_ACTION=open
	//	Provides the action that triggered the pipeline execution. Use this value to differentiate between the pull request being opened vs synchronized.
	//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-action/
	EnvDroneBuildAction = "DRONE_BUILD_ACTION"

	// EnvDroneDeployTo
	//	by env: DRONE_DEPLOY_TO
	//	Provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-deploy-to/
	EnvDroneDeployTo = "DRONE_DEPLOY_TO"

	// EnvDroneFailedStages
	//  by env:DRONE_FAILED_STAGES
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-stages/
	EnvDroneFailedStages = "DRONE_FAILED_STAGES"

	// EnvDroneFailedSteps
	//  by env:DRONE_FAILED_STEPS
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-steps/
	EnvDroneFailedSteps = "DRONE_FAILED_STEPS"

	// EnvDroneStageMachine
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-machine/
	EnvDroneStageMachine = "DRONE_STAGE_MACHINE"

	// EnvDroneStageOs
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-os/
	EnvDroneStageOs = "DRONE_STAGE_OS"

	// EnvDroneStageArch
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-arch/
	EnvDroneStageArch = "DRONE_STAGE_ARCH"

	// EnvDroneStageVariant
	//  Provides the target operating architecture variable for the current build stage. This variable is optional and is only available for arm architectures.
	//  most is: "", or v7
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-variant/
	EnvDroneStageVariant = "DRONE_STAGE_VARIANT"

	// EnvDroneStageType
	//  most use: docker
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-type/
	EnvDroneStageType = "DRONE_STAGE_TYPE"

	// EnvDroneStageKind
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-kind/
	EnvDroneStageKind = "DRONE_STAGE_KIND"

	// EnvDroneStageName
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-name/
	EnvDroneStageName = "DRONE_STAGE_NAME"

	// EnvDroneSystemVersion
	//  Provides the version of the Drone server.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-version/
	EnvDroneSystemVersion = "DRONE_SYSTEM_VERSION"

	// EnvDroneSystemHost
	//  Provides the hostname used by the Drone server. This can be combined with the protocol to construct to the server url.
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-host/
	EnvDroneSystemHost = "DRONE_SYSTEM_HOST"

	// EnvDroneSystemHostName
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-hostname/
	EnvDroneSystemHostName = "DRONE_SYSTEM_HOSTNAME"

	// EnvDroneSystemProto
	//  https or http
	//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-proto/
	EnvDroneSystemProto = "DRONE_SYSTEM_PROTO"

	// EnvDroneSystemDebug
	//	Optional boolean value. Enables debug level logging
	//	by env:DRONE_DEBUG
	//	notice this not build debug in drone, just print debug log
	//  @doc https://docs.drone.io/runner/ssh/configuration/reference/drone-debug/
	EnvDroneSystemDebug = "DRONE_DEBUG"
)

var (
	//  droneBuildStatusStatusOptSupport
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
		//  Link
		//  by env:DRONE_REPO_LINK
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-link/
		Link string
		//  ShortName
		//  by env:DRONE_REPO_NAME
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-name/
		ShortName string //   short name
		//  GroupName
		//  by env:DRONE_REPO_NAMESPACE
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-namespace/
		GroupName string //   group name
		//  FullName
		//  by env:DRONE_REPO
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo/
		FullName string //   repository full name
		//  OwnerName
		//  by env:DRONE_REPO_OWNER
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-owner/
		OwnerName string //   repo owner
		//  Scm
		//  by env:DRONE_REPO_SCM
		//  must is: git hg
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-scm/
		Scm string
		//  RemoteURL
		//  by env:DRONE_REMOTE_URL
		//  Provides the git+https url that should be used to clone the repository
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-remote-url/
		RemoteURL string //   repo remote url
		//  HttpUrl
		//  by env:DRONE_GIT_HTTP_URL
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-git-http-url/
		HttpUrl string
		//  SshUrl
		//  by env:DRONE_GIT_SSH_URL
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-git-ssh-url/
		SshUrl string
		//  Host
		//  this from HttpUrl host
		Host string
		//  HostName
		//  this from HttpUrl hostname
		HostName string
	}

	// Build info
	Build struct {
		//  WorkSpace
		//  drone’s working directory for a pipeline
		//  by env:DRONE_WORKSPACE
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-workspace/
		WorkSpace string

		//  BuildDebug
		//	by env: DRONE_BUILD_DEBUG
		//	Provides the build debug flag. This value is true when the build is open debug mode
		//	this not has official doc, just use build steps debug will change this env to true
		BuildDebug bool

		//  Trigger
		//	by env: DRONE_BUILD_TRIGGER
		//	webhook (@hook)
		//	The source of the trigger for the build.
		//	DRONE_BUILD_TRIGGER=root
		//	DRONE_BUILD_TRIGGER=f.bar
		//	DRONE_BUILD_TRIGGER=@hook
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-trigger/
		Trigger string

		//  Status
		//	by env: DRONE_BUILD_STATUS
		//	success or failure
		//	Provides the status for the current running build. If build pipelines and build steps are passing, the build status defaults to success.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-status/
		Status string

		//  Number
		//	by env: DRONE_BUILD_NUMBER
		//	Provides the build number for the current running build. Please note that this variable is not available for substitution or template injection.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-number/
		Number uint64

		//  Branch
		//  by env: DRONE_TAG
		//  Provides the tag for the current running build. This value is only populated for tag events and promotion events that are derived from tags.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-tag/
		Tag string

		//  Branch
		//  by env: DRONE_BRANCH
		//  Provides the target branch for the push or pull request. This value may be empty for tag events.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-branch/
		Branch string

		//  RepoBranch
		//	by env:DRONE_REPO_BRANCH
		//	Provides the default repository branch for the current running build
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-repo-branch/
		RepoBranch string

		//  SourceBranch
		//  by env:DRONE_SOURCE_BRANCH
		//
		//  Provides the source branch for the pull request. This value may be empty for certain source control management providers.
		//
		//  This environment variable can be used in conjunction with the target branch variable to get the pull request base and head branch.
		//
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-source-branch/
		SourceBranch string

		//  TargetBranch
		//
		//  Provides the target branch for the push or pull request. This value may be empty for tag events.
		//
		//  This environment variable can be used in conjunction with the source branch variable to get the pull request base and head branch.
		//
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-target-branch/
		TargetBranch string

		//  Link
		//	by env: DRONE_BUILD_LINK
		//	Provides a deep link the Drone build results. Please note that this variable is not available for substitution or template injection.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-link/
		Link string //   providers the current build link

		//  Event
		//	by env:DRONE_BUILD_EVENT
		//
		//	DRONE_BUILD_EVENT=push
		//	DRONE_BUILD_EVENT=pull_request
		//	DRONE_BUILD_EVENT=promote
		//	DRONE_BUILD_EVENT=rollback
		//	DRONE_BUILD_EVENT=tag
		//	DRONE_BUILD_EVENT=cron
		//	DRONE_BUILD_EVENT=custom
		//
		//  Provides the event that triggered the pipeline execution.
		//
		//  This value will be push, pull_request, tag, promote, cron, custom.
		//
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-build-event/
		Event string //   trigger event

		//  CreatedAt
		//	by env: DRONE_BUILD_CREATED
		//	Provides the unix timestamp for when the build object was created by the system.
		//	build created at ( unix timestamp )
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-created/
		CreatedAt uint64

		//  CreatedAtT
		//	by env: DRONE_BUILD_CREATED
		//	Provides the unix timestamp for when the build object was created by the system.
		//	build created at ( unix timestamp )
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-created/
		CreatedAtT string

		//  StartedAt
		//	by env: DRONE_BUILD_STARTED
		//	Provides the unix timestamp for when the build was started by the system.
		//	build created at ( unix timestamp )
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-started/
		StartAt uint64

		//  StartedAtT
		//	by env: DRONE_BUILD_STARTED
		//	Provides the unix timestamp for when the build was started by the system.
		//	build created at ( unix timestamp )
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-started/
		StartAtT string

		//  FinishedAt
		//	by env: DRONE_BUILD_FINISHED
		//	Provides the unix timestamp for when the build is finished. A running build cannot have a finish timestamp, therefore, the system always sets this value to the current timestamp.
		//	build finished at ( unix timestamp )
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-finished/
		FinishedAt uint64

		//  FinishedAtT
		//	by env: DRONE_BUILD_FINISHED
		//	Provides the unix timestamp for when the build is finished. A running build cannot have a finish timestamp, therefore, the system always sets this value to the current timestamp.
		//	build finished at ( unix timestamp )
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-finished/
		FinishedAtT string

		//  TotalBuildTime
		//  by env: DRONE_BUILD_FINISHED - DRONE_BUILD_CREATED
		//  Provides the total build time in seconds. This value is calculated by subtracting the build created timestamp from the build finished timestamp.
		TotalBuildTime string

		//  PR
		//  by env:DRONE_PULL_REQUEST
		//
		//  Provides the pull request number for the current running build. If the build is not a pull request the variable is empty.
		//
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-pull-request/
		PR string

		//  PRAction
		//	by env: DRONE_BUILD_ACTION
		//	DRONE_BUILD_ACTION=sync
		//	DRONE_BUILD_ACTION=open
		//	Provides the action that triggered the pipeline execution. Use this value to differentiate between the pull request being opened vs synchronized.
		//   @doc https://docs.drone.io/pipeline/environment/reference/drone-build-action/
		PRAction string

		//  DeployTo
		//	by env: DRONE_DEPLOY_TO
		//	Provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-deploy-to/
		DeployTo string //   build deploy to

		//  FailedStages
		//  by env:DRONE_FAILED_STAGES
		//  Provides a comma-separate list of failed pipeline stages for the current running build.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-stages/
		FailedStages string

		// FailedSteps
		//  by env:DRONE_FAILED_STEPS
		//  Provides a comma-separate list of failed pipeline steps.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-failed-steps/
		FailedSteps string
	}

	// Commit info
	Commit struct {
		//  Link
		//  by env:DRONE_COMMIT_LINK
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-link/
		Link string //   providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		//  Branch
		//  by env:DRONE_COMMIT_BRANCH
		//
		//  Provides the target branch for the push or pull request. This value may be empty for tag events.
		//
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-branch/
		Branch  string //   providers the branch for the current commit
		Message string //   providers the commit message for the current build
		//  Sha
		//  by env:DRONE_COMMIT_SHA
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-sha/
		Sha string //   providers the commit sha for the current build
		//  Ref
		//  by env:DRONE_COMMIT_REF
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-ref/
		Ref    string //   commit ref
		Author CommitAuthor
	}

	// Stage
	//  drone stage env
	Stage struct {
		//  StartedAt
		//  by env:DRONE_STAGE_STARTED
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-started/
		StartedAt uint64

		//  StartedTime
		//  by env:DRONE_STAGE_STARTED
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-started/
		StartedTime string

		//  FinishedAt
		//  by env:DRONE_STAGE_FINISHED
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-finished/
		FinishedAt uint64

		//  FinishedTime
		//  form FinishedAt
		FinishedTime string

		//  TotalStageTime
		//   by env: DRONE_STAGE_FINISHED - DRONE_STAGE_STARTED format as human
		//   this is this stage build time, not full build time see Drone.Build TotalBuildTime
		TotalStageTime string

		//  Machine
		//  by env:DRONE_STAGE_MACHINE
		//  Provides the name of the host machine on which the pipeline is currently running.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-machine/
		Machine string
		//  Os
		//  by env:DRONE_STAGE_OS
		//  Provides the target operating system for the current build stag
		//  List of all possible values: darwin dragonfly freebsd linux netbsd openbsd solaris windows
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-os/
		Os string
		//  Arch
		//  by env:DRONE_STAGE_ARCH
		//  Provides the platform architecture for the current build stage.
		//  List of all possible values: 386 amd64 arm64 arm
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-arch/
		Arch string
		//  Variant
		//  by env:DRONE_STAGE_VARIANT
		//  Provides the target operating architecture variable for the current build stage. This variable is optional and is only available for arm architectures.
		//  most is: "", or v7
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-variant/
		Variant string
		//  Type
		//  by env:DRONE_STAGE_TYPE
		//  This value is sourced from the type attribute in the yaml configuration file.
		//  most of is: docker
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-type/
		Type string
		//  Kind
		//  by env:DRONE_STAGE_KIND
		//  most of is: pipeline
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-kind/
		Kind string
		//  Name
		//  by env:DRONE_STAGE_NAME
		//  most of is build
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-stage-name/
		Name string
	}

	// CommitAuthor commit author info
	CommitAuthor struct {
		//  Username
		//  by env:DRONE_COMMIT_AUTHOR_NAME
		//  Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-name/
		Username string //   the author username for the current commit
		//  Email
		//  by env:DRONE_COMMIT_AUTHOR_EMAIL
		//  Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-email/
		Email string //   providers the author email for the current commit
		//  Name
		//  by env:DRONE_COMMIT_AUTHOR
		//  Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username).
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author/
		Name string //   providers the author name for the current commit
		//  Avatar
		//  by env:DRONE_COMMIT_AUTHOR_AVATAR
		//  Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub).
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-commit-author-avatar/
		Avatar string //   providers the author avatar for the current commit
	}

	DroneSystem struct {
		//  Version
		//  by env: DRONE_SYSTEM_VERSION
		//  Provides the version of the Drone server.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-version/
		Version string
		//  Host
		//  by env:DRONE_SYSTEM_HOST
		//  Provides the host used by the Drone server. This can be combined with the protocol to construct to the server url.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-host/
		Host string
		//  HostName
		//  by env:DRONE_SYSTEM_HOSTNAME
		//  Provides the hostname used by the Drone server. This can be combined with the protocol to construct to the server url.
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-hostname/
		HostName string
		//  Proto
		//  by env:DRONE_SYSTEM_PROTO
		//  Provides the protocol used by the Drone server. This can be combined with the hostname to construct to the server url.
		//  https or http
		//  @doc https://docs.drone.io/pipeline/environment/reference/drone-system-proto/
		Proto string

		//  DroneSystemDebug
		//	Optional boolean value. Enables debug level logging
		//	by env:DRONE_DEBUG
		//	notice this not build debug in drone, just print debug log
		//  @doc https://docs.drone.io/runner/ssh/configuration/reference/drone-debug/
		DroneSystemDebug bool
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
