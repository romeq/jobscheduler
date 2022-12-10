# Job scheduler for Go

Schedules functions to run at certain intervals. Probably the most useless library to exist.

## Examples

### With multiple jobs

```go
import "github.com/romeq/jobscheduler"

// ...

jobs := []jobscheduler.Job{
    // - execution count limit of 3
    // - executes once every second
    // - doesn't execute on init
    jobscheduler.NewJob(3, time.Second, func() { /* ... */ }, false),

    // - execution count limit of 2
    // - executes once every second
    // - executes on init
    jobscheduler.NewJob(2, time.Second, func() { /* ... */ }, true),

    // - no execution count limit
    // - executes every 200ms
    // - doesn't execute on init
    jobscheduler.NewJob(0, time.Millisecond * 200, func() { /* ... */ }, false),
}

jobscheduler.Run(jobs)
```

### With a single job

```go
import "github.com/romeq/jobscheduler"

// ...

job := jobscheduler.NewSimpleJob(time.Second * 5, func() {
    // do something fun
})

job.Run()
```
