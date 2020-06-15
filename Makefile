all: clean lexdata build test


build:
	go build -v
clean:
	rm -f handler.ReadLexem.go
	rm -f golexidl

lexdata: handler.ReadLexem.l
	golex -o handler.ReadLexem.temp         handler.ReadLexem.l
	gofmt  handler.ReadLexem.temp > handler.ReadLexem.go
	rm -f handler.ReadLexem.temp

test:
	go test -v