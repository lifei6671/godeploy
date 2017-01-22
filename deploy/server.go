package deploy

import (
	"net/http"
	"fmt"
	"io"
	"strconv"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"bytes"
	"github.com/widuu/gojson"
	"strings"
)

var DefaultConfigs = make([]DeployConfig,0);

type DeployServer struct {
	Port int //监听的端口号
}

type DeployConfig struct{
	RepositoryName string  //git 仓库名称
	LocalDir	string //本地目录地址
	Token 		string //认证Token
}

func (server DeployServer) Start()  {

	fmt.Println("service is starting ...")
	http.HandleFunc("/github",execute);
	http.HandleFunc("/gitlab",execute);

	err := http.ListenAndServe(":" + strconv.Itoa(server.Port),nil);
	if(err != nil){
		fmt.Println("ListenAndServe error: ", err.Error());
	}
	fmt.Println("service is started.");
}

func execute(w http.ResponseWriter,r *http.Request)  {
	path := r.URL.Path
	//request_type := path[strings.LastIndex(path, "."):]
	content,err := ioutil.ReadAll(r.Body);
	if err != nil{
		fmt.Println("post error:", err.Error());
	}
	fmt.Println(path)

	postBody := bytes.NewBuffer(content).String();

	jsonObject := gojson.Json(postBody);

	io.WriteString(w,bytes.NewBuffer(content).String());

	if strings.EqualFold(path,"/gitlab") {

		event := r.Header.Get("X-Gitlab-Event");
		token := r.Header.Get("X-Gitlab-Token");

		fmt.Println(token);

		if strings.EqualFold(event,"Push Hook"){
			fmt.Println(event);
			fmt.Println(jsonObject.Get("project"))
		}
		return ;
	}
	if strings.EqualFold(path,"/github") {

		return ;
	}
	io.WriteString(w,"No supported");
}

func init(){
	configFile,err := goconfig.LoadConfigFile("./conf.ini");
	if err != nil{
		fmt.Println("Config error: ",err.Error())
	}

	configs := make([]DeployConfig,0);

	configSections := configFile.GetSectionList();

	if len(configSections) > 0 {
		for _,section := range configSections {
			config := DeployConfig{
				RepositoryName : configFile.MustValue(section, "RepositoryName"),
			}
			configs = append(configs, config)
		}
	}
	DefaultConfigs = configs;
}