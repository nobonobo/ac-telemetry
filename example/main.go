package main

import (
	"fmt"
	"time"

	"github.com/nobonobo/ac-telemetry/evo"
)

func main() {
	static := evo.NewStaticInfo()
	defer static.Close()
	graphics := evo.NewGraphics()
	defer graphics.Close()
	physics := evo.NewPhysics()
	defer physics.Close()
	tick := time.NewTicker(1000 * time.Millisecond)
	var prevStatus evo.StaticInfo
	var prevGraphics evo.Graphics
	var prevPhysics evo.Physics
	for range tick.C {
		if s := static.Get(); s != nil {
			if s.TrackString() != prevStatus.TrackString() {
				fmt.Printf("s:%v\n", s)
				prevStatus = *s
			}
		}
		if g := graphics.Get(); g != nil {
			if g.PacketId != prevGraphics.PacketId {
				fmt.Printf("g:%v\n", g)
				prevGraphics = *g
			}
		}
		if p := physics.Get(); p != nil {
			if p.PacketId != prevPhysics.PacketId {
				fmt.Printf("p:%v\n", p)
				prevPhysics = *p
			}
		}
	}
}
