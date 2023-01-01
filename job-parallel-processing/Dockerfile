FROM alpine:3.17.0
RUN apk update
RUN apk add --no-cache \
	python3 \
	rabbitmq-c-utils \
	bind-tools \
	curl \
	ca-certificates

RUN ln -sf $(which python3) /usr/bin/python