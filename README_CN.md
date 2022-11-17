# APT Proxy / 轻量 APT 加速工具

[![Security Scan](https://github.com/soulteary/apt-proxy/actions/workflows/scan.yml/badge.svg)](https://github.com/soulteary/apt-proxy/actions/workflows/scan.yml) [![Release](https://github.com/soulteary/apt-proxy/actions/workflows/release.yaml/badge.svg)](https://github.com/soulteary/apt-proxy/actions/workflows/release.yaml) [![Docker Image](https://img.shields.io/docker/pulls/soulteary/apt-proxy.svg)](https://hub.docker.com/r/soulteary/apt-proxy)

<p style="text-align: center;">
  <a href="README.md" target="_blank">ENGLISH</a> | <a href="README_CN.md">中文文档</a>
</p>

<img src="screenshot/logo.png" width="64"/>

> 仅 2MB 大小的、轻量的 **APT 软件包加速工具！**

<img src="screenshot/dockerhub.png" width="600"/>

<img src="screenshot/ui.png" width="600"/>

APT Proxy 是一款轻量的、可靠的 APT 包（Ubuntu / Debian）缓存工具，能够在各种不同的操作系统环境中运行。

你可以将它作为古老的 [apt-cacher-ng](https://www.unix-ag.uni-kl.de/~bloch/acng/) 安全可靠的替代品。

## 支持运行的软硬件环境

- Linux: x86_64 / x86_32
- ARM: ARM64v8 / ARM32v6 / ARM32v7
- macOS: x86_64 / M1 ARM64v8

## 使用方法

很简单，直接运行就行了：

```bash
./apt-proxy

2022/06/12 16:15:40 running apt-proxy
2022/06/12 16:15:41 Start benchmarking mirrors
2022/06/12 16:15:41 Finished benchmarking mirrors
2022/06/12 16:15:41 using fastest mirror https://mirrors.company.ltd/ubuntu/
2022/06/12 16:15:41 proxy listening on 0.0.0.0:3142
```

当你看到类似上面的日志时，一个带有缓存功能的 APT 代理服务就启动完毕了。

然后在需要执行 `apt-get` 命令的地方，对原有命令进行简单的改写：

```bash
# `apt-get update` with apt-proxy service
http_proxy=http://your-domain-or-ip-address:3142 apt-get -o pkgProblemResolver=true -o Acquire::http=true update 
# `apt-get install vim -y` with apt-proxy service
http_proxy=http://your-domain-or-ip-address:3142 apt-get -o pkgProblemResolver=true -o Acquire::http=true install vim -y
```

当我们需要批量重复执行上面的命令，安装大量软件、或者大体积软件的时候，更新和安装的执行速度将会有**巨大的提升**。

### 选择软件源

如果你不希望软件自动选择服务器响应最快的源，这里有两种方式来指定要使用的软件源：

**使用完整的软件源链接地址**

```bash
# proxy cache for both `ubuntu` and `debian`
./apt-proxy --ubuntu=https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ --debian=https://mirrors.tuna.tsinghua.edu.cn/debian/
# proxy cache for `ubuntu` only
./apt-proxy --mode=ubuntu --ubuntu=https://mirrors.tuna.tsinghua.edu.cn/ubuntu/
# proxy cache for `debian` only
./apt-proxy --mode=debian --debian=https://mirrors.tuna.tsinghua.edu.cn/debian/
```

**使用软件源的外号**

```bash
go run apt-proxy.go --ubuntu=cn:tsinghua --debian=cn:163
2022/06/15 10:55:26 running apt-proxy
2022/06/15 10:55:26 using specify debian mirror https://mirrors.163.com/debian/
2022/06/15 10:55:26 using specify ubuntu mirror https://mirrors.tuna.tsinghua.edu.cn/ubuntu/
2022/06/15 10:55:26 proxy listening on 0.0.0.0:3142
```

软件源外号:

- cn:tsinghua
- cn:ustc
- cn:163
- cn:aliyun
- cn:huawei
- cn:tencent

### 加速 Docker 容器中的软件下载

假设我们已经通过类似下面的方式启动了一个容器：

```bash
# Ubuntu
docker run --rm -it ubuntu
# or Debian
docker run --rm -it debian
```

以及，我们在主机上已经运行了一个 Apt-Proxy 服务，我们可以使用下面的命令来加速容器内的软件的下载安装速度：

```bash
http_proxy=http://host.docker.internal:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true update && \
http_proxy=http://host.docker.internal:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true install vim -y
```

## 容器方式运行

非常简单，一条命令：

```bash
docker run -d --name=apt-proxy -p 3142:3142 soulteary/apt-proxy
```

## 支持参数

我们可以通过使用 `-h` 参数来查看程序支持的所有参数：

```bash
./apt-proxy -h

Usage of apt-proxy:
  -cachedir string
    	the dir to store cache data in (default "./.aptcache")
  -debian string
    	the debian mirror for fetching packages
  -debug
    	whether to output debugging logging
  -host string
    	the host to bind to (default "0.0.0.0")
  -mode all
    	select the mode of system to cache: all / `ubuntu` / `debian` (default "all")
  -port string
    	the port to bind to (default "3142")
  -ubuntu string
    	the ubuntu mirror for fetching packages
```

## [待完善] 二次开发调试

项目由 Golang 编写，所以我们可以通过下面的命令快速启动开发模式的应用：

```bash
go run apt-proxy.go
```

### 运行单元测试，验证覆盖率

```bash
# go test -cover ./...

?   	github.com/soulteary/apt-proxy	[no test files]
ok  	github.com/soulteary/apt-proxy/cli	1.795s	coverage: 57.6% of statements
ok  	github.com/soulteary/apt-proxy/linux	5.771s	coverage: 80.0% of statements
ok  	github.com/soulteary/apt-proxy/pkg/httpcache	1.515s	coverage: 82.6% of statements
?   	github.com/soulteary/apt-proxy/pkg/httplog	[no test files]
ok  	github.com/soulteary/apt-proxy/pkg/stream.v1	0.935s	coverage: 100.0% of statements
?   	github.com/soulteary/apt-proxy/pkg/system	[no test files]
ok  	github.com/soulteary/apt-proxy/pkg/vfs	1.348s	coverage: 59.1% of statements
?   	github.com/soulteary/apt-proxy/server	[no test files]
ok  	github.com/soulteary/apt-proxy/state	0.356s	coverage: 84.2% of statements
```

查看覆盖率报告：

```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# go test -coverprofile=coverage.out ./...
PASS
coverage: 86.7% of statements
ok  	github.com/soulteary/apt-proxy	0.485s

# go tool cover -html=coverage.out
```

### Ubuntu / Debian 使用调试

```
http_proxy=http://192.168.33.1:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true update
http_proxy=http://192.168.33.1:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true install apache2
```

## 软件协议, 依赖组件

这个项目基于 [Apache License 2.0](https://github.com/soulteary/apt-proxy/blob/master/LICENSE)，相关的项目的协议清单如下：

- License NOT Found
    - [lox/apt-proxy](https://github.com/lox/apt-proxy#readme)
- MIT License
    - [lox/httpcache](https://github.com/lox/httpcache/blob/master/LICENSE)
    - [djherbis/stream](https://github.com/djherbis/stream/blob/master/LICENSE)
- Mozilla Public License 2.0
    - [rainycape/vfs](https://github.com/rainycape/vfs/blob/master/LICENSE)