package gcore

import (
	"fmt"
	"github.com/slclub/gcore/execute"
	"github.com/slclub/gnet"
	"github.com/slclub/grouter"
	"net/http"
	"testing"
)

func BenchmarkServer(B *testing.B) {
	app := createServer()
	addroute(app)
	run_request(B, app, "GET", "/my/ho/xiaoming/boy")
}

func BenchmarkServerHello(B *testing.B) {
	app := createServer()
	addroute(app)
	run_request(B, app, "GET", "/my/page")
}

func createServer() *Engine {
	en := New()

	router := grouter.NewRouter()
	router.SetKey("router")
	en.DriverRegister(router)

	// add Process
	mida := execute.NewMiddle("before_mid")
	en.DriverRegister(mida)

	process := &execute.Process{}
	process.SetKey("handle")
	en.DriverRegister(process)

	return en
}

func addroute(en *Engine) {
	dr, ok := en.DriverRouter("")
	if !ok {
		panic("not get router")
	}
	dr.GET("/my", func(ctx gnet.Contexter) {
		ctx.Response().WriteString("hello world!")
	})
	dr.GET("/my/ho/:uid/:sex", func(ctx gnet.Contexter) {
		s1, _ := ctx.Request().GetString("uid")
		s2, _ := ctx.Request().GetString("sex")
		//fmt.Println("xiaoming")
		ctx.Response().WriteString("hello world! " + s1 + ":" + s2)
	})
	dr.GET("/my/page", func(ctx gnet.Contexter) {
		ctx.Response().WriteString("hello world!")
	})

	//dr.ServerFile("/assets/", http.Dir("./assets"))
	dr.ServerFile("/st/:filepath", ("./assets/"), true)

	bm, ok := en.DriverMiddler("before_mid")
	bm.Use(func(ctx gnet.Contexter) {
		// fmt.Println("my first run middler ware func.")
	})
}

func run_request(B *testing.B, en *Engine, method, path string) {
	// create fake request
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("GCORE: we start benchmarks")
	w := new_mock_writer()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		en.Core().ServeHTTP(w, req)
	}
}

type header_writer struct {
	header http.Header
}

func (m *header_writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *header_writer) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *header_writer) Header() http.Header {
	return m.header
}

func (m *header_writer) WriteHeader(int) {}

func new_mock_writer() *header_writer {
	return &header_writer{
		http.Header{},
	}
}
