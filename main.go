package main

import (
	"log"
	"net/http"

	"github.com/covrom/bulmactl/builder"
)

func main() {
	var v builder.ViewGroup
	v = builder.NewDocumentView("Заголовок страницы")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h, e := v.Render(nil)
		if e == nil {
			w.Write([]byte(h))
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
