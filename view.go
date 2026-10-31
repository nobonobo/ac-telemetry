package telemetry

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"
)

type View struct {
	mu     sync.RWMutex
	name   string
	handle syscall.Handle
	addr   unsafe.Pointer
}

func NewView(name string) *View {
	return &View{
		name: name,
	}
}

func (m *View) Open() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.open()
}

func (m *View) open() error {
	m.close()
	handle, err := OpenFileMappingW(syscall.FILE_MAP_READ, false, m.name)
	if err != nil {
		return err
	}
	m.handle = handle
	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, 0)
	if err != nil {
		return err
	}
	m.addr = unsafe.Pointer(addr)
	return nil
}

func (m *View) Close() (errs error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.close()
}

func (m *View) close() (errs error) {
	var err1, err2 error
	if m.addr != nil {
		err1 = syscall.UnmapViewOfFile(uintptr(m.addr))
		m.addr = nil
	}
	if m.handle != 0 {
		err2 = syscall.CloseHandle(m.handle)
		m.handle = 0
	}
	return errors.Join(err1, err2)
}

func (m *View) Ptr() unsafe.Pointer {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.addr == nil {
		if err := m.open(); err != nil {
			//log.Println("Failed to open view:", err)
			return nil
		}
	}
	return m.addr
}
