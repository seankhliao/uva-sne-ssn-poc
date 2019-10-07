.PHONY: build
build:
	go build -o poc
	sudo setcap 'cap_net_raw,cap_net_admin=eip' ./poc
