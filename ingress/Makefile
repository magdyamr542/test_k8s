SHELL=/bin/bash

up: 
	kubectl apply -f service1
	kubectl apply -f service2
	kubectl apply -f .

kill: 
	kubectl delete --ignore-not-found -f service1
	kubectl delete --ignore-not-found -f service2

restartDeployment: 
	kubectl delete deployment service1
	kubectl delete deployment service2
	$(MAKE) -s up

build: build1 build2

build2:  
	docker build ./service1 -t service1:latest; kind load docker-image service1:latest
	$(MAKE) -s cleanup

build2:  
	docker build ./service2 -t service2:latest; kind load docker-image service2:latest
	$(MAKE) -s cleanup

cleanup:  
	docker system prune -f

createIngressController:  
	helm upgrade --install ingress-nginx ingress-nginx \
	--repo https://kubernetes.github.io/ingress-nginx \
	--namespace ingress-nginx --create-namespace

deleteIngressController:  
	kubectl delete all --all -n ingress-nginx
	kubectl delete namespaces ingress-nginx

.PHONY: up kill restartDeployment build cleanup