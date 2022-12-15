package jobscheduler

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	interval := time.Millisecond
	result := 0
	args := []struct {
		jobs         []Job
		jobRunAmount int
	}{
		{
			jobRunAmount: 9,
			jobs: []Job{
				NewJob(3, interval, func() { result++ }, true),
				NewJob(2, interval, func() { result++ }, false),
				NewJob(4, interval, func() { result++ }, true),
			},
		},
		{
			jobRunAmount: 6,
			jobs: []Job{
				NewJob(1, interval, func() { result++ }, false),
				NewJob(2, interval, func() { result++ }, true),
				NewJob(3, interval, func() { result++ }, false),
			},
		},
	}

	for index, targ := range args {
		// set an execution limit
		go Run(targ.jobs)
		<-time.After(time.Duration(targ.jobRunAmount) * interval)

		if result != targ.jobRunAmount {
			t.Error("test", index+1, "failed: expected", targ.jobRunAmount, "in `result`, got", result)
		}

		result = 0
	}
}
