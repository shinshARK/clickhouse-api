FROM debian:stable-slim

# Update packages, install ca-certificates and wait-for-it
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates wait-for-it && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy your pre-built binary into the container
COPY app .

# Expose the port your application listens on
EXPOSE 3000

# Use wait-for-it to delay starting your app until ClickHouse is ready
CMD ["wait-for-it", "clickhouse:9000", "--", "./app"]
