GOARCH = amd64
GOOS = windows

all: main.exe

clean:
	rm -f ./main.exe
	rm -f ./test-command.exe

main.exe: test-command.exe
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o ./main.exe .

test-command.exe:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o ./test-command.exe ./test-command
