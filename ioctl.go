package ioctl

import (
	"syscall"
	"unsafe"
)

const (
	IOC_NONE  = 0
	IOC_WRITE = 1
	IOC_READ  = 2

	IOC_NRBITS   = 8
	IOC_TYPEBITS = 8
	IOC_SIZEBITS = 14
	IOC_DIRBITS  = 2

	IOC_NRSHIFT   = 0
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS
	IOC_DIRSHIFT  = IOC_SIZESHIFT + IOC_SIZEBITS
)

func IOC(dir, t, nr int, size uintptr) int {
	return (dir << IOC_DIRSHIFT) |
		(t << IOC_TYPESHIFT) |
		(nr << IOC_NRSHIFT) |
		(int(size) << IOC_SIZESHIFT)
}

func IO(t, nr int) int {
	return IOC(IOC_NONE, t, nr, uintptr(0))
}

func IOR(t, nr int, size uintptr) int {
	return IOC(IOC_READ, t, nr, size)
}

func IOW(t, nr int, size uintptr) int {
	return IOC(IOC_WRITE, t, nr, size)
}

func Ioctl(fd uintptr, req int, data unsafe.Pointer) (err syscall.Errno) {
	_, _, err = syscall.RawSyscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(data))
	return
}
