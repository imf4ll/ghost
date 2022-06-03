install-linux:
	mkdir /opt/ghost/
	cp ./sounds/* /opt/ghost/
	go build
	cp ghost /usr/bin/ghost

uninstall-linux:
	rm -rf /usr/bin/ghost
	rm -rf /opt/ghost/

# I can't test this on Windows, if you want to contribute,
# do that and make a pull request :)
install-win:
	go build

uninstall-win:
	rm ghost
