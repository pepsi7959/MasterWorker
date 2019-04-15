package main

type Request struct {
	HttpVersion string
	Method      string
	Headers     map[string]string
	Body        string
}

type Result struct {
	Status     int
	ResultData string
}

func main() {
	master := Master{}
	master.jobs = make(Job{}, 2)
	master.workers = make(Worker{}, 1)

	master.jobs[0].jobId = 0
	master.jobs[0].tasks = make(Task{}, 1)
	master.jobs[0].tasks.taskId = 0
	master.jobs[0].tasks.ctx = Context{}

	//set worker
	master.wokers[0].workerId = 0
	master.wokers[0].Ctx = Context{req: Request{
		Body: "Hi"}, res: Result{}}
	master.wokers[0].Run()

	wait := make(chan int, 1)
	<-wait
}
