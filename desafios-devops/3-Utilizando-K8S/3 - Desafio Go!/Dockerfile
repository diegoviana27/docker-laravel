FROM golang as builder
WORKDIR /go/src/desafio/
COPY ./src ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/desafio/desafio /app/desafio
EXPOSE 8000
ENTRYPOINT ./desafio