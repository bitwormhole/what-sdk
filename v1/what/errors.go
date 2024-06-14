package what

import (
	"fmt"

	"github.com/starter-go/vlog"
)

// HandleError ...
func HandleError(err error) {
	if err == nil {
		return
	}
	vlog.Error(err.Error())
}

// HandlePanic ...
func HandlePanic(x any) {

	if x == nil {
		return
	}

	err, ok := x.(error)
	if ok {
		HandleError(err)
		return
	}

	str, ok := x.(string)
	if ok {
		HandleError(fmt.Errorf(str))
		return
	}

	vlog.Error("what.HandlePanic(): unknow typeof panic")
}
