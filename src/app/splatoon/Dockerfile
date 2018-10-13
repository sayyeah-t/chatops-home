FROM resin/rpi-raspbian

MAINTAINER sayyeah-t <@SayYeah1121>

COPY splatoon /usr/bin

RUN mkdir /etc/take2-chatops && \
    apt-get update && \
    apt-get install -y git python && \
    mkdir -p /opt/splatnet2statink
    git clone https://github.com/frozenpandaman/splatnet2statink.git /opt/splatnet2statink && \
    pip install -r splatnet2statink/requirements.txt && \


CMD ["/usr/local/bin/splatoon"]
