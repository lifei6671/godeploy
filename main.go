package main

import (
)
import (
	"godeploy/deploy"
	"fmt"
)

func main()  {

	//config := deploy.DeployServer{
	//	Port:8080,
	//}
	//
	//config.Start();
	//
	//
	//fmt.Println(config)

	err := deploy.ScirptCmannd("E:\\go\\backshop","E:\\Go\\src\\godeploy\\scripts\\backshop.bat");

	if err != nil{
		fmt.Println(err)
	}
}
