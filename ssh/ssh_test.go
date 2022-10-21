package ssh

import (
	"log"
	"testing"
)

func TestSsh(t *testing.T) {
	ssh := &MakeConfig{
		User:     "root",
		Server:   "47.96.146.146",
		Port:     "22",
		Password: "Godman@31",
	}

	//	stdout, stderr, done, err := ssh.Run("ls", 60)
	//	if err != nil {
	//		panic("Can't run remote command: " + err.Error())
	//	} else {
	//		log.Println(done, stdout, stderr)
	//	}
	//
	err := ssh.Scp("C:\\Users\\Vector\\Desktop\\聖經簡報站：末日徵兆.mp4", "/home/3.mp4")

	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		log.Println("success")
	}
}
