# 开放式的IM服务


openim是分布式推送(IM)服务，完全基于websocket封装，围绕topic进行sub/pub
自身实现订阅管理服务，无需依赖redis


##部署架构：

![openim](document/openim.png)



web页面中适当的位置示例调用

```
http://127.0.0.1:8199/index?communicateId=3eb5e3d5cd2c71ef6fce3f391c9e9019&userName=聊天用户的ID

```

communicateId ，会话的唯一ID，也可以叫房间号

注：communicateId的生成方式为： 聊天用户A的ID_聊天用户B的ID进行字符串组合，数字小的在前面。然后进行md5

如：用户A的ID为 3322，用户B的ID为2211，组合后的字符串为 2211_3322，然后再进行md5加密生成一个字串。

userName，聊天人的ID,要求为纯数字



检查是否有新聊天记录未读

GET方法
```
http://127.0.0.1:8199/topiclist?from=15820123456

```

[接口文档](https://docs.apipost.cn/view/21b37b6699e575ef)

## 关于build.sh编译脚本
支持将git的tag编译到程序中。需要创建git的tag。

```
git tag v0.1.9

git push origin v0.1.9
```

## ##感谢 firetower 项目
本项目是从[firetower](https://github.com/holdno/firetower) 项目移植过来的。


类似的服务：

https://www.goeasy.io/

