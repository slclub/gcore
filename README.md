# gcore
gcore serve http https and so on

## Content List

- [Benchmarks](#benchmarks)
- [Engine](#engine)
  - [ListenService](#listen-service)
  - [Driver](#driver)
- [Core](#core)
  - [ServeHTTP](#server-http)
  - [allocate]
- [Flow](#flow)
  - [Recovery](#recovery)
  - [ExecuteNode](#execute-node)
  - [Flow](#flow)
- [Execute](#execute)
  - [Process](#process)
  - [Middler](#middler)

## Benchmarks

```go
BenchmarkServer-4		 1000000	      1012 ns/op	     528 B/op	       7 allocs/op
PASS
ok		github.com/slclub/gcore	1.029s
```

## Engine

Listen service include http, https and websockets and unix file.

### Listen Service

- Http
- Https
- UnixSock
- WebSocket
- Run

### Driver

- DriverRegister(driver interface{})

Register a flow node

```go
    en := gcore.New()
    
    // add Process
    mida := execute.NewMiddle("before_mid")
    en.DriverRegister(mida)
```
- Driver(name string) interface{} 

Get any type of drvier node.

- DriverRouter(name string) (grouter.Router, bool) 

Get router type of driver node.

- DriverMiddler(name string) (execute.Middler, bool)

Get Middler driver node.

- OverAllocate func() gnet.Contexter

Rewrite core.allocate method.

## Core

### Server Http

- ServeHTTP(w http.ResponseWriter, req *http.Request) 
- allocate; internal method. alloc memory.

## Flow

Execute flow.

### Execute Node

Execute flow node base interface.

```go
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
```

### Flow


```go
type Flow struct {
	nodes []interface{}
}
```

### Recovery

- RecoverNode() 

Execute node recovery like try catch.

Panic of receiving node

- RecoverCore() 

Execute flow recovery like try catch.

Panic of receiving from the whole request.

## Execute

Actual execution nodes

### Process

Implement gnet.Executer interface.

Implement IExecuteNode interface.

```go
func (p *Process) Execute(ctx gnet.Contexter) {
	exec := ctx.GetExecute()
	if exec != nil {
		exec.Execute(ctx)
	}
	ctx.GetHandler()(ctx)
}
```

### Middler

Implement gnet.Executer interface. The core executive interface. The most important interface.

Implement execute.IExecuteNode interface.

Implement execute.Middler interface.

```go
type Middler interface {
	// public excuter interface.
	flow.IExecuteNode
	// middle ware interface
	Use(gnet.HandleFunc)
	Deny(gnet.HandleFunc)

	GetHandle(i int) (gnet.HandleFunc, string)
	Combine(Middler)
	Size() int
}
```
