package main

import (
)
import (
	"godeploy/deploy"
	"fmt"
	"flag"
	"syscall"
	"io/ioutil"
	"strconv"
	"os"
)

func main()  {

	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("godeploy.pid", []byte(strconv.Itoa(pid)), 0777)
		defer os.Remove("godeploy.pid")
	}


	port := flag.Int("port", 8080, "Listen local port.")
	flag.Parse();

	config := deploy.DeployServer{
		Port	: *port,
	}

	config.Start();


	fmt.Println(config)
}
