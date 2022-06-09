install-linux:
	go install

uninstall-linux:
	rm -rf ~/go/bin/ghost

install-win:
	go build
	echo "Move the builded binary for any location you want!"
