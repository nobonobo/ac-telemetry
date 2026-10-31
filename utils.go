package telemetry

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procOpenFileMappingW = kernel32.NewProc("OpenFileMappingW")
)

func boolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// OpenFileMappingW のラッパー関数
// 戻り値のハンドルが INVALID_HANDLE_VALUE の場合はエラー
func OpenFileMappingW(
	dwDesiredAccess uint32,
	bInheritHandle bool,
	lpName string,
) (syscall.Handle, error) {
	var name16 *uint16
	var err error

	if lpName != "" {
		name16, err = syscall.UTF16PtrFromString(lpName)
		if err != nil {
			return 0, err
		}
	}

	r1, _, err := procOpenFileMappingW.Call(
		uintptr(dwDesiredAccess),
		uintptr(boolToUintptr(bInheritHandle)),
		uintptr(unsafe.Pointer(name16)),
	)

	if r1 == 0 {
		if err == syscall.Errno(0) {
			err = syscall.GetLastError()
		}
		return 0, err
	}

	return syscall.Handle(r1), nil
}

func getByteSlice(addr uintptr, byteSize int) []byte {
	if byteSize <= 0 {
		return nil
	}
	// Go 1.17+ で導入された便利関数
	return unsafe.Slice((*byte)(unsafe.Pointer(addr)), byteSize)
}

// DecodeString は [N]uint16 を Go の string に変換するユーティリティです。
func DecodeString(arr []uint16) string {
	i := 0
	for ; i < len(arr) && arr[i] != 0; i++ {
	}
	return string(utf16.Decode(arr[:i]))
}

// Coordinates は タイヤ接触点などの3D座標を表します。
type Coordinates struct {
	X float32
	Y float32
	Z float32
}
