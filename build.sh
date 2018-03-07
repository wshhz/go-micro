# 通过pb生成对应的go文件
for file in ./Proto/*
do
    if [ -d $file ]
    then
        echo $file
        protoc -I $file --go_out=plugins=micro:$file $file/*.proto
    fi
done

# 如果有错误信息,则不退出界面
if [ $? != 0 ]
then
	read -p "Press any key to continue"
fi