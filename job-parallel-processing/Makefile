prepare: 
# get used docker images
	docker pull alpine:3.17.0

# build the tester
	docker build -t rabbitmq-tester:latest -f Dockerfile .
	docker build -t queue-worker:latest -f Dockerfile.worker .

	# push to kind
	kind load docker-image rabbitmq:3.10-alpine
	kind load docker-image rabbitmq-tester:latest
	kind load docker-image queue-worker:latest
	$(MAKE) -s cleanup

cleanup:  
	docker system prune -f

broker:  
	kubectl apply -f rabbitmq-replica-set.k8s.yaml
	kubectl apply -f rabbitmq-service.k8s.yaml

connect-tester:  
	kubectl apply -f tester-pod.k8s.yaml
	kubectl wait --for=condition=ready pod  rabbitmq-tester --timeout=120s
	kubectl exec -it pods/rabbitmq-tester -- sh

job:  
	kubectl apply -f worker-job.k8s.yaml

kill:  
	kubectl delete -f .

.PHONY: prepare job connect-tester broker cleanup