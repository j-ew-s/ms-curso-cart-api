FROM golang:1.17.3  AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o  /ms-curso-cart-api



FROM gcr.io/distroless/base-debian10 

WORKDIR /
COPY --from=build /ms-curso-cart-api  /cart_api

EXPOSE 5300

USER nonroot:nonroot

CMD ["/cart_api"]