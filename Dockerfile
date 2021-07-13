# CI Build
FROM golang:1.17beta1 as builder

WORKDIR /trillian-agent

ARG GOFLAGS=""
ENV GOFLAGS=$GOFLAGS
ENV GO111MODULE=on

# Download dependencies first - this should be cacheable.
COPY go.mod go.sum ./
RUN go mod download

# Now add the local Trillian repo, which typically isn't cacheable.
COPY . .
# Build the server.
RUN go get ./cmd/trillian-agent-server

# Package only executable
# Make a minimal image.
FROM gcr.io/distroless/base

COPY --from=builder /go/bin/trillian-agent-server /

ENTRYPOINT ["./trillian-agent-server"]
