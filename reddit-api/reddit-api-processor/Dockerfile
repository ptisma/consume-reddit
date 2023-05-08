FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o reddit-api-processor ./cmd/api-processor/ 


CMD [ "./reddit-api-processor" ]