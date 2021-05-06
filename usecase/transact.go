package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
)

func create_qr() (string, error) {
	var png []byte
	err := qrcode.WriteFile("User_ID.json", qrcode.Medium, 256, "qr.png")

	return qr.png, err
}
