FROM golang:1.14.3-alpine AS build_base
# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/pallabpain/cloud-native-go
COPY . .
# Build the Go app
RUN go build -o cloud-native-go .
RUN chmod 777 cloud-native-go

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates
# Copy the cloud-native-go executable from the base image to our target image
COPY --from=build_base /go/src/github.com/pallabpain/cloud-native-go/cloud-native-go /app/cloud-native-go
# This container exposes port 8000 to the outside world
EXPOSE 8000
# Run the binary program produced by `go build`
CMD ["/app/cloud-native-go"]
