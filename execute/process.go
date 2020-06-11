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

func NewProcess() *Process {
	p := &Process{}
	p.SetKey("handle")
	//p.InitInvoker()
	return p
}

func (p *Process) Execute(ctx gnet.Contexter) {
	exec := ctx.GetExecute()
	if exec != nil {
		exec.Execute(ctx)
	}
	ctx.GetHandler()(ctx)
}
