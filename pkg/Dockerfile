FROM debian:10

RUN apt-get update && apt-get install -y gcc 

COPY script.sh /usr/bin/

WORKDIR /volume

ENTRYPOINT [ "bash" ]