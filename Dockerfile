# syntax=docker/dockerfile:1

FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /storage ./cmd/storage

FROM scratch

COPY --from=build /storage /storage
ENTRYPOINT ["/storage"]

CMD [ "/storage" ]
