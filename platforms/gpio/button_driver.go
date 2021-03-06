package gpio

import (
	"github.com/hybridgroup/gobot"
)

type ButtonDriver struct {
	gobot.Driver
	Active bool
}

func NewButtonDriver(a DigitalReader, name string, pin string) *ButtonDriver {
	return &ButtonDriver{
		Driver: gobot.Driver{
			Name: name,
			Pin:  pin,
			Events: map[string]*gobot.Event{
				"push":    gobot.NewEvent(),
				"release": gobot.NewEvent(),
			},
			Adaptor: a.(gobot.AdaptorInterface),
		},
		Active: false,
	}
}

func (b *ButtonDriver) adaptor() DigitalReader {
	return b.Driver.Adaptor.(DigitalReader)
}

func (b *ButtonDriver) Start() bool {
	state := 0
	gobot.Every(b.Interval, func() {
		newValue := b.readState()
		if newValue != state && newValue != -1 {
			state = newValue
			b.update(newValue)
		}
	})
	return true
}
func (b *ButtonDriver) Halt() bool { return true }
func (b *ButtonDriver) Init() bool { return true }

func (b *ButtonDriver) readState() int {
	return b.adaptor().DigitalRead(b.Pin)
}

func (b *ButtonDriver) update(newVal int) {
	if newVal == 1 {
		b.Active = true
		gobot.Publish(b.Events["push"], newVal)
	} else {
		b.Active = false
		gobot.Publish(b.Events["release"], newVal)
	}
}
