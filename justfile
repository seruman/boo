build:
    go build -o boo .

test:
    go test ./...

check: check-fix check-format check-vet check-staticcheck check-deadcode

check-fix:
    go fix ./...
    git diff --exit-code

check-format:
    go tool gofumpt -w .
    go tool goimports -w .
    git diff --exit-code

check-vet:
    go vet ./...

check-staticcheck:
    go tool staticcheck ./...

check-deadcode:
    test -z "$(go tool deadcode -test ./...)"
