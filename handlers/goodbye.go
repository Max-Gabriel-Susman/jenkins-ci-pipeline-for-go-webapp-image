package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger} // I don't think I quite understand this syntax
}

func (goodBye *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	goodBye.logger.Println("Goodbye World")
	rw.Write([]byte("Goodbye cruel World"))
	/*

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Goodbye %s", d)
	*/
}
