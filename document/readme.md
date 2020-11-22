# 使用说明

应用说明：

openim 聊天websocket及http服务

TopicServer 主题服务（是支持聊天服务的基础服务，必须运行它）

curl.sh openim服务的后台启动脚本
curltopic.sh TopicServer服务的后台启动脚本。

脚本使用方式：
```
./curl.sh pid|start|stop|restart|status|tail

```
分别对应 运行ID、启动、停止、重启、状态、显示动态运行信息


通过web访问服务：

http://127.0.0.1:8199/index?communicateId=1122&userName=聊天人的名字

communicateId ，会话的唯一ID，也可以房间名字

userName，聊天人的名字

跟据需要，可以参照web的调用方式，进行使用。支持多人聊天。


## 接口说明

/newcount GET类型，检查该用户是否有新聊天记录未读

http://127.0.0.1:8199/newcount?topic=1122&from=张家大公司

/history GET类型，列出历史聊天记录

http://127.0.0.1:8199/history?topic=1122&from=张家大公司




## 注意事项

请求Header参数

secret-key为api访问密钥，这个取决于本服务程序是否启用了密钥安全控制。默认不启用。

可进行安全控制的接口有：

history 沟通信息历史记录接口
