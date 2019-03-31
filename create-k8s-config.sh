echo "apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: $(echo ${K8S_CERT})
    server: ${K8S_CLUSTER_ADDR}
  name: ${K8S_CLUSTER_NAME}
contexts:
- context:
    cluster: ${K8S_CLUSTER_NAME}
    namespace: $NAMESPACE
    user: $K8S_USERNAME
  name: $K8S_USERNAME-$K8S_CLUSTER_NAME
current-context: $K8S_USERNAME-$K8S_CLUSTER_NAME
kind: Config
preferences: {}
users:
- name: $K8S_USERNAME
  user:
    client-certificate-data: \"$(echo $K8S_CLIENT_CERT)\"
    client-key-data: \"$(echo $K8S_CLIENT_KEY )\"" > ~/.kube/config