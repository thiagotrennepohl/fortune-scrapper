.DEFAULT_GOAL := help

.PHONY: ci tools clean unit-test do-cover cover build image help

NAME    = main
VERSION = 1.0.0
GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	golang.org/x/tools/cmd/cover



ci: clean unit-test build docker ## Continous Integration Steps

clean: ## Remove old binary
	-@rm -f $(NAME); \
	find vendor/* -maxdepth 0 -type d -exec rm -rf '{}' \;

unit-test:  ## Execute tests
	go test -cover ./scrapper

# ci-integration-tests:
  

do-cover: ## Test report
	./scripts/cover.sh

cover: env do-cover env-stop ## [env do-cover env-stop]

build: clean test ## [clean test] Build binary file
	docker build -t thiagotr/fortune-scrapper .

docker: ## Build Docker image
	docker login -u ${DOCKER_LOGIN} -p ${DOCKER_PASSWORD}
	docker push thiagotr/fortune-backend

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'