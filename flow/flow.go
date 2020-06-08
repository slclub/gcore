package flow

import (
	//"fmt"
	"github.com/slclub/gerror"
	"github.com/slclub/gnet"
	"github.com/slclub/gnet/permission"
)

type LinkInvoker interface {
	Invoker() permission.Invoker
	InitInvoker()
}

type IExecuteNode interface {
	gnet.Executer
	LinkInvoker

	GetKey() string
	SetKey(string)
}

var _ IExecuteNode = &ExecuteNode{}

// *************************************EXE *****************************************
type ExecuteNode struct {
	// Unique tag name.
	key string

	// process chian rights.
	invoker permission.Invoker
}

func NewExe(name string) IExecuteNode {
	n := &ExecuteNode{}
	n.SetKey(name)
	n.InitInvoker()
	return n
}

func (en *ExecuteNode) Execute(ctx gnet.Contexter) {
}

func (en *ExecuteNode) InitInvoker() {
	en.invoker = permission.NewInvoke()
}

func (en *ExecuteNode) Invoker() permission.Invoker {
	return en.invoker
}

func (en *ExecuteNode) GetKey() string {
	return en.key
}

func (en *ExecuteNode) SetKey(key string) {
	if key == "" {
		gerror.Panic(gerror.CONST_ERRNO_PANIC, "ExecuteNode key is empty")
	}
	en.key = key
}

// *************************************Flow*****************************************
type Flow struct {
	nodes []interface{}
}

func NewFlow() *Flow {
	return &Flow{
		nodes: make([]interface{}, 0),
	}
}

func (f *Flow) Add(driver interface{}) {
	nnd, ok := driver.(IExecuteNode)
	if !ok {
		gerror.Panic(gerror.CONST_ERRNO_PANIC, "Execute node not implement IExecuteNode")
		return
	}
	for _, node := range f.nodes {
		nd, ok := node.(IExecuteNode)
		if ok && nd.GetKey() == nnd.GetKey() {
			gerror.Panic(gerror.CONST_ERRNO_PANIC, "Execute node key conflict. KEY:[", nnd.GetKey(), "]")
			return
		}
	}
	f.nodes = append(f.nodes, driver)
}

func (f *Flow) Get(name string) interface{} {
	for _, node := range f.nodes {
		nd, ok := node.(IExecuteNode)
		if ok && nd.GetKey() == name {
			//fmt.Println("name", name, "key", nd.GetKey())
			return node
		}
	}
	return nil
}

func (f *Flow) GetExe(name string) (IExecuteNode, bool) {
	for _, node := range f.nodes {
		nd, ok := node.(IExecuteNode)
		if ok && nd.GetKey() == name {
			//fmt.Println("name", name, "key", nd.GetKey())
			return nd, true
		}
	}
	return nil, false

}

func (f *Flow) Next(index *int) (IExecuteNode, bool) {
	if *index >= len(f.nodes) {
		return nil, false
	}
	nd, ok := f.nodes[*index].(IExecuteNode)
	*index++
	return nd, ok
}
