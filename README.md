# Fortune Scrapper ![alt](https://travis-ci.org/thiagotrennepohl/fortune-scrapper.svg?branch=master)

### Running locally

`docker-compose -f dev-docker-compose.yml`

### Testing

`make unit-test`


### Running on Travis
> This project already has a .trais.yml, take a look

`make ci`

For running on travis the following environments are required

| Env              |               Description                |
| ---------------- | :--------------------------------------: |
| K8S_CERT         | Kubernetes CA certificate base64 encoded |
| K8S_CLUSTER_ADDR |          Kubernetes api address          |
| K8S_CLUSTER_NAME |         Kubernetes cluster name          |
| K8S_USERNAME     | Service Account with enough permissions  |
| K8S_CLIENT_CERT  |     Svc Account base64 encoded Cert      |
| K8S_CLIENT_KEY   |      Svc Account base64 encoded Key      |


## Docker push
> DOCKER_LOGIN, DOCKER_PASSWORD and DOCKER_HUB_NAMESPACE environments are required

`make docker`

### Building docker image

`make build`