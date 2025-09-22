FROM golang:alpine

ARG API_PORT

WORKDIR /api

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o godfs ./main.go

EXPOSE ${API_PORT}

CMD ["./godfs"]
