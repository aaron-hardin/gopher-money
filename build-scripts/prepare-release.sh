gofmt -s -w .
go mod tidy
go test ./...
echo "Don't forget to tag"