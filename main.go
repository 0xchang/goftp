package main

import (
	"fmt"
	"os"

	"gopkg.in/dutchcoders/goftp.v1"
)

func main() {
	var err error
	var ftp *goftp.FTP

	var port int
	var ip string
	var username string
	var password string

	fmt.Print("IP: ")
	fmt.Scanln(&ip)

	fmt.Print("Port: ")
	fmt.Scanln(&port)

	if port == 0 {
		port = 21
	}

	fmt.Print("Username: ")
	fmt.Scanln(&username)
	if username == "" {
		username = "anonymous"
	}
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	fmt.Println("username:", username, ", password:", password)

	if ftp, err = goftp.Connect(fmt.Sprintf("%v:%v", ip, port)); err != nil {
		fmt.Println("connect error")
		os.Exit(1)
	}

	defer ftp.Close()
	fmt.Println("Successfully connected to", fmt.Sprintf("%v:%v", ip, port))

	// Username / password authentication
	if err = ftp.Login(username, password); err != nil {
		fmt.Println("error username/password")
		os.Exit(1)
	}

	if err = ftp.Cwd("/"); err != nil {
		fmt.Println("get pwd error")
		os.Exit(1)
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
