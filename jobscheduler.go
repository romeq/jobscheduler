package jobscheduler

import (
	"sync"
	"time"
)

func NewJob(maxCount int, interval time.Duration, fn func(), runOnInit bool) Job {
	return Job{
		Function:        fn,
		Interval:        interval,
		MaxExecuteCount: maxCount,
		ExecOnInit:      runOnInit,
		lock:            &sync.Mutex{},
	}
}

func NewSimpleJob(interval time.Duration, fn func()) Job {
	return NewJob(0, interval, fn, false)
}

// Run starts each added job. If the same job is already running, the new one is scheduled.
func Run(jobs []Job) {
	inProgress := len(jobs)
	for _, job := range jobs {
		go func(job Job) {
			jobChannel := job.StartJob()
			<-jobChannel
			inProgress -= 1
		}(job)
	}

	for inProgress > 0 {
		continue
	}
}
