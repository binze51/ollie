package utils

import (
	"context"
	"runtime"
	"runtime/debug"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Recover(ctx context.Context) {
	if p := recover(); p != nil {
		HandleThrow(ctx, p)
	}
}

func HandleThrow(ctx context.Context, p any) {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	klog.CtxErrorf(ctx, "HandleThrow|func=%s|error=%#v|stack=%s\n", f, p, string(debug.Stack()))
	// os.Exit(-1)
}

func Go(ctx context.Context, f func()) {
	go func() {
		defer Recover(ctx)
		f()
	}()
}
