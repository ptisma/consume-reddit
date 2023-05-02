FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o reddit-api-fetcher ./cmd/api-fetcher/ 


CMD [ "./reddit-api-fetcher" ]