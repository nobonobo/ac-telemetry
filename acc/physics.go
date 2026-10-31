package acc

import telemetry "github.com/nobonobo/ac-telemetry"

type PhysicsView struct {
	*telemetry.View
}

func NewPhysics() *PhysicsView {
	return &PhysicsView{View: telemetry.NewView(`Local\acpmf_physics`)}
}

func (v *PhysicsView) Get() *Physics {
	return (*Physics)(v.Ptr())
}

// Physics は Assetto Corsa Competizione の acpmf_physics 共有メモリ構造体です。
// 全タイトルのフィールドを含んだ共通構造体 telemetry.PhysicsCommon と順序を合わせています。
// 必要に応じてタイトル固有フィールド (TurboBoost, Ballast, AirDensity, AirTemp, RoadTemp) を追加しています。
type Physics struct {
	PacketId          int32                 `json:"packet_id"`           // *パケットID
	Gas               float32               `json:"gas"`                 // *0.0–1.0 アクセル
	Brake             float32               `json:"brake"`               // *0.0–1.0 ブレーキ
	Fuel              float32               `json:"fuel"`                // *燃料残量 kg
	Gear              int32                 `json:"gear"`                // *ギア (-1=R,0=N,1–7)
	Rpms              int32                 `json:"rpms"`                // *エンジン回転数
	SteerAngle        float32               `json:"steer_angle"`         // *ステアリング角度
	SpeedKmh          float32               `json:"speed_kmh"`           // *速度 km/h
	Velocity          telemetry.Coordinates `json:"velocity"`            // ワールド座標速度 (X,Y,Z) m/s
	AccG              telemetry.Coordinates `json:"acc_g"`               // G力 (X,Y,Z)
	WheelSlip         [4]float32            `json:"wheel_slip"`          // タイヤスリップ [FL,FR,RL,RR]
	WheelLoad         [4]float32            `json:"wheel_load"`          // タイヤ荷重
	WheelsPressure    [4]float32            `json:"wheels_pressure"`     // タイヤ空気圧 (psi)
	WheelAngularSpeed [4]float32            `json:"wheel_angular_speed"` // タイヤ角速度 (rad/s)
	// ACC 固有フィールド
	TurboBoost float32 `json:"turbo_boost,omitempty"` // ターボブースト
	Ballast    float32 `json:"ballast,omitempty"`     // バラスト kg
	AirDensity float32 `json:"air_density,omitempty"` // 空気密度
	AirTemp    float32 `json:"air_temp,omitempty"`    // 気温 K
	RoadTemp   float32 `json:"road_temp,omitempty"`   // 路面温度 K
	// 以降、共通フィールド（rally と同順）
	Drs                float32                  `json:"drs,omitempty"`                 // DRS (0=closed, 1=open)
	TC                 float32                  `json:"tc,omitempty"`                  // TCトリガー値
	Heading            float32                  `json:"heading"`                       // *ヘディング角 (rad)
	Pitch              float32                  `json:"pitch"`                         // *ピッチ角 (rad)
	Roll               float32                  `json:"roll"`                          // *ロール角 (rad)
	CgHeight           float32                  `json:"cg_height,omitempty"`           // 重心高さ m
	CarDamage          [5]float32               `json:"car_damage,omitempty"`          // ボディダメージ [前,後,左,右,センター]
	NumberOfTyresOut   int32                    `json:"number_of_tyres_out,omitempty"` // コース外タイヤ本数
	PitLimiterOn       int32                    `json:"pit_limiter_on,omitempty"`      // ピットリミッター (0/1)
	Abs                float32                  `json:"abs,omitempty"`                 // ABSトリガー値
	KersCharge         float32                  `json:"kers_charge,omitempty"`         // KERSチャージ量
	KersInput          float32                  `json:"kers_input,omitempty"`          // KERS入力量
	AutoShifterOn      int32                    `json:"auto_shifter_on,omitempty"`     // オートシフター (0/1)
	RideHeight         [2]float32               `json:"ride_height,omitempty"`         // 車高 [前,後] m
	LocalAngularVel    telemetry.Coordinates    `json:"local_angular_vel,omitempty"`   // ローカル角速度 (X,Y,Z) rad/s
	FinalFF            float32                  `json:"final_ff,omitempty"`            // *最終FFB出力
	PerformanceMeter   float32                  `json:"performance_meter,omitempty"`   // パフォーマンスメーター
	EngineBrake        int32                    `json:"engine_brake,omitempty"`        // エンジンブレーキ設定
	ErsRecoveryLevel   int32                    `json:"ers_recovery_level,omitempty"`  // ERS回生レベル
	ErsPowerLevel      int32                    `json:"ers_power_level,omitempty"`     // ERS出力レベル
	ErsHeatCharging    int32                    `json:"ers_heat_charging,omitempty"`   // ERS熱充電
	ErsIsCharging      int32                    `json:"ers_is_charging,omitempty"`     // ERS充電中 (0/1)
	KersCurrentKJ      float32                  `json:"kers_current_kj,omitempty"`     // KERS現在蓄電量 KJ
	DrsAvailable       int32                    `json:"drs_available,omitempty"`       // DRS使用可能 (0/1)
	DrsEnabled         int32                    `json:"drs_enabled,omitempty"`         // DRS有効 (0/1)
	BrakeTemp          [4]float32               `json:"brake_temp,omitempty"`          // *ブレーキ温度 [FL,FR,RL,RR] K
	Clutch             float32                  `json:"clutch,omitempty"`              // *0.0–1.0 クラッチ
	TyreTempI          [4]float32               `json:"tyre_temp_i,omitempty"`         // タイヤ温度 内側 [FL,FR,RL,RR] K
	TyreTempM          [4]float32               `json:"tyre_temp_m,omitempty"`         // タイヤ温度 中央 [FL,FR,RL,RR] K
	TyreTempO          [4]float32               `json:"tyre_temp_o,omitempty"`         // タイヤ温度 外側 [FL,FR,RL,RR] K
	IsAIControlled     int32                    `json:"is_ai_controlled,omitempty"`    // AI制御中 (0/1)
	TyreContactPoint   [4]telemetry.Coordinates `json:"tyre_contact_point"`            // *タイヤ接触点 [FL,FR,RL,RR]
	TyreContactNormal  [4]telemetry.Coordinates `json:"tyre_contact_normal"`           // *接触法線
	TyreContactHeading [4]telemetry.Coordinates `json:"tyre_contact_heading"`          // *接触向き
	BrakeBias          float32                  `json:"brake_bias"`                    // *ブレーキバランス前後比
	LocalVelocity      telemetry.Coordinates    `json:"local_velocity"`                // *ローカル座標速度 (X,Y,Z) m/s
	P2PActivations     int32                    `json:"p2p_activations,omitempty"`     // P2P有効化数
	P2PStatus          int32                    `json:"p2p_status,omitempty"`          // P2P状態
	CurrentMaxRpm      int32                    `json:"current_max_rpm"`               // *現在の最大回転数
	Mz                 [4]float32               `json:"mz"`                            // *タイヤ Mz モーメント [FL,FR,RL,RR]
	Fx                 [4]float32               `json:"fx"`                            // *タイヤ Fx 力 [FL,FR,RL,RR]
	Fy                 [4]float32               `json:"fy"`                            // *タイヤ Fy 力 [FL,FR,RL,RR]
	SlipRatio          [4]float32               `json:"slip_ratio"`                    // *スリップ比 [FL,FR,RL,RR]
	SlipAngle          [4]float32               `json:"slip_angle"`                    // *スリップ角 [FL,FR,RL,RR]
	TcinAction         int32                    `json:"tcin_action,omitempty"`         // TC動作中 (0/1)
	AbsinAction        int32                    `json:"absin_action,omitempty"`        // ABS動作中 (0/1)
	SuspensionDamage   [4]float32               `json:"suspension_damage,omitempty"`   // サスペンションダメージ [FL,FR,RL,RR]
	TyreTemp           [4]float32               `json:"tyre_temp"`                     // *タイヤ温度 第4ブロック [FL,FR,RL,RR] K
}
