# https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Update and install ca-certificates.
RUN apk update && apk --no-cache add ca-certificates

WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

# Fetch dependencies. Using go get.
RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy binary.
COPY --from=builder /go/bin/main /go/bin/main

# Copy trusted certs.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go/bin/main"]