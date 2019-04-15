package masterworker

type Context struct {
}

type Task struct {
	taskId int
	ctx    Context
}

type Job struct {
	jobId int
	tasks []Task
}

type Master struct {
	jobs    []Job
	workers []Workers
}

func (m *Master) Submit(jobs Job) {
	m.jobs = jobs
	if m.worker != nil {
		for k,v range m.worker {
			m.woker[k].
	}
}
