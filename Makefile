name = "snake"
win:
	CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CGO_LDFLAGS="-static-libgcc -static -lpthread -Wl,-subsystem,windows" go build -o "./build/$(name).exe" .
linux:
	GOARCH=amd64 CGO_ENABLED=1 go build -o "./build/$(name)-linux" .
