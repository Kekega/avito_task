package main

import (
	"avito_task/handlers"
	"flag"
	"log"
	"net/http"
	"os"
)

var Version = "1.0.0"
var flagConfig = flag.String("config", "./config/dev.yml", "path to the config file")

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
