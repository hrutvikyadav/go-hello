# First Go Executable Package

## From Go site tutorial

Uses function from an imported package to greet user

## Up to this point

```bash
mkdir go-hello
cd go-hello
go mod init
nvim hello.go

go mod edit -replace hrutvik.gg/go_greet=../go_greet
go mod tidy
# go: found hrutvik.gg/go_greet in hrutvik.gg/go_greet v0.0.0-00010101000000-000000000000

go run .
```
