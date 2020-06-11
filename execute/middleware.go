package execute

import (
	// "fmt"
	"github.com/slclub/gcore/flow"
	"github.com/slclub/gnet"
	"github.com/slclub/gnet/permission"
	"github.com/slclub/utils"
)

type Middler interface {
	// public excuter interface.
	flow.IExecuteNode
	// middle ware interface
	Use(gnet.HandleFunc)
	Deny(gnet.HandleFunc)

	GetHandle(i int) (gnet.HandleFunc, string)
	Combine(Middler)
}

type MiddleWare struct {
	flow.ExecuteNode

	handle_chains []gnet.HandleFunc
	handle_names  []string
}

var _ Middler = &MiddleWare{}

func NewMiddle(name string) Middler {
	m := &MiddleWare{}
	m.SetKey(name)
	m.InitInvoker()
	return m
}

func (m *MiddleWare) Execute(ctx gnet.Contexter) {
	defer flow.RecoverNode()

	for i, handle := range m.handle_chains {
		// validate by invoker
		// id, ok := m.Invoker().GetId(m.handle_names[i])
		// fmt.Println("before middle", m.handle_names[i], id, ok, m.Invoker().Validate(0, nil))
		if !m.Invoker().ValidateByName(m.handle_names[i], ctx.GetAccess()) {
			continue
		}
		handle(ctx)
	}
}

func (m *MiddleWare) Use(handle gnet.HandleFunc) {
	name := utils.FUNC_NAME(handle)
	m.handle_chains = append(m.handle_chains, handle)
	m.handle_names = append(m.handle_names, name)

	m.Invoker().AutoSet(name, permission.SCOPE_USED)
}

func (m *MiddleWare) Deny(handle gnet.HandleFunc) {
	name := utils.FUNC_NAME(handle)
	for _, nm := range m.handle_names {
		if name == nm {
			m.Invoker().AutoSet(name, permission.SCOPE_UNUSED)
		}
	}
}

func (m *MiddleWare) GetHandle(i int) (gnet.HandleFunc, string) {
	if i >= len(m.handle_chains) {
		return nil, ""
	}
	return m.handle_chains[i], m.handle_names[i]
}

func (m *MiddleWare) Combine(one Middler) {
	if one == nil {
		return
	}
	i := 0
	for {
		handle, _ := one.GetHandle(i)
		if handle == nil {
			return
		}
		m.Use(handle)
		i++
	}
}
