package leaderelection

import (
	"flag"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	m.Run()
}
