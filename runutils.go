package runutils

import (
	"fmt"
	"math"
	"path"
	"runtime"
	"time"
)

// 获取正在运行的函数名
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func RunFileAndLine() (info string) {
	_, name, line, _ := runtime.Caller(1)
	baseName := path.Base(name)
	return fmt.Sprintf("%v#%v", baseName, line)
}

func CurrentTime() string {
	t := time.Now()
	ms := int64(math.Mod(float64(t.UnixNano()), float64(time.Second))) / int64(time.Millisecond)
	return fmt.Sprintf("%v-%v-%v %v:%v:%v:%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), ms)
}
