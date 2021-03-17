#
# Makefile
#
install:
	go build -o bin/lr .
	sudo cp bin/lr /usr/local/bin/lr
