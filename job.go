package jobscheduler

import (
	"sync"
	"time"
)

// Job struct includes all properties for a specific job.
type Job struct {
	Function        func()
	Interval        time.Duration
	ExecuteCount    int
	MaxExecuteCount int
	ExecOnInit      bool
	lock            *sync.Mutex
}

func (j *Job) StartJob() (jobDone chan bool) {
	for {
		if j.MaxExecuteCount > 0 && j.ExecuteCount >= j.MaxExecuteCount {
			jobDone <- true
			break
		}

		if !(j.ExecOnInit && j.ExecuteCount == 0) {
			<-time.After(j.Interval)
		}

		j.lock.Lock()
		j.Function()
		j.ExecuteCount++
		j.lock.Unlock()
	}

	return
}

func (j *Job) Run() {
	go j.StartJob()
}
