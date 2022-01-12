## Firecli

This is a playground repository that explores building go cli commands using cobra.

### Prerequisites

Follow instructions and install Docker Desktop from 
https://docs.docker.com/get-started/#download-and-install-docker-desktop

### Getting started

Build the image with a named tag and run it:
```zsh
docker build -t firecli .
echo "hello" | docker run -i firecli catsay
```

Before running any commands related to prometheus, let's start up two local services:
    - demoapp on port 2112 / A toy Go application that instruments the prometheus /metrics endpoint
    - prometheus on port 9000 - A local prometheus monitoring service 
```zsh
docker-compose -d up #-d for detached mode
```
We should then be able to visit localhost:9000 (prometheus) and localhost:2112/metrics (demoapp)

To stop these services:
```zsh
docker-compose down
```

To delete the local volumes created run:
```zsh
docker-compose down -v
```

### Testing

To run tests locally:
```zsh
go test -v ./... 
```

### Misc Notes

This project was initialized from [cobra](https://github.com/spf13/cobra) project scaffolding at

Init scaffolding usage:
```zsh
go install github.com/spf13/cobra/cobra
cobra init 
go run main.go
```
Or to add a command
```zsh
cobra add commandname
```

Resources for putting together some basic docker/prometheus configs:
- https://dev.to/ablx/minimal-prometheus-setup-with-docker-compose-56mp
- https://prometheus.io/docs/guides/go-application/
- https://docs.docker.com/language/golang/build-images/
