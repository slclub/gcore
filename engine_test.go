package gcore

import (
	"github.com/slclub/gcore/execute"
	"github.com/slclub/gnet"
	"github.com/slclub/grouter"
	"github.com/slclub/link"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEnginFake(t *testing.T) {
	en := New()
	en.Start()

	assert.Equal(t, "127.0.0.1:8080", en.http_addr.String())

	// add router Driver
	router := grouter.NewRouter()
	router.SetKey("router")
	en.DriverRegister(router)

	// add Process
	mida := execute.NewMiddle("before_mid")
	en.DriverRegister(mida)

	process := execute.NewProcess()
	en.DriverRegister(process)

	// test driver get
	_, ok := en.DriverMiddler("before_mid")
	assert.True(t, ok)
	da := en.Driver("before_mid")
	_, ok = da.(execute.Middler)
	assert.True(t, ok)

	router, ok = en.DriverRouter("")
	assert.Equal(t, "router", router.GetKey())
	//assert.Equal(t, "before_mid", mifaa.GetKey())

	// fake http request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Ping/xiaoming", nil)
	en.Core().ServeHTTP(w, req)
	en.core.pool.New()
	//en.Run()
}

func TestEngineWithConfig(t *testing.T) {

	en := New()
	en.Start()
	en.HttpTLS()
	en.UnixSock()

	link.Config().Set("http.enable", false)
	en.Http()
	en.HttpTLS()
	en.UnixSock()
	en.WebSocket()
	en.core.GetFlow()

	en.OverAllocate = func() gnet.Contexter {
		ctx := gnet.NewContext()
		r := gnet.NewRequest()
		s := &gnet.Response{}
		ctx.SetRequest(r)
		ctx.SetResponse(s)
		return ctx
	}
	en.core.pool.New()
	assert.NotNil(t, en.https_addr)
}

func TestEngineRun(t *testing.T) {
	en := New()
	en.Start()

	func() {
		en.Run()
	}()
	time.Sleep(4 * time.Second)
}
