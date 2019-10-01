module github.com/mhconradt/pingpong-users-api

require (
	github.com/golang/protobuf v1.3.2
	github.com/infobloxopen/protoc-gen-gorm v0.18.0
	github.com/jinzhu/gorm v1.9.1
	github.com/lib/pq v1.2.0
	github.com/micro/go-micro v1.10.0
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64
)

require github.com/mhconradt/grpc-statuses v0.0.0

replace github.com/mhconradt/grpc-statuses => ../grpc-statuses

require github.com/mhconradt/proto/status v0.0.0

replace github.com/mhconradt/proto/status => ../proto/status

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190924004331-208c0a498538 // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/jinzhu/now v1.0.1 // indirect
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/mhconradt/proto/conversation v0.0.0
	github.com/mhconradt/proto/user v0.0.0
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	google.golang.org/grpc v1.22.1
)

replace github.com/mhconradt/proto/conversation => ../proto/conversation

replace github.com/mhconradt/proto/user => ../proto/user

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

go 1.13
