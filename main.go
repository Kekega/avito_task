package main

import (
	"avito_task/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ah := handlers.NewAddMoney(l)

	sm := http.NewServeMux()
	sm.Handle("/", ah)



	//http.HandleFunc("/reserve_money/{user_id}/{service_id}/{order_id}/{sum}", func(rw http.ResponseWriter, r*http.Request) {
	//	fmt.Print("2")
	//})

	//http.HandleFunc()

	//http.HandleFunc()

	http.ListenAndServe("127.0.0.1:9090", sm)
}
