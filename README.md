# gcore
gcore serve http https and so on

## Content List

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

### 
