TOP=../../
cd ${TOP}

protoc --go_out=plugins=grpc:. gorpc/proto/echo.proto