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
