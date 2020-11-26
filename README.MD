# 开放式的IM服务


openim是分布式推送(IM)服务，完全基于websocket封装，围绕topic进行sub/pub
自身实现订阅管理服务，无需依赖redis


##部署架构：

![openim](document/openim.png)



web示例调用

```
http://127.0.0.1:8199/index?communicateId=1122&userName=聊天人的名字

communicateId ，会话的唯一ID，也可以房间名字

userName，聊天人的名字


检查是否有新聊天记录未读

GET方法

http://127.0.0.1:8199/newcount?topic=1122&from=张家大公司  

```




## 关于build.sh编译脚本
支持将git的tag编译到程序中。需要创建git的tag。

```
git tag v0.1.9

git push origin v0.1.9
```

## ##感谢 firetower 项目
本项目是从[firetower](https://github.com/holdno/firetower) 项目移植过来的。
