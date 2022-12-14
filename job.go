package jobscheduler

import (
	"sync"
	"time"
)

// Job struct includes all properties for a specific job.
type Job struct {
	Function        func()
	Interval        time.Duration
	executeCount    int
	doneChannel     chan bool
	MaxExecuteCount int
	ExecOnInit      bool
	m               *sync.Mutex
}

func NewJob(maxCount int, interval time.Duration, fn func(), runOnInit bool) Job {
	return Job{
		Function:        fn,
		Interval:        interval,
		MaxExecuteCount: maxCount,
		doneChannel:     make(chan bool, 1),
		ExecOnInit:      runOnInit,
		m:               &sync.Mutex{},
	}
}

func (j *Job) Run() {
	for {
		if j.MaxExecuteCount > 0 && j.executeCount >= j.MaxExecuteCount {
			j.doneChannel <- true
			break
		}

		if !(j.ExecOnInit && j.executeCount == 0) {
			<-time.After(j.Interval)
		}

		j.m.Lock()
		j.Function()
		j.executeCount++
		j.m.Unlock()
	}
}
