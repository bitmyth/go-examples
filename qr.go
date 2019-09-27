package main

import (
	"github.com/tuotoo/qrcode"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)

	fi, err := os.Open("qrcode.png")
	if err != nil {
		logger.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	logger.Println(qrmatrix.Content)
}
