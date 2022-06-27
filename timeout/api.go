package timeout

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

func CmdTimeOut(cmd string, ts time.Duration) {
	now := time.Now()
	exe := exec.Command("sh", "-c", "time sleep 10")

	exec.Command("sh", cmd)

	var stderr bytes.Buffer
	var stdout bytes.Buffer
	exe.Stderr = &stderr
	exe.Stdout = &stdout

	exe.SysProcAttr = &syscall.SysProcAttr{
		//ParentProcess: 10, //for windows todo
		Setpgid: true, //for linux
	}
	var finish = make(chan struct{}, 1)
	//todo: time.After使用需要了解开启和关闭
	go func() {
		select {
		case <-finish:
		case <-time.After(5 * time.Second):
			//syscall.Kill(-exe.Process.Pid, syscall.SIGKILL)  for unix
		}
	}()
	err := exe.Run()
	finish <- struct{}{}

	stdErr := stderr.String()
	stdOut := stdout.String()
	fmt.Println(stdOut, stdErr)
	fmt.Println("time:", time.Since(now))
	if err != nil {
		fmt.Println(err)
	}
}
