package leap

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestLeapMotionAdaptor() *LeapMotionAdaptor {
	return NewLeapMotionAdaptor("bot", "/dev/null")
}

func TestLeapMotionAdaptorConnect(t *testing.T) {
	t.SkipNow()
	a := initTestLeapMotionAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestLeapMotionAdaptorFinalize(t *testing.T) {
	t.SkipNow()
	a := initTestLeapMotionAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
