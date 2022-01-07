## Firecli

This is a playground repository that explores building go cli commands using cobra.

### Prerequisites

Follow instructions and install Docker Desktop from 
https://docs.docker.com/get-started/#download-and-install-docker-desktop

### Getting started

Build the image with a named tag and run it:
```zsh
docker build -t firecli .
docker run -p 9090:9090 firecli
```

### Misc Notes

This project was initialized from [cobra](https://github.com/spf13/cobra) project scaffolding at

Example usage:
```zsh
go install github.com/spf13/cobra/cobra
cobra init 
go run main.go
```

Go convention uses URLs for module names. To run local code before committing 
we should add a replace alias in go.mod

