FROM golang:1.18-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go mod download

RUN go build cmd/main.go
CMD [ "./main" ]
