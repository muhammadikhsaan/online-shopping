FROM golang:1.19

ARG PORT 8000

ENV PORT ${PORT}
ENV MODE RELEASE

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./
RUN go work sync

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/service pensiel.com/cmd

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE ${PORT}

# Run
ENTRYPOINT ["/bin/bash", "-c", "go run pensiel.com/migrations && /build/service"]