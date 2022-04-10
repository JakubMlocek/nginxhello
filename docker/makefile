#vars
IMAGENAME=nginxhello
REPO=jakmlo
IMAGEFULLNAME=${REPO}/${IMAGENAME}

.PHONY: help build push all

help:
	    @echo "Makefile commands:"
	    @echo "build"
	    @echo "push"
	    @echo "all"

.DEFAULT_GOAL := all

build:
	    @docker build -t ${IMAGEFULLNAME} .

push:
	    @docker push ${IMAGEFULLNAME}

all: build push
