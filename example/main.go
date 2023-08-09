package main

import (
	"context"
	"strconv"

	"github.com/mikelsr/raft-capnp/raft"
	ww "github.com/wetware/ww/wasm"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	self, err := ww.Init(ctx)
	if err != nil {
		panic(err)
	}
	defer self.Close()

	if len(self.Args) < 1 {
		panic("usage: ww cluster run raft_example.wasm 1")
	}

	processNumber, err := strconv.Atoi(self.Args[0])
	if err != nil {
		panic(err)
	}

	raftNode := raft.New().
		WithID(raft.DefaultID()).
		WithLogger(raft.DefaultLogger).
		WithStorage(raft.DefaultStorage()).
		WithRaftConfig(raft.DefaultConfig()).
		WithOnNewValue(raft.NilOnNewValue).
		WithRaftStore(raft.DefaultRaftStore).
		WithRaftNodeRetrieval( /* TODO */ retrieveRaftNode)
	go serveRaftRequests(ctx, self.Ctx, raftNode.Cap())
}
