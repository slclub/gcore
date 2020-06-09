package flow

import (
	"github.com/slclub/gerror"
	//"fmt"
	"github.com/slclub/gnet/defined"
	"github.com/slclub/link"
	"strconv"
	"strings"
)

func RecoverCore() {
	err := recover()
	if err == nil {
		return
	}

	debug := link.GetString("debug", "debug")
	s, ok := err.(string)
	if !ok {
		return
	}
	link.ERROR(s)
	if debug != "debug" {
		return
	}

	if debug == "panic" {
		panic(s)
	}

	i := strings.IndexByte(s, ':')
	if i <= 0 {
		return
	}

	code, er := strconv.Atoi(s[:i])

	if er != nil {
		return
	}
	if code == gerror.CONST_ERRNO_PANIC {
		panic(s)
	}
}

func RecoverNode() {
	err := recover()
	if err == nil {
		return
	}

	debug := link.GetString("debug", "debug")
	s, ok := err.(string)
	if !ok {
		return
	}
	if debug == "panic" {
		panic(s)
	}

	i := strings.IndexByte(s, ':')
	if i <= 0 {
		link.ERROR(s)
		return
	}

	code, er := strconv.Atoi(s[:i])

	if er != nil {
		return
	}
	switch code {
	case defined.CODE_JUMP_CURRENT_NODE:
		panic(s)
	case gerror.CONST_ERRNO_PANIC:
		panic(s)
	default:
		link.ERROR(s)
	}
	//fmt.Println(code)
}
