FROM golang:1.17-alpine

# RUN mkdir app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY src ./src/
COPY config.yml ./

RUN go build ./src/*.go

CMD [ "./server" ]

EXPOSE 3000