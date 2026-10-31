package evo

import telemetry "github.com/nobonobo/ac-telemetry"

type GraphicsView struct {
	*telemetry.View
}

func NewGraphics() *GraphicsView {
	return &GraphicsView{View: telemetry.NewView(`Local\acevo_pmf_graphics`)}
}

func (v *GraphicsView) Get() *Graphics {
	return (*Graphics)(v.Ptr())
}

// ACStatus は AC の動作状態を表します。
type ACStatus int32

const (
	ACOff    ACStatus = 0
	ACReplay ACStatus = 1
	ACLive   ACStatus = 2
	ACPause  ACStatus = 3
)

// ACSessionType はセッションの種類を表します。
type ACSessionType int32

const (
	ACSessionUnknown     ACSessionType = -1
	ACSessionPractice    ACSessionType = 0
	ACSessionQualify     ACSessionType = 1
	ACSessionRace        ACSessionType = 2
	ACSessionHotlap      ACSessionType = 3
	ACSessionTimeAttack  ACSessionType = 4
	ACSessionDrift       ACSessionType = 5
	ACSessionDrag        ACSessionType = 6
	ACSessionHotStint    ACSessionType = 7
	ACSessionHotlapSuper ACSessionType = 8
)

// ACFlagType はフラグの種類を表します。
type ACFlagType int32

const (
	ACNoFlag        ACFlagType = 0
	ACBlueFlag      ACFlagType = 1
	ACYellowFlag    ACFlagType = 2
	ACBlackFlag     ACFlagType = 3
	ACWhiteFlag     ACFlagType = 4
	ACCheckeredFlag ACFlagType = 5
	ACPenaltyFlag   ACFlagType = 6
)

// Graphics は Assetto Corsa の acpmf_graphics 共有メモリ構造体を表します。
// 描画時のみ更新される（*コメントの項目のみ）
// フィールド順序・型は acc-telemetry-rs の PageFileGraphic (shared_memory.rs L287)
// に完全準拠しています。Do NOT reorder fields.
type Graphics struct {
	PacketId              int32                     `json:"packet_id"`                         // *パケット番号
	Status                int32                     `json:"status,omitempty"`                  // ACStatus
	Session               int32                     `json:"session,omitempty"`                 // ACSessionType
	CurrentTime           [15]uint16                `json:"current_time"`                      // wchar_t[15] — 現在ラップタイム文字列
	LastTime              [15]uint16                `json:"last_time"`                         // wchar_t[15] — 前ラップタイム文字列
	BestTime              [15]uint16                `json:"best_time"`                         // wchar_t[15] — ベストラップタイム文字列
	Split                 [15]uint16                `json:"split"`                             // wchar_t[15] — 区間タイム文字列
	CompletedLaps         int32                     `json:"completed_laps,omitempty"`          // 完了ラップ数
	Position              int32                     `json:"position,omitempty"`                // 順位
	ICurrentTime          int32                     `json:"current_time_ms,omitempty"`         // 現在ラップタイム (ms)
	ILastTime             int32                     `json:"last_time_ms,omitempty"`            // 前ラップタイム (ms)
	IBestTime             int32                     `json:"best_time_ms,omitempty"`            // ベストラップタイム (ms)
	SessionTimeLeft       float32                   `json:"session_time_left,omitempty"`       // セッション残り時間 (秒)
	DistanceTraveled      float32                   `json:"distance_traveled"`                 // *走行距離 m
	IsInPit               int32                     `json:"is_in_pit,omitempty"`               // ピットイン中 (0/1)
	CurrentSectorIndex    int32                     `json:"current_sector_index,omitempty"`    // 現在のセクターインデックス
	LastSectorTime        int32                     `json:"last_sector_time_ms,omitempty"`     // 前セクタータイム (ms)
	NumberOfLaps          int32                     `json:"number_of_laps,omitempty"`          // セッション総ラップ数
	TyreCompound          [33]uint16                `json:"tyre_compound"`                     // wchar_t[33] — タイヤコンパウンド名
	ReplayTimeMultiplier  float32                   `json:"replay_time_multiplier,omitempty"`  // リプレイ速度倍率
	NormalizedCarPosition float32                   `json:"normalized_car_position"`           // コース上の正規化位置 (0.0–1.0)
	ActiveCars            int32                     `json:"active_cars"`                       // アクティブ車両数
	CarCoordinates        [60]telemetry.Coordinates `json:"-"`                                 // *全車両3D座標 [car_id][X,Y,Z] (最初のエントリのみ有効)
	CarId                 [60]int32                 `json:"-"`                                 // 全車両ID
	PlayerCarId           int32                     `json:"player_car_id"`                     // *プレイヤー車両ID(常にゼロ)
	PenaltyTime           float32                   `json:"penalty_time,omitempty"`            // *ペナルティ時間 (秒)
	Flag                  int32                     `json:"flag,omitempty"`                    // ACFlagType
	Penalty               int32                     `json:"penalty,omitempty"`                 // ペナルティ種別
	IdealLineOn           int32                     `json:"ideal_line_on,omitempty"`           // 理想ライン表示 (0/1)
	IsInPitLane           int32                     `json:"is_in_pit_lane,omitempty"`          // ピットレーン走行中 (0/1)
	SurfaceGrip           float32                   `json:"surface_grip,omitempty"`            // 路面グリップ係数
	MandatoryPitDone      int32                     `json:"mandatory_pit_done,omitempty"`      // 義務ピット完了 (0/1)
	WindSpeed             float32                   `json:"wind_speed,omitempty"`              // 風速 m/s
	WindDirection         float32                   `json:"wind_direction,omitempty"`          // 風向 (rad)
	IsSetupMenuVisible    int32                     `json:"is_setup_menu_visible,omitempty"`   // セットアップメニュー表示中 (0/1)
	MainDisplayIndex      int32                     `json:"main_display_index,omitempty"`      // メインディスプレイインデックス
	SecondaryDisplayIndex int32                     `json:"secondary_display_index,omitempty"` // サブディスプレイインデックス
	TC                    int32                     `json:"tc,omitempty"`                      // *TCレベル
	TCCut                 int32                     `json:"tccut,omitempty"`                   // TCカットレベル
	EngineMap             int32                     `json:"engine_map,omitempty"`              // エンジンマップ
	Abs                   int32                     `json:"abs,omitempty"`                     // *ABSレベル
	FuelXLap              int32                     `json:"fuel_x_lap,omitempty"`              // 燃料 x ラップ
	RainLights            int32                     `json:"rain_lights,omitempty"`             // レインライト点灯 (0/1)
	FlashingLights        int32                     `json:"flashing_lights,omitempty"`         // フラッシングライト (0/1)
	LightsStage           int32                     `json:"lights_stage,omitempty"`            // ライトステージ
	ExhaustTemperature    float32                   `json:"exhaust_temperature,omitempty"`     // 排気温度 K
	WiperLv               int32                     `json:"wiper_lv,omitempty"`                // ワイパーレベル
	DriverStintTotalLeft  int32                     `json:"driver_stint_total_left,omitempty"` // ドライバースティント総残り時間 (ms)
	DriverStintLeft       int32                     `json:"driver_stint_left,omitempty"`       // ドライバースティント残り時間 (ms)
	RainTyres             int32                     `json:"rain_tyres,omitempty"`              // レインタイヤ装着 (0/1)
}

// PlayerCarCoordinates はプレイヤー車両の3D座標を返します。
func (g *Graphics) PlayerCarCoordinates() telemetry.Coordinates {
	for i, id := range g.CarId {
		if id == g.PlayerCarId {
			return g.CarCoordinates[i]
		}
	}
	return telemetry.Coordinates{}
}

// CurrentTimeString は CurrentTime を Go の string に変換します。
func (g *Graphics) CurrentTimeString() string {
	return telemetry.DecodeString(g.CurrentTime[:])
}

// LastTimeString は LastTime を Go の string に変換します。
func (g *Graphics) LastTimeString() string {
	return telemetry.DecodeString(g.LastTime[:])
}

// BestTimeString は BestTime を Go の string に変換します。
func (g *Graphics) BestTimeString() string {
	return telemetry.DecodeString(g.BestTime[:])
}

// SplitString は Split を Go の string に変換します。
func (g *Graphics) SplitString() string {
	return telemetry.DecodeString(g.Split[:])
}

// TyreCompoundString は TyreCompound を Go の string に変換します。
func (g *Graphics) TyreCompoundString() string {
	return telemetry.DecodeString(g.TyreCompound[:])
}

// LastSectorTimeString は LastSectorTime を Go の string に変換します。
func (g *Graphics) LastSectorTimeString() string {
	return telemetry.DecodeString(g.Split[:])
}

// ACStatusValue は Status フィールドを ACStatus 型として返します。
func (g *Graphics) ACStatusValue() ACStatus {
	return ACStatus(g.Status)
}

// ACSessionTypeValue は Session フィールドを ACSessionType 型として返します。
func (g *Graphics) ACSessionTypeValue() ACSessionType {
	return ACSessionType(g.Session)
}

// ACFlagTypeValue は Flag フィールドを ACFlagType 型として返します。
func (g *Graphics) ACFlagTypeValue() ACFlagType {
	return ACFlagType(g.Flag)
}
