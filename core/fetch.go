package core

import "net/http"

func fetch(url string, interceptors ...func(reqPoint *http.Request)) (*http.Response, error) {
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for _, f := range interceptors {
		f(req)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	res, err := client.Do(req)
	return res, nil
}

func Get(url string) {
	_, _ = fetch(url, func(reqPoint *http.Request) {
		println("dsddsd")
	})
}
