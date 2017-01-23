#go-deploy

Golang 实现的gitlab/github 自动部署程序

程序执行的git命令依次为：

```bash
git reset --hard 本路仓库名/远程仓库名

git pull

git checkout 远程仓库名称
```

## 配置

conf.ini 文件配置说明

```ini
#仓库名称
[backshop]

#监听的事件类型：push 推送；tag_push 标签创建
EventType=push

#本地仓库物理路径，用于拉去远程代码后存放的路径
LocalDir=E:\go\backshop

#认证码
Token=123456

#拉去之前执行的脚本路径
BeforeScript=

#拉取之后执行的脚本路径
AfterScript=

#本地分支名称
LocalBranch=origin

```

## 使用

```bash
#编译
go build

#启动并设置监听的端口

godeploy -port=8080

```

## 路径

Gitlab路径类似 `http://192.168.4.104:8080/gitlab`

Github 路径：`http://192.168.4.104:8080/github`

