# Building and deploying the Docker Image

## Building Code
```
cd k8ssvcaccttokentest
GOOS=linux go build -o ./app .
```

## Deploy to docker
```
docker build -t <some-registry>/<your-image-name>:tag . 
docker push <some-registry>/<your-image-name>:tag
```
Notice the "." in the end of docker build




