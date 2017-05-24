.SILENT :
.PHONY :

APPNAME:=keycloak-todomvc-slides

# Define base image
BASEIMAGE=node:6-onbuild

# Define port
PORT?=3000
PORTS_FLAGS=-p $(PORT):1948

# Custom run flags
RUN_CUSTOM_FLAGS?=$(PORTS_FLAGS)
RUN_FLAGS=--rm -it $(RUN_CUSTOM_FLAGS)

# Include common Make tasks
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
makefiles:=$(root_dir)/../makefiles
include $(makefiles)/help.Makefile
include $(makefiles)/docker.Makefile

# Build static website
build-site:
	npm install
	npm run dist
	mkdir -p _build/assets _build/img
	cp -R css fonts _build/assets/
	cp img/*.* _build/img


