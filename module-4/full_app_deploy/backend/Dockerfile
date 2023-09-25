FROM golang:latest as base

RUN apt update && \
  apt upgrade -y

WORKDIR /app
COPY . .
RUN go mod tidy && \
  GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o app .

FROM scratch as prd

WORKDIR /app
COPY --from=base /app .

EXPOSE 8888

ENTRYPOINT ["./app"]
