# litemall-go
 

使用go微服务重写[litemall](https://github.com/linlinjava/litemall)api接口

## 环境
- Apple M1
- go version go1.19.4 darwin/arm64
- docker version 20.10.16
- mysql  Ver 8.0.32 for Linux on aarch64 (MySQL Community Server - GPL)
- other

## 安装支持m1的nacos

```azure
docker run --name nacos-standalone -e MOOE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d anonymy/nacos-server-m1:2.0.3
```
访问链接http://127.0.0.1:8848/nacos，输入密码，默认nacos:nacos

## consul
```azure
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul consul agent -dev -client=0.0.0.0
```
默认UI界面端口： http://127.0.0.1:8500/
