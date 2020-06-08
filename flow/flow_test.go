package flow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	n := NewExe("nana")
	assert.Equal(t, "nana", n.GetKey())

	assert.Panics(t, func() { NewExe("") })
	n.Invoker()
}

func TestFlow(t *testing.T) {
	f := NewFlow()
	f.Add(NewExe("node1"))
	f.Add(NewExe("node2"))

	assert.Panics(t, func() { f.Add(NewExe("node1")) })
	assert.Panics(t, func() { f.Add(nil) })
	assert.Panics(t, func() { f.Add(25) })

	n := f.Get("node1")
	vn, ok := n.(IExecuteNode)
	assert.True(t, ok)
	assert.Equal(t, "node1", vn.GetKey())

	n = f.Get("noden")
	vn, ok = n.(IExecuteNode)
	assert.False(t, ok)
	assert.Nil(t, vn)

	vn, ok = f.GetExe("node2")
	assert.True(t, ok)
	assert.Equal(t, "node2", vn.GetKey())

	vn, ok = f.GetExe("noden")
	assert.False(t, ok)
	assert.Nil(t, vn)

	index := 0
	for {
		vn, ok = f.Next(&index)
		if !ok {
			break
		}
	}

	assert.Equal(t, 2, index)
}
