package idg

import (
	"testing"

	"github.com/rs/xid"
)

func TestID(t *testing.T) {
	guid := xid.New()
	println(guid.String())
	println(guid.Value)
	println(guid.Pid())
	println(guid.Time().String())
	println(guid.Counter())
}
