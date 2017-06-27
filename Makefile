.SILENT :
.PHONY : config-keycloak deploy undeploy cli

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

## Configure keycloak
config-keycloak:
	docker-compose $(COMPOSE_FILES) run keycloak_config

## Deploy containers to Docker host
deploy:
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

## Start a CLI
cli:
	echo "Starting CLI..."
	docker run --rm -it \
		--add-host="devbox:172.17.0.1" \
		--env-file todomvc-cli/conf.env \
		ncarlier/keycloak-todomvc-cli

