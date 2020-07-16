// Copyright © 2015-2020 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package test

import (
	"testing"
	"time"
)

// Cleanup wraps a testing.Test or Benchmark with Program
type Cleanup struct {
	testing.TB
}

// Program runs a cleanup task to its End
func (cleanup Cleanup) Program(options ...interface{}) {
	cleanup.Helper()
	p, err := Begin(cleanup.TB, options...)
	if err == nil {
		err = p.End()
	}
	if err != nil {
		cleanup.Log(err)
	}
}

func (cleanup Cleanup) ProgramRetry(tries int, options ...interface{}) {
	var err error
	var p *Program
	cleanup.Helper()

	for try := 0; try < tries; try++ {
		p, err = Begin(cleanup.TB, options...)
		if err == nil {
			err = p.End()
		}
		if err == nil {
			break
		}
		if *VVV {
			cleanup.Log("Retry ", err)
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		Pause.Prompt("Fail ", err)
		cleanup.Log(err)
	}
}
