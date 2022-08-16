
# 使用方式：
# sh protoc_gen.sh ${输入的协议的目录名称}
# 如： "sh protoc_gen.sh fancy_painter" 是是生成 fancy_painter 目录下的协议文件的桩代码
# 注意的是：输入的目录名称和协议文件的前缀需要一致。比如 fancy_painter 和 fancy_painter.proto


# 1. 参数指定输入协议文件
INPUT_PROTO=$1 #第一个参数为输入的协议文件
INPUT_PROTO_FILE=${INPUT_PROTO}.proto

# 2. 依赖文件拷贝 [这个很重要]
cp ../pkg/google ${INPUT_PROTO} -r       # 外部google库

ls ${INPUT_PROTO}
echo '输入proto文件: '$INPUT_PROTO_FILE
echo '开始生成pb协议文件...'

# 3. 用于生成pb文件
protoc --proto_path=${INPUT_PROTO} --proto_path=./ --go_out=plugins=grpc:. ${INPUT_PROTO_FILE}
# 4. 用于生成pb文件，pb.go文件， gateway文件
protoc --proto_path=${INPUT_PROTO} --proto_path=./ --go_out . --go-grpc_out . --grpc-gateway_out . ${INPUT_PROTO_FILE}
# protoc --go_out=plugins=grpc:. --grpc-gateway_out . ${INPUT_PROTO_FILE}

# 5. 依赖文件清理
rm -rf ${INPUT_PROTO}/google

echo '生成pb协议文件完成...'


## FIXME:
## 公共代码生成之后，目前还需要手动拷贝，不是很灵活！！ ==> 以后在优化吧 !! 

