# 指定基础镜像
FROM golang
 
# 维护人信息
MAINTAINER liu "344776209@qq.com"
 
# 工作目录，即执行go命令的目录
WORKDIR $GOPATH/src/myblog-gin
 
# 将本地内容添加到镜像指定目录
ADD . $GOPATH/src/myblog-gin
 
# 设置开启go mod
RUN go env -w GO111MODULE=on
# 设置go代理
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 构建go应用
RUN go build -mod=mod main.go
 
# 指定镜像内部服务监听的端口
EXPOSE 3000
 
# 镜像默认入口命令，即go编译后的可执行文件
ENTRYPOINT ["./main"]