package main

import (
	"context"

	"github.com/wetware/ww/api/cluster"
	"github.com/wetware/ww/pkg/csp"
)

func executorFromHost(ctx context.Context, host cluster.Host) (csp.Executor, error) {
	f, _ := host.Executor(ctx, nil)
	<-f.Done()

	res, err := f.Struct()
	if err != nil {
		return csp.Executor{}, err
	}

	return csp.Executor(res.Executor()), nil
}
