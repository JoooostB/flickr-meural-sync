FROM golang:1.17-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser for a non-root container
ENV USER=appuser
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src
COPY . .

# Fetch dependencies.
# Using go mod with go 1.11
RUN go get -d -v

# Add support for custom args, to allow multi-arch builds
ARG opts="CGO_ENABLED=0 GOOS=linux GOARCH=amd64"
ENV opts="${opts}"
# Build the binary.
RUN go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/flickr-meural-sync .
########################v######
FROM scratch

ENV GIN_MODE=release

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/flickr-meural-sync /go/bin/flickr-meural-sync

# Use an unprivileged user.
USER appuser:appuser

# Run the binary.
ENTRYPOINT ["/go/bin/flickr-meural-sync"]
