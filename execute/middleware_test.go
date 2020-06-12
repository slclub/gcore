package execute

import (
	"github.com/slclub/gnet"
	"github.com/slclub/gnet/permission"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleNew(t *testing.T) {
	mid := NewMiddle("mid0")

	f1 := func(gnet.Contexter) {}
	mid.Use(f1)
	mid.Use(f1)
	assert.Equal(t, 1, mid.Size())
	scope := mid.Invoker().GetScopeById(0)
	assert.Equal(t, permission.SCOPE_USED, scope)
	mid.Deny(f1)
	assert.Equal(t, 1, mid.Size())
	scope = mid.Invoker().GetScopeById(0)
	assert.Equal(t, permission.SCOPE_UNUSED, scope)

	// assert panic
	assert.Panics(t, func() { mid.Use(nil) })

}

func TestMiddleExecute(t *testing.T) {
	mid := NewMiddle("mid0")

	count := 0
	f1 := func(ctx gnet.Contexter) {
		count++
	}
	mid.Use(f1)

	f2 := func(gnet.Contexter) {
		count++
	}
	mid.Use(f2)
	mid.Execute(gnet.NewContext())

	assert.Equal(t, 2, count)

	// combine.
	count = 0
	mid2 := NewMiddle("mid1")
	f3 := func(ctx gnet.Contexter) {
		count++
	}
	mid2.Use(f3)
	mid.Combine(mid2)
	mid.Combine(nil)
	mid.Execute(gnet.NewContext())
	assert.Equal(t, 3, count)

	// GetHandle
	count = 0
	for {
		_, name := mid.GetHandle(count)
		if name == "" {
			break
		}

		count++
	}

	assert.Equal(t, 3, mid.Size())
	assert.Equal(t, 3, count)
}

func TestProcess(t *testing.T) {
	mid := NewMiddle("mid0")

	count := 0
	f1 := func(ctx gnet.Contexter) {
		count++
	}
	mid.Use(f1)

	f2 := func(gnet.Contexter) {
		count++
	}
	mid.Use(f2)

	handle := func(ctx gnet.Contexter) {
		count++
	}
	ctx := gnet.NewContext()
	ctx.SetHandler(handle)

	proc := NewProcess()

	count = 0
	proc.Execute(ctx)
	assert.Equal(t, 1, count)

	ctx.SetExecute(mid)

	count = 0
	proc.Execute(ctx)
	assert.Equal(t, 3, count)

}
