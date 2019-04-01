.DEFAULT_GOAL := help

.PHONY: ci tools clean unit-test do-cover cover build image help

NAME    = main
VERSION = 1.0.0
PROJECT_NAME=fortune-scrapper
DOCKER_HUB_NAMESPACE=thiagotr


ci: clean unit-test build docker ## Continous Integration Steps

create-kube-config: ## Create Kubernetes Configs
	mkdir ~/.kube || true && ./create-k8s-config.sh

install-kubectl: ## Install Kubectl
	curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.14.0/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl

clean: ## Remove old binary
	-@rm -f $(NAME); \
	find vendor/* -maxdepth 0 -type d -exec rm -rf '{}' \;

unit-test:  ## Execute unit tests
	go test -cover ./scrapper

build: clean  ## [clean test] Build binary file
	docker build -t thiagotr/fortune-scrapper .

docker-login: ## Logins into docker hub
	docker login -u ${DOCKER_LOGIN} -p ${DOCKER_PASSWORD}

docker: docker-login ##  Push docker image to docker hub
	docker push ${DOCKER_HUB_NAMESPACE}/${PROJECT_NAME}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'