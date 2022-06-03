FROM golang:1.17 AS builder

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

COPY cmd/configs/app.yaml configs/app.yaml
RUN go build -o /bin/url_shortener_gb ./cmd/

CMD ["/bin/url_shortener_gb"]


