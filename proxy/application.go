package proxy

type Application struct{}

func(a *Application) handleRequest(url, method string)(int, string){
	if url == "app/status" && method == "GET"{
		return 200, "OK"
	}

	if url == "create/user" && method == "POST"{
		return 200, "User Created"
	}

	return 404, "Not Found"
}