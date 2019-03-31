#!/bin/bahs

set -e

ENVIRONMENT=$2
NAMESPACE="staging"
if [[ "$ENVIRONMENT" == "production" ]]; then
  ENVIRONMENT="production"
  NAMESPACE="default"
fi


changeDeploymentEntryPoint(){
  COLOR="blue"
  IS_BLUE_UP=$(kubectl --namespace=$NAMESPACE get deployments | grep "$ENVIRONMENT-fortune-app-blue" | wc -l)
  if [[ $IS_BLUE_UP -eq 1 ]]; then
    COLOR="green"
  fi
  changeDeployment $ENVIRONMENT $COLOR
}

changeDeployment(){
  ENVIRONMENT=$1
  COLOR=$2
  echo $MONGO_ADDR
  echo "Updating $COLOR , $ENVIRONMENT $NAMESPACE"
  DEPLOYMENT=$COLOR ENVIRONMENT=${ENVIRONMENT} NAMESPACE=$NAMESPACE MONGO_ADDR=${MONGO_ADDR} envsubst < k8s_fortune_app.yml | kubectl apply -f -
  echo "Waiting 30 secs for image pull and healthcheck"
  sleep 30
  checkHealth $ENVIRONMENT $COLOR
}

checkHealth(){
  ENVIRONMENT=$1
  COLOR=$2
  replicas=$(kubectl --namespace=$NAMESPACE get deployment $ENVIRONMENT-fortune-app-$COLOR -o jsonpath={.spec.replicas})
  echo $replicas
  avaliablePods=$(kubectl --namespace=$NAMESPACE get deployment $ENVIRONMENT-fortune-app-$COLOR -o jsonpath={.status.availableReplicas})
  echo $avaliablePods
  if [[ "$replicas" == "$avaliablePods" ]]; then
    changeService $COLOR
  else
   echo "Deployment failed"
   exit 1
  fi
}

changeService(){
  COLOR=$1
 echo "Updating service $COLOR , $ENVIRONMENT $NAMESPACE"
  DEPLOYMENT=$COLOR ENVIRONMENT=${ENVIRONMENT} NAMESPACE=${NAMESPACE}  envsubst < k8s_fortune_app_service.yml | kubectl apply -f -
  if [[ "$COLOR" == "blue" ]]; then
    kubectl --namespace=$NAMESPACE delete deployment $ENVIRONMENT-fortune-app-green
  else
    kubectl --namespace=$NAMESPACE delete deployment $ENVIRONMENT-fortune-app-blue
  fi
}

while [[ $# -gt 0 ]]
do
key=$1

case $key in
    --set-deployment)
    changeDeploymentEntryPoint $@
    ;;
    *)
    ;;
esac
shift
done

