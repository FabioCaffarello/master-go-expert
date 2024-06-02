FROM alpine:latest
WORKDIR /app
COPY dist/server .

# Install CA certificates
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates && curl


CMD ["./server"]
