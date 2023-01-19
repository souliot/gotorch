.PHONY: all install build env clean run

PROJ_MAIN=./cmd
OUT_ROOT=_build

GIT_REPO=github.com/souliot/gotorch
APP_NAME=pytorch
APP_INFO=$(GIT_REPO)/version

PROJ_REV=`git rev-parse --short HEAD`
PROJ_BRANCH=`git rev-parse --abbrev-ref HEAD`
PROJ_BUILD_USER=`git config user.name`
PROJ_BUILD_TIME=`date +%FT%T%z`

all: install

install: build

build: 
	go build -v -x -ldflags " \
		-X $(APP_INFO).Revision=$(PROJ_REV) \
		-X $(APP_INFO).Branch=$(PROJ_BRANCH) \
		-X $(APP_INFO).BuildUser=$(PROJ_BUILD_USER) \
		-X $(APP_INFO).BuildTiime=$(PROJ_BUILD_TIME)" \
		-o $(OUT_ROOT)/bin/$(APP_NAME) \
		$(PROJ_MAIN)

env:
	mkdir -p $(OUT_ROOT)/{bin,conf}

clean:
	rm -rf ./$(OUT_ROOT) 

run:
	$(OUT_ROOT)/bin/$(APP_NAME)
