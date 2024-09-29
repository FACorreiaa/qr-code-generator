package main

import (
	"github.com/skip2/go-qrcode"
)

func main() {
	//assetsDir := "assets"
	//content := filepath.Join(assetsDir, "contact.png")
	//
	//err := os.MkdirAll(assetsDir, os.ModePerm)
	//if err != nil {
	//	panic("Failed to create assets directory")
	//}

	err := qrcode.WriteFile("https://pamozi.de/contact/", qrcode.Medium, 125, "qr.png")
	if err != nil {
		panic(err)
	}
}
