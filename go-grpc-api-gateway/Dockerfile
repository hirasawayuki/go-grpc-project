FROM golang:1.18-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN go mod download
EXPOSE 3000

CMD [ "go", "run", "cmd/main.go" ]
