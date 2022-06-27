package multiroutine

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func MultiAdd() error {
	infos := []string{
		"100",
		"20",
		"30", "50",
	}
	var errLog error
	var wg sync.WaitGroup
	fmt.Println("CPU NUM", runtime.NumCPU())
	filedatas := make([]string, len(infos))
	//routines := make(chan struct{}, runtime.NumCPU())
	routines := make(chan struct{}, 2)
	for index, info := range infos {
		if errLog != nil {
			break
		}
		routines <- struct{}{}
		wg.Add(1)
		go func(index int, fileName string) {
			defer func() {
				wg.Done()
				<-routines
				if err := recover(); err != nil { //todo:实现时添加panic打印
					//var buf [4096]byte
					//n := runtime.Stack(buf[:], false)
					//logkit.Error("uploadMulitFileToUssS3 panic", zap.String("panicInfo", string(buf[:n])), zap.String("traceId", traceId))
					errLog = errors.New("uploadMulitFileToUssS3 panic")
				}
			}()
			//todo
			time.Sleep(time.Second)
			fmt.Println("index:", index, "fileName:", fileName)
			filedatas[index] = fileName
		}(index, info)
	}
	wg.Wait()
	fmt.Println(filedatas)
	return errLog
}
