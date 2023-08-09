package main

import (
	"context"

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
