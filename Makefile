install-linux:
	mkdir /opt/goshot/
	cp ./sounds/* /opt/goshot/
	go build
	cp goshot /usr/bin/goshot

uninstall-linux:
	rm -rf /usr/bin/goshot
	rm -rf /opt/goshot/

# I can't test this on Windows, if you want to contribute,
# do that and make a pull request :)
install-win:
	go build

uninstall-win:
	rm goshot
