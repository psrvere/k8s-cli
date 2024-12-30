This is a cli application to interact with k8s. This implements following functionalities:

Deployment
- Create Deployment: `k8scli create deployment [name] [flags]`
- Delete Deployment: `k8scli delete deployment [name]`

Pod
- Create Pod: `k8scli run pod [name] [flags]`
- Delete Pod: `k8scli delete pod [name]`

CRD
- Create a CustomPod CRD: `k8scli create crd`

CustomResource
- Create a CustomPod Resource: `k8scli create cr [name] [flags]`
