package gohtml

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo obtem o título de uma pagina html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			htmlFiltred := r.FindStringSubmatch(string(html))

			if len(htmlFiltred) > 1 {
				c <- htmlFiltred[1]
			} else {
				c <- "not found"
			}

		}(url)
	}
	return c
}
