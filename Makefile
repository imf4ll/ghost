install-linux:
	mkdir /opt/ghost/
	cp ./sounds/* /opt/ghost/
	go install

uninstall-linux:
	rm -rf ~/go/bin/ghost
	rm -rf /opt/ghost/

# I can't test this on Windows, if you want to contribute,
# do that and make a pull request :)
install-win:
	go build

uninstall-win:
	rm ghost
