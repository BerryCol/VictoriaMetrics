ARG go_builder_image
FROM $go_builder_image
STOPSIGNAL SIGINT
RUN apt-get update && apt-get install -y gcc make && rm -rf /var/lib/apt/lists/*
