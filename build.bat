# 通过pb生成对应的go文件
protoc -I ./Proto --go_out=plugins=micro:./Proto ./Proto/consignment.proto
