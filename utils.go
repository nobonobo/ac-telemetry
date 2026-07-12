package telemetry

import "unicode/utf16"

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

type Telemetry struct {
	StaticInfo *StaticInfo `json:"static_info,omitempty"`
	Graphics   *Graphics   `json:"graphics,omitempty"`
	Physics    *Physics    `json:"physics,omitempty"`
}
