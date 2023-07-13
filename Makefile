update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/talon-one/talon_go/v3

test:
	go test -mod=vendor -run TestGetInventory ./cmd