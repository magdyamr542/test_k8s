SHELL=/bin/bash

up: 
	kubectl apply -f caller
	kubectl apply -f app

kill: 
	kubectl delete --ignore-not-found -f app
	kubectl delete --ignore-not-found -f caller

caller: buildCaller
	kubectl apply -f caller

app: buildApp
	kubectl apply -f app/app-backend.k8s.yaml

build:  buildCaller buildApp

buildApp: 
	docker build ./app -t goapp:latest; kind load docker-image goapp:latest
	$(MAKE) cleanup

buildCaller: 
	docker build ./caller -t gocaller:latest; kind load docker-image gocaller:latest
	$(MAKE) cleanup

cleanup:  
	docker system prune -f