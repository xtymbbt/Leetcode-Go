package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("https://101.201.121.145/merryChristmas/index.html", qrcode.Medium, 256, "./golang_qrcode.png")
	if err != nil {
		fmt.Printf("err is: %#v\n", err)
	}
}