FROM golang:1.18-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN go mod download

CMD [ "go", "run", "cmd/main.go" ]
