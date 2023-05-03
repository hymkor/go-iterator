all:
	go fmt
	go build

test:
	go test -v

demo:
	go run example.go

demo2:
	go run example2.go

docs:
	example-into-readme
