package jobscheduler

// Run starts each added job. If the same job is already running, the new one is scheduled.
func Run(jobs []Job) {
	for _, job := range jobs {
		go func(task Job) {
			task.Run()
		}(job)
		<-job.doneChannel
	}
}
