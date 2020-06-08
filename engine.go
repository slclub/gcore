package gcore

import (
	"github.com/slclub/gcore/execute"
	"github.com/slclub/gnet/addr"
	"github.com/slclub/grouter"
	"github.com/slclub/link"
	"net"
	"net/http"
	"os"
)

type Engine struct {
	http_addr  addr.Addr
	https_addr addr.Addr
	web_socket addr.Addr
	core       *Core
}

func New() *Engine {
	en := &Engine{}
	en.core = NewCore()
	en.http_addr = addr.NewAddr(
		link.GetString("http.host", ""),
		link.GetString("http.port", ""),
		link.GetString("http.name", ""),
	)
	en.https_addr = addr.NewAddr(
		link.GetString("https.host", ""),
		link.GetString("https.port", ""),
		link.GetString("https.name", ""),
	)
	return en
}

// ****************************************serve*******************************************
func (en *Engine) Run() {
	en.Http()
	en.HttpTLS()
	en.UnixSock()
}

func (en *Engine) Http() {
	open := link.Config().GetBool("http.enable")
	if !open {
		return
	}
	err := http.ListenAndServe(en.http_addr.String(), en.core)
	link.INFO(err)
}

func (en *Engine) HttpTLS() {
	cert_file := link.GetString("https.cert_file", "")
	key_file := link.GetString("https.key_file", "")
	open := link.Config().GetBool("https.enable")
	if cert_file == "" || key_file == "" {
		return
	}

	if !open {
		return
	}

	err := http.ListenAndServeTLS(en.https_addr.String(), cert_file, key_file, en.core)
	link.INFO(err)
}

func (en *Engine) UnixSock() (err error) {
	defer link.INFO("[UNIX_SOCK][SERVE]", err)

	file := link.GetString("unix.sock_file", "")
	open := link.Config().GetBool("unix.enable")
	if file == "" || !open {
		return
	}
	listener, err := net.Listen("unix", file)
	if err != nil {
		return
	}

	defer listener.Close()
	defer os.Remove(file)
	err = http.Serve(listener, en.core)
	return
}

// TODO
func (en *Engine) WebSocket() (err error) {
	return
}

// ****************************************serve*******************************************

// ****************************************driver*******************************************

// register driver to execute flow
func (en *Engine) DriverRegister(driver interface{}) {
	en.core.GetFlow().Add(driver)
}

// get router driver
func (en *Engine) DriverRouter(name string) (grouter.Router, bool) {
	if name == "" {
		name = "router"
	}
	v := en.core.GetFlow().Get(name)
	r, ok := v.(grouter.Router)
	return r, ok
}

// get public type driver
// need to assert by yourself.
func (en *Engine) Driver(name string) interface{} {
	return en.core.GetFlow().Get(name)
}

// get middle ware driver.
func (en *Engine) DriverMiddler(name string) (execute.Middler, bool) {
	v := en.core.GetFlow().Get(name)
	dr, ok := v.(execute.Middler)
	return dr, ok
}
