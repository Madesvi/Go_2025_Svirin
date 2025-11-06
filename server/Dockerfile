FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    netcat-openbsd 

WORKDIR /app

COPY server_v2.sh .
COPY *.html ./
COPY *.http ./

RUN chmod +x server_v2.sh

EXPOSE 3000

CMD ["./server_v2.sh"]
