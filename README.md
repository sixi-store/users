# users
用户系统

### 交叉编译
生成linux系统可执行文件: `CGO_ENABLED=0 GOOS=linux go build .`

### 生成pd.go
1.到pb目录
2.`protoc --go_out=plugins=grpc:.  --proto_path=../proto  user.proto`
