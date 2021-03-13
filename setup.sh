#!/usr/bin/env bash

set -e

# the HostIP is specific to minikube for EKS for example you would have to curl the load balancer

CURRENT_PATH=$(pwd)
APP_NAME="template-app"
INGRESS_APP_NAME="ingress-nginx"
IP=$(kubectl get pods -l app.kubernetes.io/name=${INGRESS_APP_NAME} -o jsonpath="{...hostIP}" | awk '{print $1}')

# helper func to hide popd and pushd stack prints
# https://stackoverflow.com/questions/25288194/dont-display-pushd-popd-stack-across-several-bash-scripts-quiet-pushd-popd
pushd() {
	command pushd "$@" >/dev/null
}

popd() {
	command popd "$@" >/dev/null
}

# flow start
echo -e "\n\e[7mRunning go test...\e[0m\n"
pushd ${CURRENT_PATH}/pkg
go test
popd

echo -e "\n\e[7mVerifying if the template is valid...\e[0m\n"
pushd ${CURRENT_PATH}/deploy
helm template ${APP_NAME} . -f values.yaml --debug
popd

echo -e "\n\e[7mRunning helm test...\e[0m\n"
pushd ${CURRENT_PATH}/deploy
helm test ${INGRESS_APP_NAME}
popd

echo -e "\n\e[7mInstalling app...\e[0m\n"
if [[ $(helm list | grep -c ${APP_NAME}) -eq 1 ]]; then
	pushd ${CURRENT_PATH}/deploy
	helm upgrade ${APP_NAME} . -f values.yaml --atomic --wait
	popd
else
	pushd ${CURRENT_PATH}/deploy
	helm install ${APP_NAME} . -f values.yaml --wait
	popd
fi

echo -e "\n\e[7mThe application has been installed and is ready to recieve requests on ${IP}!\e[0m"
