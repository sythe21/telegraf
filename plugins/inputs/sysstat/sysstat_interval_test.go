// +build !race
// +build linux

package sysstat

import (
	"os/exec"
	"testing"
	"time"

	"github.com/sythe21/telegraf/testutil"
)

// TestInterval verifies that the correct interval is created. It is not
// run with -race option, because in that scenario interval between the two
// Gather calls is greater than wantedInterval.
func TestInterval(t *testing.T) {
	// overwriting exec commands with mock commands
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	var acc testutil.Accumulator

	s.interval = 0
	wantedInterval := 3

	err := s.Gather(&acc)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Duration(wantedInterval) * time.Second)

	err = s.Gather(&acc)
	if err != nil {
		t.Fatal(err)
	}

	if s.interval != wantedInterval {
		t.Errorf("wrong interval: got %d, want %d", s.interval, wantedInterval)
	}
}