# Golang api

Simple api in golang using gin and docker.

## initializing

### Requirements:

- Golang
- Gin

### Install the dependencies:

    go mod tidy

### Run the project:

    go mod run main.go

###

    curl http://localhost:8000/products

output: 10px is equal to 0.625rem

### Build the project:

    go build main.go

###

    ./main

###

### Using docker:

#### Build api image:

    docker build -t go-basic-api .

#### Running container:

    docker-compose up -d go-basic-db

#### Listing containers:

    docker container ls

#### Connecting to database;

    docker exec -it go_db psql -U postgres -d code043

#### Receiving logs:

    docker logs -f go-app

## Resources:

- **Vscode:**
  - **install:** https://code.visualstudio.com/download
- **Golang:**
  - **install:** https://go.dev/doc/install
- **Gin:**
  - **repo:** https://github.com/gin-gonic/gin
