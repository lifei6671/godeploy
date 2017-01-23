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

var DefaultConfigs = make(map[string]DeployConfig,0);

type DeployServer struct {
	Port int //监听的端口号
}

type DeployConfig struct{
	RepositoryName string  //git 仓库名称
	LocalDir	string //本地目录地址
	Token 		string //认证Token
	BeforeScript	string //执行脚本
	AfterScript	string	//拉取之后执行的脚本
	LocalBranch	string //本地分支名称
	EventType	string //监听的事件类型
}

func (server DeployServer) Start()  {

	fmt.Println("service is starting ...")
	http.HandleFunc("/github",githubFunc);
	http.HandleFunc("/gitlab",gitlabFunc);

	err := http.ListenAndServe(":" + strconv.Itoa(server.Port),nil);
	if(err != nil){
		fmt.Println("ListenAndServe error: ", err.Error());
	}
	fmt.Println("service is started.");
}

//处理 gitlab 的请求
func gitlabFunc(w http.ResponseWriter,r *http.Request){
	path := r.URL.Path
	//request_type := path[strings.LastIndex(path, "."):]
	content,err := ioutil.ReadAll(r.Body);
	if err != nil{
		fmt.Println("post error:", err.Error());
	}
	fmt.Println(path)
	//读取请求内容
	postBody := bytes.NewBuffer(content).String();


	//如果是Gitlab
	if strings.EqualFold(path,"/gitlab") {

		event := r.Header.Get("X-Gitlab-Event");
		token := r.Header.Get("X-Gitlab-Token");

		remoteBranch := gojson.Json(postBody).Get("project").Get("default_branch").Tostring();

		projectName := gojson.Json(postBody).Get("project").Get("name").Tostring();

		if section,ok := DefaultConfigs[projectName];ok {

			if !strings.EqualFold(token,section.Token) {
				fmt.Println("Token error");
				return ;
			}

			//如果是Push事件
			if strings.EqualFold(event, "Push Hook") && strings.EqualFold(section.EventType,"push") {
				Command(section,remoteBranch);

			} else if strings.EqualFold(event, "Tag Push Hook") && strings.EqualFold(section.EventType,"tag_push") {
				Command(section,remoteBranch);
			}
		}
		return ;
	}
	fmt.Println("No supported request");

	io.WriteString(w,"No supported request");
}

//处理 github 的请求
func githubFunc(w http.ResponseWriter,r *http.Request)  {
	path := r.URL.Path
	//request_type := path[strings.LastIndex(path, "."):]
	content,err := ioutil.ReadAll(r.Body);
	if err != nil{
		fmt.Println("post error:", err.Error());
	}
	fmt.Println(path)
	//读取请求内容
	postBody := bytes.NewBuffer(content).String();

	//如果是Github
	if strings.EqualFold(path,"/github") {

		event := r.Header.Get("X-GitHub-Event");
		token := r.Header.Get("X-Hub-Signature");

		remoteBranch := gojson.Json(postBody).Get("repository").Get("default_branch").Tostring();

		projectName := gojson.Json(postBody).Get("repository").Get("name").Tostring();

		if section,ok := DefaultConfigs[projectName];ok {

			if !strings.EqualFold(token,section.Token) {
				fmt.Println("Token error");
				return ;
			}

			if err != nil {

				fmt.Println("exec cmmand error:",err.Error())
			}

			return ;
			//如果是Push事件
			if strings.EqualFold(event, "push") && strings.EqualFold(section.EventType,"push") {
				Command(section,remoteBranch);

			} else if strings.EqualFold(event, "create") && strings.EqualFold(section.EventType,"tag_push") {
				Command(section,remoteBranch);
			}
		}

		return ;
	}
	fmt.Println("No supported request");
	io.WriteString(w,"No supported");
}

func Command(config DeployConfig, remoteBranch string)  {
	if config.BeforeScript != "" {
		err := ScirptComannd(config.LocalDir,config.BeforeScript);
		if err != nil {
			fmt.Println("before script error:",err.Error());
			return ;
		}
	}
	err := GitCommand(config.LocalDir ,remoteBranch,config.LocalBranch);
	if err != nil {
		fmt.Println("exec cmmand error:",err.Error());
		return ;
	}
	if config.AfterScript != "" {
		err := ScirptComannd(config.LocalDir,config.AfterScript);
		if err != nil {
			fmt.Println("after script error:",err.Error());
			return ;
		}
	}
}

func init(){
	configFile,err := goconfig.LoadConfigFile("./conf.ini");
	if err != nil{
		fmt.Println("Config error: ",err.Error())
	}

	configs := make(map[string]DeployConfig,0);

	configSections := configFile.GetSectionList();

	if len(configSections) > 0 {
		for _,section := range configSections {
			config := DeployConfig{
				RepositoryName 	: configFile.MustValue(section, "RepositoryName"),
				LocalDir	: configFile.MustValue(section,"LocalDir"),
				Token		: configFile.MustValue(section,"Token"),
				AfterScript	: configFile.MustValue(section,"AfterScript",""),
				BeforeScript	: configFile.MustValue(section,"BeforeScript",""),
				LocalBranch	: configFile.MustValue(section,"LocalBranch","origin"),
				EventType	: configFile.MustValue(section,"EventType","push"),
			}
			configs[section] = config
		}
	}
	DefaultConfigs = configs;
}