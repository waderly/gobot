package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/ttyACM0")
	servo := gpio.NewServoDriver(firmataAdaptor, "servo", "3")

	work := func() {
		gobot.Every(1*time.Second, func() {
			i := uint8(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	gbot.Robots = append(gbot.Robots,
		gobot.NewRobot("servoBot", []gobot.Connection{firmataAdaptor}, []gobot.Device{servo}, work))

	gbot.Start()
}
