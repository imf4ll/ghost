install:
	mkdir /opt/goshot/
	cp ./sounds/* /opt/goshot/
	go build
	cp goshot /usr/bin/goshot

uninstall:
	rm -rf /usr/bin/goshot
	rm -rf /opt/goshot/
