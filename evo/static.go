package evo

import telemetry "github.com/nobonobo/ac-telemetry"

// StaticInfoView wraps a view onto the ACEvo static shared‑memory segment.
type StaticInfoView struct {
    *telemetry.View
}

// NewStaticInfo opens the ACEvo static memory (Local\\acevo_pmf_static).
func NewStaticInfo() *StaticInfoView {
    return &StaticInfoView{View: telemetry.NewView(`Local\acevo_pmf_static`)}
}

func (v *StaticInfoView) Get() *StaticInfo { return (*StaticInfo)(v.Ptr()) }

// StaticInfo mirrors the layout of acevo‑shared‑memory's SPageFileStatic.
// All fields are ordered exactly as in the original C struct.
type StaticInfo struct {
    SMVersion                [15]uint16 `json:"sm_version"`                     // SM バージョン
    ACVersion                [15]uint16 `json:"ac_version"`                     // AC バージョン
    NumberOfSessions         int32      `json:"number_of_sessions,omitempty"`   // セッション数
    NumCars                  int32      `json:"num_cars,omitempty"`             // 車両数
    CarModel                 [33]uint16 `json:"car_model"`                      // 車種名
    Track                    [33]uint16 `json:"track"`                          // トラック名
    PlayerName               [33]uint16 `json:"player_name"`                    // プレイヤー名
    PlayerSurname            [33]uint16 `json:"player_surname"`                 // 姓
    PlayerNick               [33]uint16 `json:"player_nick"`                    // ニックネーム
    TrackConfiguration       [33]uint16 `json:"track_configuration"`            // コース構成
    StartingGrip             float32    `json:"starting_grip,omitempty"`        // 初期グリップ係数
    Session                  int32      `json:"session,omitempty"`              // セッション種別 (Practice=0, Qualify=1, Race=2, …)
    SessionLength            int32      `json:"session_length,omitempty"`       // セッション時間 (秒)
    MaxTurboBoost            float32    `json:"max_turbo_boost,omitempty"`      // 最大ターボブースト
    MaxRpm                   int32      `json:"max_rpm,omitempty"`              // 最大回転数
    MaxTorque                float32    `json:"max_torque,omitempty"`           // 最大トルク
    MaxPower                 float32    `json:"max_power,omitempty"`            // 最大出力
    MaxFuel                  float32    `json:"max_fuel,omitempty"`             // 燃料タンク容量
    SuspensionMaxTravel      [4]float32 `json:"suspension_max_travel"`          // サスペンション最大トラベル
    TyreRadius               [4]float32 `json:"tyre_radius"`                    // タイヤ半径
    PenaltiesEnabled         int32      `json:"penalties_enabled,omitempty"`    // ペナルティ有効
    AidFuelRate              float32    `json:"aid_fuel_rate,omitempty"`        // 燃料アシスト率
    AidTireRate              float32    `json:"aid_tire_rate,omitempty"`        // タイヤアシスト率
    AidMechanicalDamage      float32    `json:"aid_mechanical_damage,omitempty"`// メカダメージアシスト率
    AidAllowTyreBlankets     int32      `json:"aid_allow_tyre_blankets,omitempty"` // タイヤブランケット許可
    AidStability             float32    `json:"aid_stability,omitempty"`        // スタビリティアシスト
    AidAutoClutch            int32      `json:"aid_auto_clutch,omitempty"`      // 自動クラッチアシスト
    AidAutoBlip              int32      `json:"aid_auto_blip,omitempty"`        // 自動ブリップアシスト
    HasDRS                   int32      `json:"has_drs,omitempty"`              // DRS 有無
    HasERS                   int32      `json:"has_ers,omitempty"`              // ERS 有無
    HasKERS                  int32      `json:"has_kers,omitempty"`             // KERS 有無
    KersMaxJ                 float32    `json:"kers_max_j,omitempty"`           // KERS 最大エネルギー (kJ)
    EngineBrakeSettingsCount int32      `json:"engine_brake_settings_count,omitempty"` // エンジンブレーキ設定数
    ErsPowerControllerCount  int32      `json:"ers_power_controller_count,omitempty"`  // ERS パワーコントローラ数
    TrackSplineLength        float32    `json:"track_spline_length"`            // トラック全長 (m)
    ErsMaxJ                  float32    `json:"ers_max_j,omitempty"`            // ERS 最大エネルギー (kJ)
    IsTimedRace              int32      `json:"is_timed_race,omitempty"`        // タイムレースか
    HasExtraLap              int32      `json:"has_extra_lap,omitempty"`        // エキストララップ有無
    CarSkin                  [33]uint16 `json:"car_skin"`                       // 車両スキン
    ReversedGridPositions    int32      `json:"reversed_grid_positions,omitempty"` // リバースグリッド数
    PitWindowStart           int32      `json:"pit_window_start,omitempty"`      // ピットウィンドウ開始 (秒)
    PitWindowEnd             int32      `json:"pit_window_end,omitempty"`        // ピットウィンドウ終了 (秒)
    IsOnline                 int32      `json:"is_online,omitempty"`            // オンラインか
    DryTyresName             [33]uint16 `json:"dry_tyres_name"`                 // ドライタイヤ名
    WetTyresName             [33]uint16 `json:"wet_tyres_name"`                 // ウェットタイヤ名
}

// Helper methods to decode UTF‑16 strings.
func (s *StaticInfo) SMVersionString() string { return telemetry.DecodeString(s.SMVersion[:]) }
func (s *StaticInfo) ACVersionString() string { return telemetry.DecodeString(s.ACVersion[:]) }
func (s *StaticInfo) CarModelString() string  { return telemetry.DecodeString(s.CarModel[:]) }
func (s *StaticInfo) TrackString() string    { return telemetry.DecodeString(s.Track[:]) }
func (s *StaticInfo) PlayerNameString() string { return telemetry.DecodeString(s.PlayerName[:]) }
func (s *StaticInfo) PlayerSurnameString() string { return telemetry.DecodeString(s.PlayerSurname[:]) }
func (s *StaticInfo) PlayerNickString() string { return telemetry.DecodeString(s.PlayerNick[:]) }
func (s *StaticInfo) TrackConfigurationString() string { return telemetry.DecodeString(s.TrackConfiguration[:]) }
func (s *StaticInfo) CarSkinString() string { return telemetry.DecodeString(s.CarSkin[:]) }
func (s *StaticInfo) DryTyresNameString() string { return telemetry.DecodeString(s.DryTyresName[:]) }
func (s *StaticInfo) WetTyresNameString() string { return telemetry.DecodeString(s.WetTyresName[:]) }
