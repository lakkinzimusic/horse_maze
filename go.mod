module github.com/lakkinzimusic/horse_maze_micro/horse_maze

go 1.13

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/go-log/log v0.2.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/lakkinzimusic/horse_maze v0.0.0-20200511165125-6534bca6cfb6
	github.com/lakkinzimusic/shippy-service-vessel v0.0.0-20200516185548-72a586cc34db
	github.com/micro/go-micro v1.18.0
	github.com/miekg/dns v1.1.29 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	google.golang.org/grpc v1.29.1
)
