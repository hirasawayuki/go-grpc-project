FROM golang:1.18-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build cmd/main.go
EXPOSE 3000
CMD [ "./main" ]
