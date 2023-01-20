TARGET_OS   = "windows" "linux" "darwin"
TARGET_ARCH = "amd64" "arm64"

clean: 
	rm -rf .build
	rm -rf test/targets/*

build: clean
	for os in $(TARGET_OS);do                                                          \
		for arch in $(TARGET_ARCH); do                                                 \
			env GOOS=$$os GOARCH=$$arch go build -o .build/$$os/$$arch/go src/main.go; \
		done;                                                                          \
	done;

