package flow

import (
	"github.com/slclub/gerror"
	"github.com/slclub/gnet/defined"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecovery(t *testing.T) {

	assert.Panics(t, panic_recovery_node)

	assert.NotPanics(t, not_panic_recovery_core)
	assert.Panics(t, panic_recovery_core)
}

func TestNotRecover(t *testing.T) {
	defer RecoverCore()
	defer RecoverNode()

	panic("nothing need to panic")
}

func panic_recovery_node() {
	defer RecoverNode()

	func() {
		gerror.Panic(defined.CODE_JUMP_CURRENT_NODE, "jump node")
	}()
}

func panic_recovery_core() {
	defer RecoverCore()
	gerror.Panic(gerror.CONST_ERRNO_PANIC, "panic server")
}

func not_panic_recovery_core() {
	defer RecoverCore()
	gerror.Panic(defined.CODE_JUMP_CURRENT_NODE, "jump node")
}
