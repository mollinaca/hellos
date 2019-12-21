package main

import "fmt"

var version string

/*
 golang 自体のバージョンはハードコーディング
 app version は build optionにて指定する
 # go build -ldflags "-X main.version=1.0.0" hello.go
*/

func main() {
  fmt.Printf("Hello World\n")
  println("golang version: go version go1.10.4 linux/amd64")
  println("app version: " + version)
}
