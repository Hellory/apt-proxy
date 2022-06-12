# APT Proxy

Small and Reliable APT packages cache tool, supports both Ubuntu and Debian.

You can safely use it instead of [apt-cacher-ng](https://www.unix-ag.uni-kl.de/~bloch/acng/).

## (WIP) Usage

- Binaries
- Docker

## Development

coverage:

```bash
# go test -cover ./...

?   	github.com/soulteary/apt-proxy	[no test files]
ok  	github.com/soulteary/apt-proxy/cli	2.492s	coverage: 68.4% of statements
ok  	github.com/soulteary/apt-proxy/linux	7.420s	coverage: 76.7% of statements
ok  	github.com/soulteary/apt-proxy/pkgs/httpcache	2.575s	coverage: 82.7% of statements
?   	github.com/soulteary/apt-proxy/pkgs/httplog	[no test files]
ok  	github.com/soulteary/apt-proxy/pkgs/stream.v1	1.238s	coverage: 100.0% of statements
ok  	github.com/soulteary/apt-proxy/pkgs/vfs	2.255s	coverage: 59.4% of statements
?   	github.com/soulteary/apt-proxy/proxy	[no test files]
```

View coverage report:

```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# go test -coverprofile=coverage.out ./...
PASS
coverage: 86.7% of statements
ok  	github.com/soulteary/apt-proxy	0.485s

# go tool cover -html=coverage.out
```



### (WIP) Development

```bash
go run apt-proxy.go
```

## Ubuntu / Debian Debugging

```
http_proxy=http://192.168.33.1:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true update
http_proxy=http://192.168.33.1:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true install apache2
```

## Licenses, contains dependent software

This project is under the [Apache License 2.0](https://github.com/soulteary/apt-proxy/blob/master/LICENSE), and base on those software (or codebase).

- License NOT Found
    - [lox/apt-proxy](https://github.com/lox/apt-proxy#readme)
- MIT License
    - [lox/httpcache](https://github.com/lox/httpcache/blob/master/LICENSE)
    - [djherbis/stream](https://github.com/djherbis/stream/blob/master/LICENSE)
    - [stretchr/testify](https://github.com/stretchr/testify/blob/master/LICENSE)
- Mozilla Public License 2.0
    - [rainycape/vfs](https://github.com/rainycape/vfs/blob/master/LICENSE)
