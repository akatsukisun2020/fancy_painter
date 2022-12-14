module github.com/akatsukisun2020/fancy_painter/mainline

go 1.16

require (
	github.com/akatsukisun2020/fancy_painter/proto/fancy_common v0.0.0-00010101000000-000000000000 // indirect
	github.com/alibabacloud-go/darabonba-openapi v0.1.18
	github.com/alibabacloud-go/dypnsapi-20170525 v1.0.6
	github.com/alibabacloud-go/tea v1.1.19
	github.com/alibabacloud-go/tea-utils v1.4.5
	github.com/alibabacloud-go/tea-xml v1.1.2 // indirect
	github.com/clbanning/mxj/v2 v2.5.6 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gomodule/redigo/redis v0.0.1
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	google.golang.org/api v0.30.0
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/akatsukisun2020/fancy_painter/proto/fancy_common => ./proto/fancy_common
