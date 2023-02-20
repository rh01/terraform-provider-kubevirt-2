// +build freebsd openbsd netbsd dragonfly darwin linux solaris

package remote

import (
	"errors"
	"io/ioutil"
	"os"

<<<<<<< HEAD
	"github.com/nxadm/tail"
	"golang.org/x/sys/unix"
=======
	"github.com/hpcloud/tail"
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
)

func NewOutputInterceptor() OutputInterceptor {
	return &outputInterceptor{}
}

type outputInterceptor struct {
	redirectFile *os.File
	streamTarget *os.File
	intercepting bool
	tailer       *tail.Tail
	doneTailing  chan bool
}

func (interceptor *outputInterceptor) StartInterceptingOutput() error {
	if interceptor.intercepting {
		return errors.New("Already intercepting output!")
	}
	interceptor.intercepting = true

	var err error

	interceptor.redirectFile, err = ioutil.TempFile("", "ginkgo-output")
	if err != nil {
		return err
	}

<<<<<<< HEAD
	// This might call Dup3 if the dup2 syscall is not available, e.g. on
	// linux/arm64 or linux/riscv64
	unix.Dup2(int(interceptor.redirectFile.Fd()), 1)
	unix.Dup2(int(interceptor.redirectFile.Fd()), 2)
=======
	// Call a function in ./syscall_dup_*.go
	// If building for everything other than linux_arm64,
	// use a "normal" syscall.Dup2(oldfd, newfd) call. If building for linux_arm64 (which doesn't have syscall.Dup2)
	// call syscall.Dup3(oldfd, newfd, 0). They are nearly identical, see: http://linux.die.net/man/2/dup3
	syscallDup(int(interceptor.redirectFile.Fd()), 1)
	syscallDup(int(interceptor.redirectFile.Fd()), 2)
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")

	if interceptor.streamTarget != nil {
		interceptor.tailer, _ = tail.TailFile(interceptor.redirectFile.Name(), tail.Config{Follow: true})
		interceptor.doneTailing = make(chan bool)

		go func() {
			for line := range interceptor.tailer.Lines {
				interceptor.streamTarget.Write([]byte(line.Text + "\n"))
			}
			close(interceptor.doneTailing)
		}()
	}

	return nil
}

func (interceptor *outputInterceptor) StopInterceptingAndReturnOutput() (string, error) {
	if !interceptor.intercepting {
		return "", errors.New("Not intercepting output!")
	}

	interceptor.redirectFile.Close()
	output, err := ioutil.ReadFile(interceptor.redirectFile.Name())
	os.Remove(interceptor.redirectFile.Name())

	interceptor.intercepting = false

	if interceptor.streamTarget != nil {
		interceptor.tailer.Stop()
		interceptor.tailer.Cleanup()
		<-interceptor.doneTailing
		interceptor.streamTarget.Sync()
	}

	return string(output), err
}

func (interceptor *outputInterceptor) StreamTo(out *os.File) {
	interceptor.streamTarget = out
}
