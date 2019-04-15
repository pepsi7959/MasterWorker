package masterworker

import "time"

type Worker struct {
	wokerId int
	ctx     Context
	stop    chan int
}

func (w *Worker) Run() {
	for {
		select {
		case <-w.stop:
			return
		default:
			time.Sleep(50 * time.Millisecond)
			//get a new job
			if w.ctx != nil {
				//do the job
				if w.ctx.res != nil {
					w.ctx.res.Data = "Hello"
					return
				}

			}
		}
	}
}
