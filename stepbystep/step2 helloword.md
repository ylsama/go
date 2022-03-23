make dir first hello_word
```bash
mkdir ylsama
cd ylsama
mkdir hello_world
cd hello_word
lvim main.go
```

```go
package main

func main() {
	println("Hello World!")
}
```
```bash
go mod init ylsama/hello_world
go run main.go
go build -o hello_world
./hello_world
```


