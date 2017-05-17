.SILENT :
.PHONY : build-jar build-app config-keycloak build deploy undeploy

# Artifact version
VERSION?=0.0.1-SNAPSHOT

# Execution profile
PROFILE?=local

# Builded artifact
app_dependencies=todomvc-app/node_modules/

# Builded artifact
artifact=todomvc-api/build/libs/todomvc-api-$(VERSION).jar

# Compose files
define COMPOSE_FILES
	-f docker-compose.yml \
	-f docker-compose.config.yml \
	-f docker-compose.app.yml
endef

# Include common Make tasks
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
makefiles:=$(root_dir)/makefiles
include $(makefiles)/help.Makefile
include $(makefiles)/compose.Makefile

## Build gradle artifact
build-jar:
	echo "Building gradle artifact..."
	./todomvc-api/gradlew -p todomvc-api clean build

## Build gradle artifact
build-app:
	echo "Downloading npm dependencies..."
	$(shell cd ./todomvc-app; npm install)

$(artifact):
	echo "$(artifact) artifact not builded. Building..."
	$(MAKE) build-jar

$(app_dependencies):
	echo "$(app_dependencies) does not exists. Downloading..."
	$(MAKE) build-app

## Configure keycloak
config-keycloak:
	docker-compose $(COMPOSE_FILES) run keycloak_config

## Build services
build: $(artifact) $(app_dependencies)
	echo "Building services ..."
	docker-compose $(COMPOSE_FILES) build

## Deploy containers to Docker host
deploy: build
	echo "Deploying infrastructure..."
	-cat .env
	docker-compose up -d
	$(MAKE) config service=keycloak
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.app.yml \
		up -d
	echo "Congrats! Infrastructure deployed."

## Un-deploy API from Docker host
undeploy:
	echo "Un-deploying infrastructure..."
	docker-compose $(COMPOSE_FILES) down
	echo "Infrastructure un-deployed."

