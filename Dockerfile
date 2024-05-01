FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM golang:alpine AS final

WORKDIR /app

COPY --from=build /app/main /app/main


EXPOSE 8080
EXPOSE 80

ENTRYPOINT [ "./main" ]