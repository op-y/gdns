# gdns

gdns 是一个本地DNS服务的简单实现。

## 项目起源

gdns 的灵感来源于之前工作中一个本地DNS的项目。

当时大家水平有限，做得比较简陋（但是能润！），期间还因为proxy只能指定一个nameserver导致线上故障。

考虑到还是有些场景下，会需要一个方便配置的本地DNS服务，所以参考之前项目，重新实现。

当前进行了初步的结构搭建和DEMO服务实现，还处于玩具状态，希望后续能逐步完善。

## 启动方式

```
# 本地先配置好Golang环境 尽量选择高版本吧

# clone项目到本地
git clone https://github.com/op-y/gdns

# 编译项目
go build -o gdns main.go

# 根据需要调整配置config.toml

# 直接启动即可开发调试
./gdns

# HTTP API 参考docs下接口说明

# DNS 解析可以使用 dig nslookup 等工具调试
dig www.baidu.com @127.0.0.1:30053

```

