module github.com/mikelsr/ww-raft-example

go 1.21

require (
	capnproto.org/go/capnp/v3 v3.0.0-alpha.30
	github.com/mikelsr/raft-capnp v0.0.0-20230806194035-41db6507c6f8
	github.com/wetware/ww v0.0.1-beta.30.0.20230715172926-5b96c77026c4
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/tetratelabs/wazero v1.2.1 // indirect
	go.etcd.io/raft/v3 v3.0.0-20230805183326-89c97ed7f982 // indirect
	golang.org/x/sync v0.3.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	zenhack.net/go/util v0.0.0-20230512035932-3f5fdfa0f47d // indirect
)

replace (
	github.com/mikelsr/raft-capnp v0.0.0-20230806194035-41db6507c6f8 => /home/mikel/Code/github.com/mikelsr/raft-capnp
	github.com/wetware/ww v0.0.1-beta.30.0.20230715172926-5b96c77026c4 => /home/mikel/Code/github.com/wetware/ww
)
