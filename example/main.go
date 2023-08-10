package main

import (
	"context"
	"strconv"

	"capnproto.org/go/capnp/v3"
	"github.com/mikelsr/raft-capnp/raft"
	"github.com/wetware/ww/api/cluster"
	"github.com/wetware/ww/pkg/csp"
	ww "github.com/wetware/ww/wasm"
)

const totalNodes uint64 = 3

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

	processNumber, err := strconv.ParseUint(self.Args[0], 10, 64)
	if err != nil {
		panic(err)
	}

	// The host points to its executor.
	host := cluster.Host(self.Caps[0])
	executor, err := executorFromHost(ctx, host)
	if err != nil {
		panic(err)
	}
	defer executor.Release()

	raftNode := raft.New().
		WithID(processNumber).
		WithLogger(raft.DefaultLogger).
		WithStorage(raft.DefaultStorage()).
		WithRaftConfig(raft.DefaultConfig()).
		WithOnNewValue(raft.NilOnNewValue).
		WithRaftStore(raft.DefaultRaftStore).
		WithRaftNodeRetrieval( /* TODO */ retrieveRaftNode)
	go serveRaftRequests(ctx, self.Ctx, raftNode.Cap())
	go raftNode.Start(ctx)

	peers := make(map[uint64]csp.Proc)

	// First node will spawn more nodes
	if processNumber == 1 {
		for i := processNumber; i < totalNodes; i++ {
			proc, release := executor.ExecFromCache(
				ctx,
				[]byte(self.Hash),
				self.Pid,
				capnp.Client(csp.NewArgs(strconv.FormatUint(uint64(i), 10))),
				capnp.Client(host),
			)
			defer release()
			peers[i] = proc
		}
	}

}
