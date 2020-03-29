# client-go-sample

## go mod
```shell script
go mod init 
go mod download
go mod tidy
go mod vendor
```

### step
```shell script
export GO111MODULE=on

# https://goproxy.cn/
export GOPROXY=https://goproxy.cn

go mod init

go get k8s.io/client-go@v0.17.4

```