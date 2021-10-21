build:
	go build -o hello_linux hello.go

build_osx:
	GOOS=darwin go build -o hello_mac hello.go

build_win:
	GOOS=windows go build hello.go

all: build build_osx build_win

clean:
	@rm -f hello.exe hello_mac hello_linux
