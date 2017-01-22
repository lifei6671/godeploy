package main

import (
	"fmt"
	"godeploy/deploy"
)

func main()  {

	config := deploy.DeployServer{
		Port:8080,
	}

	config.Start();

	fmt.Println("aaa",config);
}
