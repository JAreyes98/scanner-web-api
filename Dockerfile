FROM golang:1.18
LABEL authors="jdreyes"

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD go run main.go
#CMD ["app"]