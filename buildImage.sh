!#/bin/bash

docker build ./lb/app -t goapp:latest; kind load docker-image goapp:latest