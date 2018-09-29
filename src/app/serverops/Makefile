CONFDIR=/etc/take2-chatops
BINDIR=/usr/local/bin

all: dep build
build:
	go build -v -o serverops main.go
clean:
	go clean
	rm -f serverops
install:
	sudo mkdir -p $(CONFDIR)
	if [ ! -e /etc/take2-chatops/serverops.conf ]; then\
		cp ../../../etc/take2-chatops/take2-chatops.conf serverops.conf;\
		echo "" >> serverops.conf;\
		echo "[serverops]" >> serverops.conf;\
		echo "nodename =" >> serverops.conf;\
		sudo cp serverops.conf /etc/take2-chatops/serverops.conf;\
		sudo cp serverops /usr/local/bin;\
	fi
uninstall:
	sudo rm -f $(BINDIR)/serverops
	sudo rm -rf $(CONFDIR)/serverops.conf
dep:
	go get -v github.com/nlopes/slack
	go get -v github.com/go-ini/ini
