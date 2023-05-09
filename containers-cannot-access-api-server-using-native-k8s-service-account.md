# containers-cannot-access-api-server-using-native-kubernetes-service-account  

## Issue Description

All containers running inside ACI instances with Kubernetes version 1.24+ cannot access the API server using native Kubernetes service account because they cannot find the service account token.

## Cause

The service account token for native Kubernetes service account will not be automatically generated by the API server starting Kubernetes version 1.24 where the feature [Reduction of Secret-based Service Account Tokens](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token) is enabled by default.


## Diagnostics

1. Verify the kubernetes version running inside the cluster is 1.24 or higher using "kubectl version"
2. Create a service account and verify that no secret are automatically created.
```
kubectl create namespace test-namespace
kubectl create sa test-svc-acct -n test-namespace
kubectl get sa -n test-namespace
kubectl get secret -n test-namespace
```

## Mitigation

1. Create a secret manually and assign it to the service account

```
apiVersion: v1
kind: Secret
type: kubernetes.io/service-account-token
metadata:
  name: test-svc-acct-token
  namespace: test-namespace
  annotations:
    kubernetes.io/service-account.name: "test-svc-acct"

```

2. Describe the secret. You will notice a token was created for it. Now the service account uses the secret token we manually created to access the API Server.

```
kubectl describe secret test-svc-acct-token -n test-namespace

```
