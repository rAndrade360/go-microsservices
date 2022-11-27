FROM golang:alpine as builder

RUN apk add --no-cache git tzdata

WORKDIR /go/src

ENV GO111MODULE=on \
  GOPROXY=direct \
  GOSUMDB=off \
  CGO_ENABLED=0 \
  GOOS=linux  \
  GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -o /go/bin/main .


FROM scratch

COPY --from=builder /go/bin/main .

CMD ["./main"]