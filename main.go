package main

import (
)
import (
	"godeploy/deploy"
	"fmt"
	"flag"
)

func main()  {

	port := flag.Int("port", 8080, "Listen local port.")
	flag.Parse();
	fmt.Println(*port);

	config := deploy.DeployServer{
		Port	: *port,
	}

	config.Start();


	fmt.Println(config)
	//
	//err := deploy.ScirptCmannd("E:\\go\\backshop","E:\\Go\\src\\godeploy\\scripts\\backshop.bat");
	//
	//if err != nil{
	//	fmt.Println(err)
	//}
}
