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

### Build the project:

    go build main.go

###

    ./main

###

### Using docker:

#### Build api image:

    docker build -t go-basic-api .

#### Running container database:

    docker-compose up -d go-basic-db

#### Running container api:

    docker-compose up -d go-api-app

#### Listing containers:

    docker container ls

#### Connecting to database;

    docker exec -it go-basic-db psql -U postgres -d code043

###

psql (12.20 (Debian 12.20-1.pgdg120+1)) \
 Type "help" for help.

code043=#

##### Product table:

    create table product (
        id serial primary key,
        product_name varchar(50) not null,
        price numeric(10, 2) not null
    );

#### Criate table:

psql (12.20 (Debian 12.20-1.pgdg120+1)) \
 Type "help" for help.

code043=#create table product (
id serial primary key,
product_name varchar(50) not null,
price numeric(10, 2) not null
);

#### Insert product:

     insert into product (product_name, price) values('Coffee', 20);

###

psql (12.20 (Debian 12.20-1.pgdg120+1)) \
 Type "help" for help.

code043=#insert into product (product_name, price) values('Coffee', 20);

#### Receiving logs:

    docker logs -f go-api-app

### Send request:

    curl http://localhost:8000/products

output: [{"id_product":1,"name":"Coffee","price":20}]

### Using Rest Client:

    get http://localhost:8000/products

#####

```javascript
[
  {
    id_product: 1,
    name: "Coffee",
    price: 20,
  },
];
```

## Resources:

- **Vscode:**
  - **install:** https://code.visualstudio.com/download
- **Golang:**
  - **install:** https://go.dev/doc/install
- **Gin:**
  - **repo:** https://github.com/gin-gonic/gin
- **Docker:**

  - **install:** https://docs.docker.com/get-started/get-docker/

- **Rest Client:**
  - **install:** https://marketplace.visualstudio.com/items?itemName=humao.rest-client
