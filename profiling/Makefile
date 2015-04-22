all: factorial fib

%: %.go
	go build $<
clean:
	rm -f *.prof factorial fib
