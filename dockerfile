FROM golang:1.22.3-alpine

WORKDIR /Forum

RUN apk update && apk add --no-cache \
    bash \
    build-base \
    sqlite-dev \
    && rm -rf /var/cache/apk/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o forum ./cmd/main.go

EXPOSE 4444

LABEL name="yhrouk,zabdelal,aboutamgh,adraoui,selasly"

CMD ["./forum"]