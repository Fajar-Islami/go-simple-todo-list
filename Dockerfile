# Start by building the application.
FROM golang:alpine3.17 as build
LABEL stage=dockerbuilder
WORKDIR /app
COPY . .

# Build the binary
RUN go build -o apps cmd/main.go

# Now copy it into our base image.
FROM alpine:3.9

# Copy bin file
WORKDIR /app
COPY --from=build /app/apps /app/apps
COPY --from=build /app/migrations /app/migrations
RUN mkdir /app/logs

EXPOSE 8090
ENTRYPOINT ["/app/apps"]