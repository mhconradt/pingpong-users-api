run:
	go build
	./pingpong-users-api
proto:
	protoc --go_out=./proto/ --gorm_out=./proto/ --micro_out=./proto/ -I=${GOPATH}/src -I=./proto/ ./proto/users.proto
