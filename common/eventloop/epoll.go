package eventloop

import (
	"syscall"
)

const EPOLL_ET = 1 << 31
const EPOLL_ERR = syscall.EPOLLERR
const EPOLL_READ = syscall.EPOLLIN
const EPOLL_WRITE = syscall.EPOLLOUT
const EPOLL_READWRITE = syscall.EPOLLIN | syscall.EPOLLOUT

var epfd int
var events = [512]syscall.EpollEvent{}

func epollInit() (err error) {
	epfd, err = syscall.EpollCreate1(syscall.EPOLL_CLOEXEC)
	return
}

func EpollAdd(fd int, event uint32) (err error) {
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &syscall.EpollEvent{
		Fd:     int32(fd),
		Events: EPOLL_ET | syscall.EPOLLERR | event,
	})
	return
}

func EpollModify(fd int, event uint32) (err error) {
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_MOD, fd, &syscall.EpollEvent{
		Fd:     int32(fd),
		Events: EPOLL_ET | syscall.EPOLLERR | event,
	})
	return
}

func EpollDelete(fd int) (err error) {
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_DEL, fd, &syscall.EpollEvent{
		Fd: int32(fd),
	})
	return
}

func EpollWait() (ev []syscall.EpollEvent, err error) {
	if n, err := syscall.EpollWait(epfd, events[:], -1); err != nil || n <= 0 {
		return nil, err
	} else {
		return events[:n], nil
	}
}
