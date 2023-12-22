# Kubernetes kompose
### Prerequisites
- [Kompose](https://kompose.io/installation/)

### Create Kubernetes manifests files (Kompose)
From the root directory with explicit variables in the `docker-compose.yml`, run the following command:
```shell
# Create Kubernetes manifests files
kompose convert -f docker-compose.yml
```
Move the files into the appropriate directory:

<p>
    <img alt="kubernetes" src="k8s_structure.png" width="450"/>
</p>

Don't forget to add the docker image path in the `app-deployment.yaml` file:
```yaml
...
spec:
  containers:
  - image: maximehrt/devops-project-app:latest
...
```
To synchronize the app service with **Azure** add in the `app-service.yaml` file:
```yaml
...
spec:
  type: LoadBalancer
...
```

### Persistent Volume Claim
Create a persistent volume claim for Redis database to store data in a persistent volume which means that the data will be stored even if the pod is deleted.

Add a `redis-pvc.yaml` file in the `k8s/redis` directory:
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: redis-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```