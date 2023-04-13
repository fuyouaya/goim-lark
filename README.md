```

|——apps(微服务)
|——————interfaces (api)
|——————————cmd
|——————————dig
|——————————internal
|————————————————config 
|————————————————ctrl (控制器)
|————————————————dto (网络请求传参的模型)
|————————————————ctrl (控制器)      
|————————————————router (路由分组)  
|————————————————server (服务启动)
|————————————————service ()         
|——————auth
|——————————client (gprc 对外接口)
|——————————cmd (程序入口)
|——————————dig (代码注入)
|——————————internal
|——————————————————config
|——————————————————server
|——————————————————service

|——assets (静态资源)    

|——business (公共业务代码)

|——config ()

|——docker (集群)

|——docs (文档)

|——domain (数据库交互)
|——————po (表模型)

|——examples (测试)

|——pkg (公共代码)
|——————common
|——————middleware
|——————proto
|——————utils (公共代码)

|——scripts (脚本文件)

```