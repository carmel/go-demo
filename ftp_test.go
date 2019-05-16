package main

import (
	"gopkg.in/dutchcoders/goftp.v1"
	"log"
	"os"
)

func main() {
	var err error
	var ftp *goftp.FTP

	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")
	if ftp, err = goftp.ConnectDbg("47.96.146.146:21"); err != nil {
		panic(err)
	}
	log.Println("Successfully connected")
	defer ftp.Close()
	log.Println("Successfully connected")

	// Username / password authentication
	if err = ftp.Login("root", "Godman@31"); err != nil {
		panic(err)
	}

	var file *os.File
	if file, err = os.Open("C:\\Users\\Vector\\Desktop\\聖經簡報站：末日徵兆.mp4"); err != nil {
		panic(err)
	}
	if err := ftp.Stor("/home/2.mp4", file); err != nil {
		panic(err)
	}

}
