# Build image
FROM debian:12 AS tor
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -qq && apt-get install -y apt-transport-https wget gnupg
RUN echo "deb [signed-by=/usr/share/keyrings/tor-archive-keyring.gpg] https://deb.torproject.org/torproject.org bookworm main" >/etc/apt/sources.list.d/tor.list
RUN wget -qO- https://deb.torproject.org/torproject.org/A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89.asc | gpg --dearmor | tee /usr/share/keyrings/tor-archive-keyring.gpg >/dev/null
RUN apt-get update -qq && apt-get install -y tor deb.torproject.org-keyring

# Deploy image
FROM tor
RUN useradd --create-home -d /data -s /bin/bash onionpipe
RUN mkdir -p /app/data-dir && \
    chmod 770 /app/data-dir && \
    chown -R onionpipe:onionpipe /app/data-dir
COPY dist/server .
WORKDIR /app
USER onionpipe
ENTRYPOINT [ "/server" ]