package evo

import (
	telemetry "github.com/nobonobo/ac-telemetry"
)

type StaticInfoView struct {
	*telemetry.View
}

func NewStaticInfo() *StaticInfoView {
	return &StaticInfoView{View: telemetry.NewView(`Local\acevo_pmf_static`)}
}

func (v *StaticInfoView) Get() *StaticInfo {
	return (*StaticInfo)(v.Ptr())
}

// StaticInfo は Assetto Corsa の acpmf_static 共有メモリ構造体を表します。
// セッション開始時のみ更新される（*コメントの項目のみ）
// フィールド順序・型は AssettoCorsaRallyTelemetryReader の StaticInfo.cs
// (StructLayout Pack=4, CharSet=Unicode) に完全準拠しています。
type StaticInfo struct {
	SMVersion                [15]uint16 `json:"sm_version"`                            // SMバージョン
	ACVersion                [15]uint16 `json:"ac_version"`                            // ACバージョン
	NumberOfSessions         int32      `json:"number_of_sessions,omitempty"`          // セッション数
	NumCars                  int32      `json:"num_cars,omitempty"`                    // 車両数
	CarModel                 [33]uint16 `json:"car_model"`                             // *カーモデル名
	Track                    [33]uint16 `json:"track"`                                 // トラック
	PlayerName               [33]uint16 `json:"player_name"`                           // プレイヤー名
	PlayerSurname            [33]uint16 `json:"player_surname"`                        // プレイヤー姓
	PlayerNick               [33]uint16 `json:"player_nick"`                           // プレイヤーニックネーム
	SectorCount              int32      `json:"sector_count,omitempty"`                // セクター数
	MaxTorque                float32    `json:"max_torque,omitempty"`                  // 最大トルク
	MaxPower                 float32    `json:"max_power,omitempty"`                   // 最大パワー
	MaxRpm                   int32      `json:"max_rpm,omitempty"`                     // 最大回転数
	MaxFuel                  float32    `json:"max_fuel,omitempty"`                    // 最大燃料
	SuspensionMaxTravel      [4]float32 `json:"suspension_max_travel"`                 // サスペンション最大トラベル
	TyreRadius               [4]float32 `json:"tyre_radius"`                           // タイヤ半径
	MaxTurboBoost            float32    `json:"max_turbo_boost,omitempty"`             // 最大ターボブースト
	Deprecated_1             float32    `json:"deprecated_1,omitempty"`                // 廃止
	Deprecated_2             float32    `json:"deprecated_2,omitempty"`                // 廃止
	PenaltiesEnabled         int32      `json:"penalties_enabled,omitempty"`           // ペナルティ有効
	AidFuelRate              float32    `json:"aid_fuel_rate,omitempty"`               // 燃料アシスト
	AidTireRate              float32    `json:"aid_tire_rate,omitempty"`               // タイヤアシスト
	AidMechanicalDamage      float32    `json:"aid_mechanical_damage,omitempty"`       // メカニカルダメージアシスト
	AidAllowTyreBlankets     int32      `json:"aid_allow_tyre_blankets,omitempty"`     // タイヤブランケット許可
	AidStability             float32    `json:"aid_stability,omitempty"`               // スタビリティアシスト
	AidAutoClutch            int32      `json:"aid_auto_clutch,omitempty"`             // オートクラッチアシスト
	AidAutoBlip              int32      `json:"aid_auto_blip,omitempty"`               // オートブリップアシスト
	HasDRS                   int32      `json:"has_drs,omitempty"`                     // DRS有効
	HasERS                   int32      `json:"has_ers,omitempty"`                     // ERS有効
	HasKERS                  int32      `json:"has_kers,omitempty"`                    // KERS有効
	KersMaxJ                 float32    `json:"kers_max_j,omitempty"`                  // KERS最大エネルギー
	EngineBrakeSettingsCount int32      `json:"engine_brake_settings_count,omitempty"` // エンジンブレーキ設定数
	ErsPowerControllerCount  int32      `json:"ers_power_controller_count,omitempty"`  // ERSパワーコントロール数
	TrackSplineLength        float32    `json:"track_spline_length"`                   // *トラックの標準長 m
	TrackConfiguration       [33]uint16 `json:"track_configuration"`                   // トラックコンフィギュレーション
	ErsMaxJ                  float32    `json:"ers_max_j,omitempty"`                   // ERS最大エネルギー
	IsTimedRace              int32      `json:"is_timed_race,omitempty"`               // タイムレース
	HasExtraLap              int32      `json:"has_extra_lap,omitempty"`               // エキストララップ
	CarSkin                  [33]uint16 `json:"car_skin"`                              // 車両スキン
	ReversedGridPositions    int32      `json:"reversed_grid_positions,omitempty"`     // リバースグリッドポジション
	PitWindowStart           int32      `json:"pit_window_start,omitempty"`            // ピットウィンドウ開始
	PitWindowEnd             int32      `json:"pit_window_end,omitempty"`              // ピットウィンドウ終了
	IsOnline                 int32      `json:"is_online,omitempty"`                   // オンライン
	DryTyresName             [33]uint16 `json:"dry_tyres_name"`                        // ドライタイヤ名
	WetTyresName             [33]uint16 `json:"wet_tyres_name"`                        // ウェットタイヤ名
}

func (s *StaticInfo) SMVersionString() string {
	return telemetry.DecodeString(s.SMVersion[:])
}

func (s *StaticInfo) ACVersionString() string {
	return telemetry.DecodeString(s.ACVersion[:])
}

func (s *StaticInfo) CarModelString() string {
	return telemetry.DecodeString(s.CarModel[:])
}

func (s *StaticInfo) TrackString() string {
	return telemetry.DecodeString(s.Track[:])
}

func (s *StaticInfo) PlayerNameString() string {
	return telemetry.DecodeString(s.PlayerName[:])
}

func (s *StaticInfo) PlayerSurnameString() string {
	return telemetry.DecodeString(s.PlayerSurname[:])
}

func (s *StaticInfo) PlayerNickString() string {
	return telemetry.DecodeString(s.PlayerNick[:])
}

func (s *StaticInfo) TrackConfigurationString() string {
	return telemetry.DecodeString(s.TrackConfiguration[:])
}

func (s *StaticInfo) CarSkinString() string {
	return telemetry.DecodeString(s.CarSkin[:])
}

func (s *StaticInfo) DryTyresNameString() string {
	return telemetry.DecodeString(s.DryTyresName[:])
}

func (s *StaticInfo) WetTyresNameString() string {
	return telemetry.DecodeString(s.WetTyresName[:])
}
