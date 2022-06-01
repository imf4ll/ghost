install:
	go build
	cp goshot /usr/bin/goshot

uninstall:
	rm -rf /usr/bin/goshot
