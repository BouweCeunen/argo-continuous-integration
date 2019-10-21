# argo-continuous-integration
Continuous integration with Argo which supports bitbucket private webhooks.

[![DockerHub Badge](https://dockeri.co/image/bouwe/argo-continuous-integration)](https://hub.docker.com/r/bouwe/argo-continuous-integration)

[![jetson-nano](https://bouweceunen.com/vision/jetson-nano.gif)](https://bouweceunen.com/vision/jetson-nano.gif)

## Installation
```
kubectl apply -f kubernetes/
```

## Usage
Once deployed, add a webhook to your bitbucket repositories ```https://<your-argo-domain>/webhook```.

### ingress
Change the url in the ingress before deploying. This Ingress will route traffic from your Argo domain to the Argo CI implementation, Bitbucket will use this domain as a webhook.

### starter workflow 
The webhook is going to be called and it will start up a 'starter' Workflow, which in turn will pull the repository and start the ```argo.yml``` Workflow which has to reside in the root of the repository. The argo.yml file in this repository consists of a ```sshPrivateKeySecret``` named 'bitbucket-creds', this has to be a secret in the namespace you will deploy this in. Feel free to rename this and adjust this in the argo.yml file. This Secret is needed to be able to pull your repository from Bitbucket and start the actual Workflow.
