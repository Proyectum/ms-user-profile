# k8s deployments

## Deploy secrets
```shell
kubectl create secret generic ms-user-profile --from-env-file=.env
```