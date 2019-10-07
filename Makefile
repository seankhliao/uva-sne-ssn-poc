.PHONY: build
build:
	go build
	sudo setcap 'cap_net_raw,cap_net_admin=eip' ./uva-sne-ssn-poc
