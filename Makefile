all:
	go build -o strand
	@mv strand /usr/local/bin/

clean:
	@rm -f /usr/local/bin/strand
	go clean
