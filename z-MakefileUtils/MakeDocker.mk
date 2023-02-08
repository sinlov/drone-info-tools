#ENV_INFO_BUILD_DOCKER_TAG = ${ENV_DIST_VERSION}
ENV_INFO_BUILD_DOCKER_TAG = latest
ENV_INFO_BUILD_DOCKER_FILE = Dockerfile
ENV_INFO_DOCKER_REPOSITORY=${ROOT_NAME}
ENV_INFO_DOCKER_OWNER=bridgewwater
# if set ENV_INFO_PRIVATE_DOCKER_REGISTRY= will push to hub.docker.com
# private docker registry use harbor must create project name as ${ENV_INFO_DOCKER_OWNER}
#ENV_INFO_PRIVATE_DOCKER_REGISTRY=harbor.xxx.com/
ENV_INFO_PRIVATE_DOCKER_REGISTRY=
ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE ?= ${ENV_INFO_DOCKER_OWNER}/${ENV_INFO_DOCKER_REPOSITORY}
ENV_INFO_BUILD_DOCKER_FROM_IMAGE ?= alpine:3.17

ENV_INFO_TEST_BUILD_DOCKER_FILE ?= Dockerfile.s6
ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE ?= golang:1.17.13-buster
ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER ?= "test-parent-${ENV_INFO_DOCKER_REPOSITORY}"
ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME ?= "test-${ENV_INFO_DOCKER_REPOSITORY}"

## change this for ip-v4 get
#ROOT_LOCAL_IP_V4_LINUX = $$(ifconfig enp8s0 | grep inet | grep -v inet6 | cut -d ':' -f2 | cut -d ' ' -f1)
#ROOT_LOCAL_IP_V4_DARWIN = $$(ifconfig en0 | grep inet | grep -v inet6 | cut -d ' ' -f2)
#
#localIPLinux:
#	@echo "=> now run as docker with linux"
#	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_LINUX)"
#
#localIPDarwin:
#	@echo "=> now run as docker with darwin"
#	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_DARWIN)"

dockerEnv:
	@echo "== docker env print start"
	@echo "INFO_REPOSITORY                     ${ENV_INFO_DOCKER_REPOSITORY}"
	@echo "INFO_OWNER                          ${ENV_INFO_DOCKER_OWNER}"
	@echo "INFO_BUILD_DOCKER_IMAGE_NAME        ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}"
	@echo "INFO_BUILD_DOCKER_TAG               ${ENV_INFO_BUILD_DOCKER_TAG}"
	@echo "INFO_BUILD_DOCKER_FILE              ${ENV_INFO_BUILD_DOCKER_FILE}"
	@echo "INFO_BUILD_FROM_IMAGE               ${ENV_INFO_BUILD_DOCKER_FROM_IMAGE}"
	@echo "INFO_PRIVATE_DOCKER_REGISTRY        ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}"
	@echo "REGISTRY tag as:                    ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"
	@echo ""
	@echo "INFO_TEST_BUILD_PARENT_IMAGE        ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}"
	@echo "INFO_TEST_BUILD_PARENT_CONTAINER    ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER}"
	@echo ""
	@echo "INFO_TEST_TAG_BUILD_CONTAINER_NAME  ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}"
	@echo ""
	@echo "== docker env print end"

dockerAllPull:
	docker pull ${ENV_INFO_BUILD_DOCKER_FROM_IMAGE}
	docker pull ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}

dockerCleanImages:
	(while :; do echo 'y'; sleep 3; done) | docker image prune

dockerCleanPruneAll:
	(while :; do echo 'y'; sleep 3; done) | docker container prune
	(while :; do echo 'y'; sleep 3; done) | docker image prune

dockerRunContainerParentBuild:
	@echo "run rm container image: ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}"
	# docker run -d --rm --name ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER} ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}
	docker run -d --rm --name ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER} ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE} tail -f /dev/null
	@echo ""
	@echo "-> run rm container name: ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER}"
	@echo "-> into container use: docker exec -it ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER} bash"

dockerRmContainerParentBuild:
	-docker rm -f ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER}

dockerPruneContainerParentBuild: dockerRmContainerParentBuild
	-docker rmi -f ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}

dockerTestBuildLatest:
	docker build --rm=true --tag ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG} --file ${ENV_INFO_TEST_BUILD_DOCKER_FILE} .

dockerTestRunLatest:
	docker image inspect --format='{{ .Created}}' ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	-docker run --rm --name ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} \
	-e RUN_MODE=dev \
	${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	# for inner check can use like this
	# docker run -it -d --entrypoint /bin/sh --name ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	-docker inspect --format='{{ .State.Status}}' ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}

dockerTestLogLatest:
	-docker logs ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}

dockerTestRmLatest:
	-docker rm -f ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}

dockerTestRmiLatest:
	-INFO_TEST_TAG_BUILD_CONTAINER_NAME=$(ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME) \
	INFO_BUILD_DOCKER_SOURCE_IMAGE=$(ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE) \
	ROOT_DOCKER_IMAGE_TAG=$(ENV_INFO_BUILD_DOCKER_TAG) \
	INFO_BUILD_DOCKER_TAG=${ENV_INFO_BUILD_DOCKER_TAG} \
	docker rmi -f ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}

dockerTestRestartLatest: dockerTestRmLatest dockerTestRmiLatest dockerTestBuildLatest dockerTestRunLatest
	@echo "restart ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

dockerTestStopLatest: dockerTestRmLatest dockerTestRmiLatest
	@echo "stop and remove ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

dockerTestPruneLatest: dockerTestStopLatest
	@echo "prune and remove ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

dockerRmiBuild:
	-docker rmi -f ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	-docker rmi -f ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}

dockerBuild:
	docker build --rm=true --tag ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG} --file ${ENV_INFO_BUILD_DOCKER_FILE} .

dockerTag:
	docker tag ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG} ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}

dockerBeforePush: dockerRmiBuild dockerBuild dockerTag
	@echo "===== then now can push to ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}"

dockerPushBuild: dockerBeforePush
	docker push ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	@echo "=> push ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

helpDocker:
	@echo "=== this make file can include MakeDocker.mk then use"
	@echo "- must has file: [ ${ENV_INFO_BUILD_DOCKER_FILE} ${ENV_INFO_TEST_BUILD_DOCKER_FILE}" ]
	@echo "- then change tag as:                       INFO_BUILD_DOCKER_TAG"
	@echo "- then change repository as:                INFO_REPOSITORY"
	@echo "- then change owner as:                     INFO_OWNER"
	@echo "- then change private docker repository as: INFO_PRIVATE_DOCKER_REGISTRY"
	@echo "- then change build parent image as:        INFO_TEST_BUILD_PARENT_IMAGE"
	@echo "- then change build image as:               INFO_BUILD_FROM_IMAGE"
	@echo "- check by task"
	@echo "$$ make dockerEnv"
	@echo ""
	@echo "- first use can pull images"
	@echo "$$ make dockerAllPull"
	@echo ""
	@echo "- then use to show how to build docker parent image"
	@echo "$$ make dockerRunContainerParentBuild"
	@echo "- and prune resource at parent image"
	@echo "$$ make dockerPruneContainerParentBuild"
	@echo ""
	@echo "- test run container use ./${ENV_INFO_TEST_BUILD_DOCKER_FILE}"
	@echo "$$ make dockerTestRestartLatest"
	@echo "- prune test container and image"
	@echo "$$ make dockerTestPruneLatest"
	@echo ""
	@echo "- build and tag as use ./${ENV_INFO_BUILD_DOCKER_FILE}"
	@echo "$$ make dockerBeforePush"
	@echo "- prune build"
	@echo "$$ make dockerRmiBuild"
	@echo "- then final push use"
	@echo "$$ make dockerPushBuild"