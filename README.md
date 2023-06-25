# ms-curso-cart-api
Cart

# Catalog Mod
go mod init github.com/j-ew-s/ms-curso-cart-api

## Porta
5300

# Comunica com outros serviços:

1 - AUTH GRPC   pela porta :5500 


# DOCKER 

docker build --tag cart-api:v0.0.1 -f Dockerfile   .

docker compose build

docker compose up

