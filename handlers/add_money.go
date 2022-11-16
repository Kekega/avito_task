package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AddMoney struct {
	l *log.Logger
}

func NewAddMoney(l *log.Logger) *AddMoney {
	return &AddMoney{l}
}

func (a *AddMoney) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("money")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Add money %s", d)
}
