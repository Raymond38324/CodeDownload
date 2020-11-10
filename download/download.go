package main

import (
	"io"
	"net/http"
	"os"
)

var headers = [...][2]string{
	{"User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"},
	{"Accept-Encoding", "gzip, deflate"},
	{"Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7"},
	{"Cache-Control", "no-cache"},
	{"Host", "idea.medeming.com"},
	{"Pragma", "no-cache"},
	{"Proxy-Connection", "keep-alive"},
	{"Upgrade-Insecure-Requests", "1"},
}

func DownloadFile(url, path string) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	for _, value := range headers {
		request.Header.Add(value[0], value[1])
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, e := os.Create(path)
	if e != nil {
		return e
	}
	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}
	return nil
}
