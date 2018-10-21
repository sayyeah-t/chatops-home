CONFDIR=/etc/take2-chatops
BINDIR=/usr/local/bin
SERVICEFILEDIR=/usr/lib/systemd/system

all: dep build
build:
	go build -v -o splatoon main.go
clean:
	go clean
	rm -f splatoon
	rm -f splatoon.conf
install:
	sudo mkdir -p $(CONFDIR)
	if [ ! -e $(CONFDIR)/splatoon.conf ]; then\
		cp ../../../etc/take2-chatops/take2-chatops.conf splatoon.conf;\
		echo "" >> splatoon.conf;\
		echo "[splatoon]" >> splatoon.conf;\
		echo "iksm_session =" >> splatoon.conf;\
		echo "api_key =" >> splatoon.conf;\
		echo "uploader_path = /opt/splatnet2statink" >> splatoon.conf;\
		sudo cp splatoon.conf $(CONFDIR)/splatoon.conf;\
		cp ../../../etc/systemd/system/take2-chatops.service splatoon.service;\
		sed -i 's/take2-chatops/splatoon/g' splatoon.service;\
		sudo cp splatoon.service $(SERVICEFILEDIR)/splatoon.service;\
		sudo cp splatoon $(BINDIR);\
	fi
uninstall:
	sudo rm -f $(BINDIR)/splatoon
	sudo rm -f $(CONFDIR)/splatoon.conf
dep:
	go get -v github.com/nlopes/slack
	go get -v github.com/go-ini/ini
