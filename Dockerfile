FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN make build

# Run stage (will remove source code in order to reduce image size)
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/bin/server .

# Use config maps over copying app.env in the image
# COPY app.env .

EXPOSE 8080

CMD [ "" ]
ENTRYPOINT [ "/app/server" ]