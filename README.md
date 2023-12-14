## Docker

#### Build

```
docker build -t publisher-image -f Dockerfile .
```

#### Run

```
docker run -d -p 8080:8080 --name publisher-container publisher-image
```

#### Registry

> my-docker-registry is the name of the container registry where the image is hosted. Common container registries include Docker Hub, Google Container Registry (gcr.io), Amazon Elastic Container Registry (ECR)

## RabbitMQ

To start RabbitMQ as a foreground process, use the following command, specifying the path to the RabbitMQ environment configuration file

```
CONF_ENV_FILE="/usr/local/etc/rabbitmq/rabbitmq-env.conf" /usr/local/opt/rabbitmq/sbin/rabbitmq-server
```

To stop RabbitMQ, you can press Ctrl+C in the terminal where RabbitMQ is running. This will gracefully shut down the RabbitMQ server.

#### Amazon MQ

Currently, the production instance of RabbitMQ is hosted on an Amazon MQ instance named `sms-broker`

## Deployment

### Local

Start `minikube`

```
minikube start
```

Direct `minikube` to use the `docker` env. Any `docker build ...` commands after this command is run will build inside the `minikube` registry and will not be visible in Docker Desktop. `minikube` uses its own docker daemon which is separate from the docker daemon on your host machine. Running `docker images` inside the `minikube` vm will show the images accessible to `minikube`

```
eval $(minikube docker-env)
```

```
docker build -t sms-publisher-image:latest .
```

#### Environment Variables (if needed)

```
kubectl create secret generic rabbitmq-secret --from-env-file=./.env
```

```
kubectl apply -f ./k8s/sms-publisher.deployment.yaml
```

```
kubectl apply -f ./k8s/sms-publisher.service.yaml
```

```
kubectl get deployments
```

```
kubectl get pods
```

```
minikube service sms-publisher-service
```

After running the last comment the application will be able to be accessed in the browser at the specified port that `minikube` assigns.

#### Troubleshooting

```
minikube ssh 'docker images'
```

```
kubectl logs <pod-name>
```

```
kubectl logs -f <pod-name>
```
