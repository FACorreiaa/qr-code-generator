package main

import (
	"os"
	"path/filepath"

	"github.com/skip2/go-qrcode"
)

func main() {
	assetsDir := "assets"
	content := filepath.Join(assetsDir, "contact.png")

	err := os.MkdirAll(assetsDir, os.ModePerm)
	if err != nil {
		panic("Failed to create assets directory")
	}

	err = qrcode.WriteFile(content, qrcode.Medium, 256, "qr.png")
	if err != nil {
		panic(err)
	}
}
