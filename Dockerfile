FROM golang:1.16-alpine

WORKDIR /app

ENV GIN_MODE=release

COPY . .

RUN go mod download

RUN go build ./

EXPOSE 8080

CMD [ "./tiles-backend-go" ]