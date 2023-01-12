FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy
RUN ls -ltr
RUN go mod download

ENV MICROSERVICECDEMONEWAPI "local"

RUN go build -o imdbapi-service

CMD ["./imdbapi-service"]