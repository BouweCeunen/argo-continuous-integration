# argo-continuous-integration
Continuous integration with Argo which supports bitbucket private webhooks.

[![DockerHub Badge](https://dockeri.co/image/bouwe/argo-continuous-integration)](https://hub.docker.com/r/bouwe/argo-continuous-integration)

## Installation
```
kubectl apply -f kubernetes/
```

## Usage
Once deployed, add a webhook to your bitbucket repositories ```https://<your-argo-domain>/webhook``` and change the url in the ingress before deploying.