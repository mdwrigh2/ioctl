package ioctl

import (
	"syscall"
	"unsafe"
)

const (
	_IOC_NONE  = 0
	_IOC_WRITE = 1
	_IOC_READ  = 2

	_IOC_NRBITS   = 8
	_IOC_TYPEBITS = 8
	_IOC_SIZEBITS = 14
	_IOC_DIRBITS  = 2

	_IOC_NRSHIFT   = 0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS
)

func _IOC(dir, t, nr int, size uintptr) int {
	return (dir << _IOC_DIRSHIFT) |
		(t << _IOC_TYPESHIFT) |
		(nr << _IOC_NRSHIFT) |
		(int(size) << _IOC_SIZESHIFT)
}

func _IO(t, nr int) int {
	return _IOC(_IOC_NONE, t, nr, uintptr(0))
}

func _IOR(t, nr int, size uintptr) int {
	return _IOC(_IOC_READ, t, nr, size)
}

func _IOW(t, nr int, size uintptr) int {
	return _IOC(_IOC_WRITE, t, nr, size)
}

func Ioctl(fd uintptr, req int, data unsafe.Pointer) (err syscall.Errno) {
	_, _, err = syscall.RawSyscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(data))
	return
}
