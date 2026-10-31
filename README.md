# ac-telemetry

Assetto Corsa の共有メモリから取得できるテレメトリデータをパースするための Go ライブラリです。

AC / ACC / ACR に対応しており、3種類の構造体データ（StaticInfo, Graphics, Physics）を取得できます。

## プロジェクト概要

このプロジェクトは Assetto Corsa の `acpmf` 共有メモリからデータを取得し、Go の構造体にパースするライブラリです。
以下の3つのデータタイプをサポートしています：

- **StaticInfo** - セッション開始時のみ更新される静的情報（車両情報、トラック情報、設定など）
- **Graphics** - 描画時のみ更新されるグラフィックス情報（ラップタイム、順位、フラグなど）
- **Physics** - 1KHzサイクルで更新される物理演算情報（速度、ギア、タイヤ状態など）

## インストール

```bash
go get github.com/nobonobo/ac-telemetry
```

## 使用方法

```go
import "github.com/nobonobo/ac-telemetry"
```

`Telemetry` 構造体を通じて各データを取得できます：

```go
telemetry := &Telemetry{
    StaticInfo: &StaticInfo{},
    Graphics:   &Graphics{},
    Physics:    &Physics{},
}
```

## 構造体一覧

### StaticInfo

`acpmf_static` 共有メモリ構造体を表します。セッション開始時のみ更新されます。

| フィールド | 型 | 説明 |
|-----------|------|------|
| SMVersion | [15]uint16 | SMバージョン |
| ACVersion | [15]uint16 | ACバージョン |
| NumberOfSessions | int32 | セッション数 |
| NumCars | int32 | 車両数 |
| CarModel | [33]uint16 | *カーモデル名 |
| Track | [33]uint16 | *トラック |
| PlayerName | [33]uint16 | プレイヤー名 |
| PlayerSurname | [33]uint16 | プレイヤー姓 |
| PlayerNick | [33]uint16 | プレイヤーニックネーム |
| SectorCount | int32 | セクター数 |
| MaxTorque | float32 | 最大トルク |
| MaxPower | float32 | 最大パワー |
| MaxRpm | int32 | 最大回転数 |
| MaxFuel | float32 | 最大燃料 |
| SuspensionMaxTravel | [4]float32 | サスペンション最大トラベル |
| TyreRadius | [4]float32 | タイヤ半径 |
| MaxTurboBoost | float32 | 最大ターボブースト |
| PenaltiesEnabled | int32 | ペナルティ有効 |
| AidFuelRate | float32 | 燃料アシスト |
| AidTireRate | float32 | タイヤアシスト |
| AidMechanicalDamage | float32 | メカニカルダメージアシスト |
| AidAllowTyreBlankets | int32 | タイヤブランケット許可 |
| AidStability | float32 | スタビリティアシスト |
| AidAutoClutch | int32 | オートクラッチアシスト |
| AidAutoBlip | int32 | オートブリップアシスト |
| HasDRS | int32 | DRS有効 |
| HasERS | int32 | ERS有効 |
| HasKERS | int32 | KERS有効 |
| KersMaxJ | float32 | KERS最大エネルギー |
| EngineBrakeSettingsCount | int32 | エンジンブレーキ設定数 |
| ErsPowerControllerCount | int32 | ERSパワーコントロール数 |
| TrackSplineLength | float32 | *トラックの標準長 m |
| TrackConfiguration | [33]uint16 | トラックコンフィギュレーション |
| ErsMaxJ | float32 | ERS最大エネルギー |
| IsTimedRace | int32 | タイムレース |
| HasExtraLap | int32 | エキストララップ |
| CarSkin | [33]uint16 | 車両スキン |
| ReversedGridPositions | int32 | リバースグリッドポジション |
| PitWindowStart | int32 | ピットウィンドウ開始 |
| PitWindowEnd | int32 | ピットウィンドウ終了 |
| IsOnline | int32 | オンライン |
| DryTyresName | [33]uint16 | ドライタイヤ名 |
| WetTyresName | [33]uint16 | ウェットタイヤ名 |

### Graphics

`acpmf_graphics` 共有メモリ構造体を表します。描画時のみ更新されます。

| フィールド | 型 | 説明 |
|-----------|------|------|
| PacketId | int32 | *パケット番号 |
| Status | int32 | ACStatus (Off/Replay/Resume/Pause) |
| Session | int32 | ACSessionType (Practice/Qualify/Race/Hotlap等) |
| CurrentTime | [15]uint16 | 現在ラップタイム文字列 |
| LastTime | [15]uint16 | 前ラップタイム文字列 |
| BestTime | [15]uint16 | ベストラップタイム文字列 |
| Split | [15]uint16 | 区間タイム文字列 |
| CompletedLaps | int32 | 完了ラップ数 |
| Position | int32 | 順位 |
| CurrentTimeMs | int32 | 現在ラップタイム (ms) |
| LastTimeMs | int32 | 前ラップタイム (ms) |
| BestTimeMs | int32 | ベストラップタイム (ms) |
| SessionTimeLeft | float32 | セッション残り時間 (秒) |
| DistanceTraveled | float32 | *走行距離 m |
| IsInPit | int32 | ピットイン中 (0/1) |
| CurrentSectorIndex | int32 | 現在のセクターインデックス |
| LastSectorTime | int32 | 前セクタータイム (ms) |
| NumberOfLaps | int32 | セッション総ラップ数 |
| TyreCompound | [33]uint16 | タイヤコンパウンド名 |
| ReplayTimeMultiplier | float32 | リプレイ速度倍率 |
| NormalizedCarPosition | float32 | コース上の正規化位置 (0.0-1.0) |
| ActiveCars | int32 | アクティブ車両数 |
| CarCoordinates | [60]telemetry.Coordinates | *全車両3D座標 [car_id][X,Y,Z] |
| CarId | [60]int32 | 全車両ID |
| PlayerCarId | int32 | *プレイヤー車両ID (常にゼロ) |
| PenaltyTime | float32 | *ペナルティ時間 (秒) |
| Flag | int32 | ACFlagType (No/Blue/Yellow/Black/White/Checkered/Penalty) |
| Penalty | int32 | ペナルティ種別 |
| IdealLineOn | int32 | 理想ライン表示 (0/1) |
| IsInPitLane | int32 | ピットレーン走行中 (0/1) |
| SurfaceGrip | float32 | 路面グリップ係数 |
| MandatoryPitDone | int32 | 義務ピット完了 (0/1) |
| WindSpeed | float32 | 風速 m/s |
| WindDirection | float32 | 風向 (rad) |
| IsSetupMenuVisible | int32 | セットアップメニュー表示中 (0/1) |
| MainDisplayIndex | int32 | メインディスプレイインデックス |
| SecondaryDisplayIndex | int32 | サブディスプレイインデックス |
| TC | int32 | *TCレベル |
| TCCut | int32 | TCカットレベル |
| EngineMap | int32 | エンジンマップ |
| Abs | int32 | *ABSレベル |
| FuelXLap | int32 | 燃料 x ラップ |
| RainLights | int32 | レインライト点灯 (0/1) |
| FlashingLights | int32 | フラッシングライト (0/1) |
| LightsStage | int32 | ライトステージ |
| ExhaustTemperature | float32 | 排気温度 K |
| WiperLv | int32 | ワイパーレベル |
| DriverStintTotalLeft | int32 | ドライバースティント総残り時間 (ms) |
| DriverStintLeft | int32 | ドライバースティント残り時間 (ms) |
| RainTyres | int32 | レインタイヤ装着 (0/1) |

### Physics

`acpmf_physics` 共有メモリ構造体を表します。1KHzサイクルで更新されます。

| フィールド | 型 | 説明 |
|-----------|------|------|
| PacketId | int32 | *パケットID |
| Gas | float32 | *0.0-1.0 アクセル |
| Brake | float32 | *0.0-1.0 ブレーキ |
| Fuel | float32 | *燃料残量 kg |
| Gear | int32 | *ギア (-1=R, 0=N, 1-7) |
| Rpms | int32 | *エンジン回転数 |
| SteerAngle | float32 | *ステアリング角度 |
| SpeedKmh | float32 | *速度 km/h |
| Velocity | telemetry.Coordinates | *ワールド座標速度 (X,Y,Z) m/s |
| AccG | telemetry.Coordinates | *G力 (X,Y,Z) |
| WheelSlip | [4]float32 | *タイヤスリップ [FL,FR,RL,RR] |
| WheelLoad | [4]float32 | *タイヤ荷重 [FL,FR,RL,RR] |
| WheelsPressure | [4]float32 | *タイヤ空気圧 [FL,FR,RL,RR] psi |
| WheelAngularSpeed | [4]float32 | *タイヤ角速度 [FL,FR,RL,RR] rad/s |
| TyreWear | [4]float32 | タイヤ摩耗 [FL,FR,RL,RR] |
| TyreDirtyLevel | [4]float32 | タイヤ汚れ [FL,FR,RL,RR] |
| TyreCoreTemperature | [4]float32 | *タイヤコア温度 [FL,FR,RL,RR] K |
| SuspensionTravel | [4]float32 | *サスペンションストローク [FL,FR,RL,RR] m |
| DRS | float32 | DRS (0=closed, 1=open) |
| TC | float32 | TCトリガー値 |
| Heading | float32 | *ヘディング角 (rad) |
| Pitch | float32 | *ピッチ角 (rad) |
| Roll | float32 | *ロール角 (rad) |
| CgHeight | float32 | 重心高さ m |
| CarDamage | [5]float32 | ボディダメージ [前,後,左,右,センター] |
| NumberOfTyresOut | int32 | コース外タイヤ本数 |
| PitLimiterOn | int32 | ピットリミッター (0/1) |
| Abs | float32 | ABSトリガー値 |
| KersCharge | float32 | KERSチャージ量 |
| KersInput | float32 | KERS入力量 |
| AutoShifterOn | int32 | オートシフター (0/1) |
| RideHeight | [2]float32 | 車高 [前,後] m |
| TurboBoost | float32 | ターボブースト |
| Ballast | float32 | バラスト kg |
| AirDensity | float32 | 空気密度 |
| AirTemp | float32 | 気温 K |
| RoadTemp | float32 | 路面温度 K |
| LocalAngularVel | telemetry.Coordinates | ローカル角速度 (X,Y,Z) rad/s |
| FinalFF | float32 | *最終FFB出力 |
| PerformanceMeter | float32 | パフォーマンスメーター |
| EngineBrake | int32 | エンジンブレーキ設定 |
| ErsRecoveryLevel | int32 | ERS回生レベル |
| ErsPowerLevel | int32 | ERS出力レベル |
| ErsHeatCharging | int32 | ERS熱充電 |
| ErsIsCharging | int32 | ERS充電中 (0/1) |
| KersCurrentKJ | float32 | KERS現在蓄電量 KJ |
| DRSAvailable | int32 | DRS使用可能 (0/1) |
| DrsWithEnabled | int32 | DRS有効 (0/1) |
| BrakeTemp | [4]float32 | *ブレーキ温度 [FL,FR,RL,RR] K |
| Clutch | float32 | *0.0-1.0 クラッチ |
| TyreTempI | [4]float32 | タイヤ温度 内側 [FL,FR,RL,RR] K |
| TyreTempM | [4]float32 | タイヤ温度 中央 [FL,FR,RL,RR] K |
| TyreTempO | [4]float32 | タイヤ温度 外側 [FL,FR,RL,RR] K |
| IsAIControlled | int32 | AI制御中 (0/1) |
| TyreContactPoint | [4]Coordinates | *タイヤ接触点 [FL,FR,RL,RR] |
| TyreContactNormal | [4]Coordinates | *タイヤ接触法線 [FL,FR,RL,RR] |
| TyreContactHeading | [4]Coordinates | *タイヤ接触向き [FL,FR,RL,RR] |
| BrakeBias | float32 | *ブレーキバランス前後比 |
| LocalVelocity | telemetry.Coordinates | *ローカル座標速度 (X,Y,Z) m/s |
| P2PActivations | int32 | P2P有効化数 |
| P2PStatus | int32 | P2P状態 |
| CurrentMaxRpm | int32 | *現在の最大回転数 |
| Mz | [4]float32 | *タイヤ Mz モーメント [FL,FR,RL,RR] |
| Fx | [4]float32 | *タイヤ Fx 力 [FL,FR,RL,RR] |
| Fy | [4]float32 | *タイヤ Fy 力 [FL,FR,RL,RR] |
| SlipRatio | [4]float32 | *スリップ比 [FL,FR,RL,RR] |
| SlipAngle | [4]float32 | *スリップ角 [FL,FR,RL,RR] |
| TcinAction | int32 | TC動作中 (0/1) |
| AbsinAction | int32 | ABS動作中 (0/1) |
| SuspensionDamage | [4]float32 | サスペンションダメージ [FL,FR,RL,RR] |
| TyreTemp | [4]float32 | *タイヤ温度 第4ブロック [FL,FR,RL,RR] K |

## フィールドについて

`*` コメントが付いているフィールドは、ACR（Assetto Corsa Rally）で有効なデータが得られるものです。
AC（Assetto Corsa）および ACC（Assetto Corsa Competizione）では未確認です。

## メソッド

各構造体には、文字列変換用のメソッドが定義されています：

- `SMVersionString()` - SMバージョンを文字列に変換
- `ACVersionString()` - ACバージョンを文字列に変換
- `CarModelString()` - カーモデル名を文字列に変換
- `TrackString()` - トラック名を文字列に変換
- `PlayerNameString()` - プレイヤー名を文字列に変換
- `CurrentTimeString()` - 現在ラップタイムを文字列に変換
- `LastTimeString()` - 前ラップタイムを文字列に変換
- `BestTimeString()` - ベストラップタイムを文字列に変換
- `TyreCompoundString()` - タイヤコンパウンド名を文字列に変換

型安全なアクセス用メソッド：

- `ACStatusValue()` - Status を ACStatus 型として返す
- `ACSessionTypeValue()` - Session を ACSessionType 型として返す
- `ACFlagTypeValue()` - Flag を ACFlagType 型として返す

## ライセンス

MIT License