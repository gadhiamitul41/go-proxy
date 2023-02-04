package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go from vercel !!!!",
		})
	})

	server.GET("/proxy", func(context *Context) {
		url := context.Query("url")
		ua := context.Query("ua")
		c := context.Query("c")
		res, err := Request(url, ua, c)
		if err != nil {
			log.Println(err)
		}
		context.SetHeader("Content-Type", "application/octet-stream")
		context.Data(200, res)
	})
	server.Handle(w, r)
}

func Request(url string, userAgent string, cookie string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("x-forwarded-for", "49.34.11.50")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return body, err
	}
	//    defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	return resBody, err
}
