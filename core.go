package gcore

import (
	"github.com/slclub/gcore/flow"
	"github.com/slclub/gnet"
	"net/http"
	"sync"
)

type Core struct {
	run_flow *flow.Flow
	pool     sync.Pool
}

func NewCore() *Core {
	core := &Core{}
	core.run_flow = flow.NewFlow()
	core.pool.New = func() interface{} {
		return core.allocate()
	}
	return core
}

func (core *Core) allocate() gnet.Contexter {
	ctx := gnet.NewContext()
	r := gnet.NewRequest()
	s := &gnet.Response{}
	ctx.SetRequest(r)
	ctx.SetResponse(s)
	return ctx
}

// ServeHTTP conforms to the http.Handler interface.
func (core *Core) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	defer flow.RecoverCore()
	ctx := core.pool.Get().(gnet.Contexter)
	ctx.Request().InitWithHttp(req)
	ctx.Response().InitSelf(w)

	index := 0
	for {
		exe, ok := core.run_flow.Next(&index)
		if !ok {
			break
		}
		exe.Execute(ctx)
	}

	ctx.Reset()
	core.pool.Put(ctx)

}

func (core *Core) GetFlow() *flow.Flow {
	return core.run_flow
}
