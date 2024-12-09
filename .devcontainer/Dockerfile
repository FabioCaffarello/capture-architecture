ARG VARIANT="jammy"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

# Install necessary packages for SSH agent forwarding
RUN apt-get update && \
    apt-get install -y openssh-client
