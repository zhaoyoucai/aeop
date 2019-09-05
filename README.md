# aeop
速卖通平台API

## 安装
go get github.com/zhaoyoucai/aeop

## 使用说明
配置文件在conf/config.conf,将获取的app_key等信息设置好,在top/base.go中修改其路径  

[app]  
app_key = 2374718  
app_secret = 2b36fcbc572sa764c93c6386347bbf0  
redirect_uri = https://www.shangshudb.com/  
server_url = https://gw.api.taobao.com/router/rest  


## 依赖
go get github.com/Unknwon/goconfig
读取配置文件的库
