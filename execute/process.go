package execute

import (
	"github.com/slclub/gcore/flow"
	"github.com/slclub/gnet"
)

// most simple executer.
// run the handle offer by router.
type Process struct {
	flow.ExecuteNode
}

func (p *Process) Execute(ctx gnet.Contexter) {
	ctx.GetHandler()(ctx)
}
