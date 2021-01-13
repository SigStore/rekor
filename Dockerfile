FROM registry.access.redhat.com/ubi8/go-toolset AS builder
ENV GOPATH=$APP_ROOT

WORKDIR $APP_ROOT/src/
ADD go.mod go.sum $APP_ROOT/src/
RUN go mod download

# Add source code
ADD ./ $APP_ROOT/src/

RUN go build ./cmd/server
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o server_debug ./cmd/server

# Multi-Stage production build
FROM registry.access.redhat.com/ubi8/ubi-minimal as deploy

# Retrieve the binary from the previous stage
COPY --from=builder /opt/app-root/src/server /usr/local/bin/rekor-server

# Set the binary as the entrypoint of the container
CMD ["rekor-server", "serve"]

# debug compile options & debugger
FROM deploy as debug

# overwrite server and include debugger
COPY --from=builder /opt/app-root/src/server_debug /usr/local/bin/rekor-server
COPY --from=builder /usr/bin/dlv /usr/local/bin/dlv
