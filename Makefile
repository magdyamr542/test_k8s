SHELL=/bin/bash

all: 
	kubectl apply -f lb/caller
	kubectl apply -f lb/app

caller: buildCaller
	kubectl apply -f lb/caller

app: buildApp
	kubectl apply -f lb/app/app-backend.k8s.yaml

buildApp: 
	docker build ./lb/app -t goapp:latest; kind load docker-image goapp:latest

buildCaller: 
	docker build ./lb/caller -t gocaller:latest; kind load docker-image gocaller:latest

kill: 
	kubectl delete -f lb/app/
	kubectl delete -f lb/caller/
