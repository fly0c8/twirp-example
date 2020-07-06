protoc --proto_path=$GOPATH/src:. \
--twirp_out=. \
--go_out=. \
--twirp_swagger_out=. \
service.proto