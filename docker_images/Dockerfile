FROM ubuntu:22.04

RUN apt-get update

COPY greeting.sh /app/

RUN chmod +x /app/greeting.sh

WORKDIR /app

CMD ["/bin/bash", "/app/greeting.sh"]








