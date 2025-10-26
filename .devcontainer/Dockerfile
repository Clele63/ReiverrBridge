FROM mcr.microsoft.com/devcontainers/base:trixie

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        npm \
        curl \
        bash-completion \
    && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION=1.25.0
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:/home/vscode/go/bin:${PATH}"

RUN mkdir -p /workspaces/ReiverrBridge
WORKDIR /workspaces/ReiverrBridge

USER vscode

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest