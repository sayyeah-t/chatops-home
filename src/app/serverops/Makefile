CONFDIR=/etc/take2-chatops
BINDIR=/usr/local/bin
SERVICEFILEDIR=/usr/lib/systemd/system

all: dep build
build:
	go build -v -o serverops main.go
clean:
	go clean
	rm -f serverops
	rm -f serverops.conf
	rm -f serverops.service
install:
	sudo mkdir -p $(CONFDIR)
	if [ ! -e /etc/take2-chatops/serverops.conf ]; then\
		cp ../../../etc/take2-chatops/take2-chatops.conf serverops.conf;\
		echo "" >> serverops.conf;\
		echo "[serverops]" >> serverops.conf;\
		echo "nodename =" >> serverops.conf;\
		sudo cp serverops.conf /etc/take2-chatops/serverops.conf;\
		cp ../../../etc/systemd/system/take2-chatops.service serverops.service;\
		sed -i 's/take2-chatops/serverops/g' serverops.service;\
		sudo cp serverops.service $(SERVICEFILEDIR)/serverops.service;\
		sudo cp serverops $(BINDIR);\
	fi
uninstall:
	sudo rm -f $(BINDIR)/serverops
	sudo rm -f $(CONFDIR)/serverops.conf
	sudo rm -f $(SERVICEFILEDIR)/serverops.service
dep:
	go get -v github.com/nlopes/slack
	go get -v github.com/go-ini/ini
