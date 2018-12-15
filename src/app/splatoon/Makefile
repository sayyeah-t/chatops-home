APPNAME=splatoon
CONFDIR=/etc/take2-chatops
BINDIR=/usr/local/bin
SERVICEFILEDIR=/usr/lib/systemd/system
BASECONFDIR=../../../etc/take2-chatops
BASESERVICEFILEDIR=../../../etc/systemd/system
DEB_WORKDIR=package

all: dep build
build:
	go build -v -o $(APPNAME) main.go
	cat $(BASECONFDIR)/take2-chatops.conf $(BASECONFDIR)/$(APPNAME).conf > $(APPNAME).conf
	sed 's/take2-chatops/$(APPNAME)/g' $(BASESERVICEFILEDIR)/take2-chatops.service > $(APPNAME).service
clean:
	go clean
	rm -f $(APPNAME)
	rm -f $(APPNAME).conf
	rm -f $(APPNAME).service
	rm -rf $(APPNAME)*.deb
	rm -rf $(DEB_WORKDIR)
install:
	sudo mkdir -p $(CONFDIR)
	if [ ! -e $(CONFDIR)/$(APPNAME).conf ]; then\
		sudo cp $(APPNAME).conf $(CONFDIR)/$(APPNAME).conf;\
		sudo cp $(APPNAME).service $(SERVICEFILEDIR)/$(APPNAME).service;\
		sudo cp $(APPNAME) $(BINDIR);\
	fi
uninstall:
	sudo rm -f $(BINDIR)/$(APPNAME)
	sudo rm -f $(CONFDIR)/$(APPNAME).conf
	sudo rm -f $(SERVICEFILEDIR)/$(APPNAME).service
dep:
	go get -v github.com/nlopes/slack
	go get -v github.com/go-ini/ini
package:
	mkdir -p $(DEB_WORKDIR)$(BINDIR)
	cp $(APPNAME) $(DEB_WORKDIR)$(BINDIR)
	mkdir -p $(DEB_WORKDIR)$(CONFDIR)
	cp $(APPNAME).conf $(DEB_WORKDIR)$(CONFDIR)/_$(APPNAME).conf
	mkdir -p $(DEB_WORKDIR)$(SERVICEFILEDIR)
	cp $(APPNAME).service $(DEB_WORKDIR)$(SERVICEFILEDIR)
	cp -R DEBIAN $(DEB_WORKDIR)
	fakeroot dpkg-deb --build $(DEB_WORKDIR) .
