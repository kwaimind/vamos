build:
	go build && mv vamos ~/go/bin/

run:
	go run main.go

test:
	go test ./...

lint:
	golangci-lint run

# Release builds
release-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o vamos
	tar -czf vamos-$(VERSION)-darwin-amd64.tar.gz vamos
	rm vamos

release-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o vamos
	tar -czf vamos-$(VERSION)-darwin-arm64.tar.gz vamos
	rm vamos

release: release-darwin-amd64 release-darwin-arm64
	@echo "Release tarballs created. Don't forget to:"
	@echo "1. Create GitHub release for v$(VERSION)"
	@echo "2. Upload both .tar.gz files"
	@echo "3. Update Formula/vamos.rb with new SHA256 hashes"
	@echo ""
	@echo "Get SHA256 hashes:"
	@shasum -a 256 vamos-$(VERSION)-darwin-amd64.tar.gz
	@shasum -a 256 vamos-$(VERSION)-darwin-arm64.tar.gz