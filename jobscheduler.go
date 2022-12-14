package jobscheduler

// Run starts each added job. If the same job is already running, the new one is scheduled.
func Run(jobs []Job) (completed []bool) {
	for _, job := range jobs {
		go func(task Job) {
			task.Run()
			completed = append(completed, <-task.doneChannel)
		}(job)
	}
	return
}
