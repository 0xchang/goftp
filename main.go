package main

import (
	"flag"
	"fmt"

	"gopkg.in/dutchcoders/goftp.v1"
)

func main() {
	var err error
	var ftp *goftp.FTP

	var port int
	var ip string
	var username string
	var password string

	flag.StringVar(&ip, "t", "127.0.0.1", "ip")
	flag.IntVar(&port, "p", 21, "port")
	flag.StringVar(&username, "u", "anonymous", "user")
	flag.StringVar(&password, "pwd", "", "password")
	flag.Parse()

	if ftp, err = goftp.Connect(fmt.Sprintf("%v:%v", ip, port)); err != nil {
		panic(err)
	}

	defer ftp.Close()
	fmt.Println("Successfully connected to", fmt.Sprintf("%v:%v", ip, port))

	// Username / password authentication
	if err = ftp.Login(username, password); err != nil {
		panic(err)
	}

	if err = ftp.Cwd("/"); err != nil {
		panic(err)
	}

	var curpath string
	if curpath, err = ftp.Pwd(); err != nil {
		panic(err)
	}

	fmt.Printf("Current path: %s", curpath)

	// Get directory listing
	var files []string
	if files, err = ftp.List("/"); err != nil {
		panic(err)
	}
	fmt.Println("Directory listing:\n", files)

}
