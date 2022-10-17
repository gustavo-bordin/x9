package request

type Request struct {
	Method string
	Url    string
}

func NewRequest(method, url string) Request {
	return Request{
		Method: method,
		Url:    url,
	}
}

func StartCheck() {

}
