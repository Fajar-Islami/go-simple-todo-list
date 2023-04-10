# Start by building the application.
FROM golang:alpine3.17 as build
LABEL stage=dockerbuilder
WORKDIR /app
COPY . .

# Make docs swagger
RUN command -v swag >/dev/null 2>&1 || { go install github.com/swaggo/swag/cmd/swag@v1.8.10; }
RUN swag init -g ./internal/delivery/http/main.go --output ./docs/

# Build the binary
RUN go build -o apps cmd/main.go

# Now copy it into our base image.
FROM alpine:3.9

# Copy bin file
WORKDIR /app
COPY --from=build /app/apps /app/apps
COPY --from=build /app/docs /app/docs
COPY --from=build /app/migrations /app/migrations
RUN mkdir /app/logs

EXPOSE 8090
ENTRYPOINT ["/app/apps"]