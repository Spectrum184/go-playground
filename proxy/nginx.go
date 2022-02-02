package proxy

type Nginx struct{
	application *Application
	maxAllowRequest int
	rateLimiter map[string]int
}

func NewNginxServer() *Nginx  {
	return &Nginx{
		application: &Application{},
		maxAllowRequest: 2,
		rateLimiter: make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string)(int, string)  {
	allowed := n.checkRateLimiting(url)

	if !allowed{
		return 403 , "Not Allowed"
	}

	return n.application.handleRequest(url, method)
}

func(n *Nginx)checkRateLimiting(url string) bool{
if n.rateLimiter[url] == 0{
	n.rateLimiter[url] = 1
}

if n.rateLimiter[url] > n.maxAllowRequest{
	return false
}

n.rateLimiter[url] ++
	return true
}