package masterworker

type Request struct {
	HttpVersion string
	Headers     map[string]string
	Body        string
}

type Result struct {
	Status int
	Data   string
}

type Context struct {
	req Request
	res Result
}
