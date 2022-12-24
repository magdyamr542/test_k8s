SHELL=/bin/bash

up: 
	kubectl apply -f lb/caller
	kubectl apply -f lb/app

kill: 
	kubectl delete --ignore-not-found -f lb/app/
	kubectl delete --ignore-not-found -f lb/caller/

caller: buildCaller
	kubectl apply -f lb/caller

app: buildApp
	kubectl apply -f lb/app/app-backend.k8s.yaml

build:  buildCaller buildApp

buildApp: 
	docker build ./lb/app -t goapp:latest; kind load docker-image goapp:latest
	$(MAKE) cleanup

buildCaller: 
	docker build ./lb/caller -t gocaller:latest; kind load docker-image gocaller:latest
	$(MAKE) cleanup

cleanup:  
	docker system prune -f