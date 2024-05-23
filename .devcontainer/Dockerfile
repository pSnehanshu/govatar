FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Install prerequisites
RUN apt-get update

# Install n
ENV N_PREFIX="/home/vscode/n"
ENV PATH "$N_PREFIX/bin:$PATH"
RUN curl -L https://bit.ly/n-install | bash -s -- -y

# Install Node.js
RUN n lts

# Install Go
ENV GOLANG_VERSION 1.22.3
ENV GOPATH "/opt/go"
ENV PATH "$PATH:/usr/local/go/bin:$GOPATH/bin"
RUN curl -sSL https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xzf -

# Install Air, Atlas
RUN go install github.com/cosmtrek/air@latest; \
  curl -sSf https://atlasgo.sh | sh

# Give permission to all users
RUN chmod -R 777 $GOPATH

# Install Redis and PSQL CLI
RUN apt-get install -y redis-tools postgresql-client
