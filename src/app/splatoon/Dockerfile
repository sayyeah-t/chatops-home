FROM resin/rpi-raspbian:buster

MAINTAINER sayyeah-t <@SayYeah1121>

COPY splatoon /usr/bin
COPY requirements.txt /

RUN mkdir /etc/take2-chatops && \
    apt-get update && \
    apt-get install -y git python python-pip && \
    mkdir -p /opt/splatnet2statink && \
    git clone https://github.com/frozenpandaman/splatnet2statink.git \
    /opt/splatnet2statink && \
    pip install -r /requirements.txt


CMD ["/usr/bin/splatoon"]
