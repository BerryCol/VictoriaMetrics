# All these commands must run from repository root.

DOCKER_NAMESPACE := victoriametrics

ROOT_IMAGE ?= alpine:3.15.0
CERTS_IMAGE := alpine:3.15.0
GO_VERSION=1.17.6
GO_BUILDER_IMAGE := golang:$(GO_VERSION)-alpine
GO_BUILDER_IMAGE_ARM64 := golang:$(GO_VERSION)
BUILDER_IMAGE := local/builder:2.0.0-$(shell echo $(GO_BUILDER_IMAGE) | tr :/ __)
BUILDER_IMAGE_ARM64 := local/builder/arm:2.0.0-$(shell echo $(GO_BUILDER_IMAGE_ARM64) | tr :/ __)
BASE_IMAGE := local/base:1.1.3-$(shell echo $(ROOT_IMAGE) | tr :/ __)-$(shell echo $(CERTS_IMAGE) | tr :/ __)

package-base:
	(docker image ls --format '{{.Repository}}:{{.Tag}}' | grep -q '$(BASE_IMAGE)$$') \
		|| docker build \
			--build-arg root_image=$(ROOT_IMAGE) \
			--build-arg certs_image=$(CERTS_IMAGE) \
			--tag $(BASE_IMAGE) \
			deployment/docker/base

package-base-arm64:
	(docker image ls --format '{{.Repository}}:{{.Tag}}' | grep -q '$(BASE_IMAGE)-arm64$$') \
		|| docker build \
		    --platform=linux/arm64 \
			--build-arg root_image=$(ROOT_IMAGE) \
			--build-arg certs_image=$(CERTS_IMAGE) \
			--tag $(BASE_IMAGE)-arm64 \
			deployment/docker/base

package-builder-arm64:
	(docker image ls --format '{{.Repository}}:{{.Tag}}' | grep -q '$(BUILDER_IMAGE_ARM64)$$') \
		|| docker build \
		    --platform=linux/arm64 \
			--build-arg go_builder_image=$(GO_BUILDER_IMAGE_ARM64) \
			--tag $(BUILDER_IMAGE_ARM64) \
			-f deployment/docker/builder/arm64.Dockerfile \
			deployment/docker/builder


package-builder:
	(docker image ls --format '{{.Repository}}:{{.Tag}}' | grep -q '$(BUILDER_IMAGE)$$') \
		|| docker build \
			--build-arg go_builder_image=$(GO_BUILDER_IMAGE) \
			--tag $(BUILDER_IMAGE) \
			deployment/docker/builder

app-via-docker: package-builder
	mkdir -p gocache-for-docker
	docker run --rm \
		--user $(shell id -u):$(shell id -g) \
		--mount type=bind,src="$(shell pwd)",dst=/VictoriaMetrics \
		-w /VictoriaMetrics \
		--mount type=bind,src="$(shell pwd)/gocache-for-docker",dst=/gocache \
		--env GOCACHE=/gocache \
		--env GO111MODULE=on \
		$(DOCKER_OPTS) \
		$(BUILDER_IMAGE) \
		go build $(RACE) -mod=vendor -trimpath \
			-ldflags "-extldflags '-static' $(GO_BUILDINFO)" \
			-tags 'netgo osusergo nethttpomithttp2 musl' \
			-o bin/$(APP_NAME)$(APP_SUFFIX)-prod $(PKG_PREFIX)/app/$(APP_NAME)

app-via-docker-build-arm64: package-builder-arm64
	mkdir -p gocache-for-docker
	docker run --rm \
		--user $(shell id -u):$(shell id -g) \
		--mount type=bind,src="$(shell pwd)",dst=/VictoriaMetrics \
		-w /VictoriaMetrics \
		--mount type=bind,src="$(shell pwd)/gocache-for-docker",dst=/gocache \
		--env GOCACHE=/gocache \
		--env GO111MODULE=on \
		--platform=linux/arm64 \
		--env CGO_ENABLED=1 \
		--env GOOS=linux \
		--env GOARCH=arm64 \
		$(BUILDER_IMAGE_ARM64) \
		go build $(RACE) -mod=vendor \
			-ldflags "-extldflags '-static' $(GO_BUILDINFO)" \
			-tags 'netgo osusergo nethttpomithttp2' \
			-o bin/$(APP_NAME)$(APP_SUFFIX)-prod $(PKG_PREFIX)/app/$(APP_NAME)


app-via-docker-windows: package-builder
	mkdir -p gocache-for-docker
	docker run --rm \
		--user $(shell id -u):$(shell id -g) \
		--mount type=bind,src="$(shell pwd)",dst=/VictoriaMetrics \
		-w /VictoriaMetrics \
		--mount type=bind,src="$(shell pwd)/gocache-for-docker",dst=/gocache \
		--env GOCACHE=/gocache \
		--env GO111MODULE=on \
		$(DOCKER_OPTS) \
		$(BUILDER_IMAGE) \
		go build $(RACE) -mod=vendor -trimpath \
			-ldflags "-s -w -extldflags '-static' $(GO_BUILDINFO)" \
			-tags 'netgo osusergo nethttpomithttp2' \
			-o bin/$(APP_NAME)-windows$(APP_SUFFIX)-prod.exe $(PKG_PREFIX)/app/$(APP_NAME)

package-via-docker: package-base
	(docker image ls --format '{{.Repository}}:{{.Tag}}' | grep -q '$(DOCKER_NAMESPACE)/$(APP_NAME):$(PKG_TAG)$(APP_SUFFIX)$(RACE)$$') || (\
		$(MAKE) app-via-docker && \
		docker build \
			--build-arg src_binary=$(APP_NAME)$(APP_SUFFIX)-prod \
			--build-arg base_image=$(BASE_IMAGE) \
			--tag $(DOCKER_NAMESPACE)/$(APP_NAME):$(PKG_TAG)$(APP_SUFFIX)$(RACE) \
			-f app/$(APP_NAME)/deployment/Dockerfile bin)

publish-via-docker: \
	app-via-docker-amd64 \
	app-via-docker-arm \
	app-via-docker-arm64 \
	app-via-docker-ppc64le \
	app-via-docker-386
	docker buildx build \
		--platform=linux/amd64,linux/arm,linux/arm64,linux/ppc64le,linux/386 \
		--build-arg certs_image=$(CERTS_IMAGE) \
		--build-arg root_image=$(ROOT_IMAGE) \
		--build-arg APP_NAME=$(APP_NAME) \
		--tag $(DOCKER_NAMESPACE)/$(APP_NAME):$(PKG_TAG)$(RACE) \
		--tag $(DOCKER_NAMESPACE)/$(APP_NAME):latest$(RACE) \
		-o type=image \
		-f app/$(APP_NAME)/multiarch/Dockerfile \
		--push \
		bin

run-via-docker: package-via-docker
	docker run -it --rm \
		--user $(shell id -u):$(shell id -g) \
		--net host \
		$(DOCKER_OPTS) \
		$(DOCKER_NAMESPACE)/$(APP_NAME):$(PKG_TAG)$(APP_SUFFIX)$(RACE) $(ARGS)

app-via-docker-goarch:
	APP_SUFFIX='-$(GOARCH)' \
	DOCKER_OPTS='--env CGO_ENABLED=$(CGO_ENABLED) --env GOOS=linux --env GOARCH=$(GOARCH)' \
	$(MAKE) app-via-docker

app-via-docker-windows-goarch:
	APP_SUFFIX='-$(GOARCH)' \
	DOCKER_OPTS='--env CGO_ENABLED=0 --env GOOS=windows --env GOARCH=$(GOARCH)' \
	$(MAKE) app-via-docker-windows

app-via-docker-goarch-cgo:
	CGO_ENABLED=1 $(MAKE) app-via-docker-goarch

app-via-docker-goarch-nocgo:
	CGO_ENABLED=0 $(MAKE) app-via-docker-goarch

app-via-docker-pure:
	APP_SUFFIX='-pure' DOCKER_OPTS='--env CGO_ENABLED=0' $(MAKE) app-via-docker

app-via-docker-amd64:
	GOARCH=amd64 $(MAKE) app-via-docker-goarch-cgo

app-via-docker-arm:
	GOARCH=arm $(MAKE) app-via-docker-goarch-nocgo

app-via-docker-arm64:
	GOARCH=arm64 APP_SUFFIX='-arm64' \
	$(MAKE) app-via-docker-build-arm64

app-via-docker-ppc64le:
	GOARCH=ppc64le $(MAKE) app-via-docker-goarch-nocgo

app-via-docker-386:
	GOARCH=386 $(MAKE) app-via-docker-goarch-nocgo

app-via-docker-windows-amd64:
	GOARCH=amd64 $(MAKE) app-via-docker-windows-goarch

package-via-docker-goarch:
	APP_SUFFIX='-$(GOARCH)' \
	DOCKER_OPTS='--env CGO_ENABLED=$(CGO_ENABLED) --env GOOS=linux --env GOARCH=$(GOARCH)' \
	$(MAKE) package-via-docker

package-via-docker-goarch-cgo:
	CGO_ENABLED=1 $(MAKE) package-via-docker-goarch

package-via-docker-goarch-nocgo:
	CGO_ENABLED=0 $(MAKE) package-via-docker-goarch

package-via-docker-pure:
	APP_SUFFIX='-pure' DOCKER_OPTS='--env CGO_ENABLED=0' $(MAKE) package-via-docker

package-via-docker-amd64:
	GOARCH=amd64 $(MAKE) package-via-docker-goarch-cgo

package-via-docker-arm:
	GOARCH=arm $(MAKE) package-via-docker-goarch-nocgo

package-via-docker-arm64: package-base-arm64
	GOARCH=arm64 APP_SUFFIX='-arm64' \
	DOCKER_OPTS='--platform=linux/arm64 --env CGO_ENABLED=1 --env GOOS=linux --env GOARCH=$(GOARCH)' \
	$(MAKE) app-via-docker-build-arm64
	docker build \
                --platform=linux/arm64 \
    			--build-arg src_binary=$(APP_NAME)-arm64-prod \
    			--build-arg base_image=$(BASE_IMAGE)-arm64 \
    			--tag $(DOCKER_NAMESPACE)/$(APP_NAME):$(PKG_TAG)-arm64$(RACE) \
    			-f app/$(APP_NAME)/deployment/Dockerfile bin

package-via-docker-ppc64le:
	GOARCH=ppc64le $(MAKE) package-via-docker-goarch-nocgo

package-via-docker-386:
	GOARCH=386 $(MAKE) package-via-docker-goarch-nocgo

remove-docker-images:
	docker image ls --format '{{.Repository}}\t{{.ID}}' | grep $(DOCKER_NAMESPACE)/ | grep -v /builder | awk '{print $$2}' | xargs docker image rm -f
