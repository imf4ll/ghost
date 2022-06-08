install-linux:
	mkdir /opt/ghost/
	cp ./sounds/* /opt/ghost/
	go install

uninstall-linux:
	rm -rf ~/go/bin/ghost
	rm -rf /opt/ghost/

install-win:
	mkdir ${HOME}\\AppData\\Roaming\\ghost
	cp ./sounds/* ${HOME}\\AppData\\Roaming\\ghost\\
	go build
	echo "Move the builded binary for any location you want!"

uninstall-win:
	rmdir ${HOME}\\AppData\\Roaming\\ghost
