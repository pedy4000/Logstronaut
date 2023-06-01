FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .

RUN apk add -q build-base 
RUN make build

# Run stage (will remove source code in order to reduce image size)
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/bin/server .
COPY app.env .
COPY db/migration ./db/migration

CMD [ "" ]
ENTRYPOINT [ "/app/server" ]