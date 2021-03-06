package main

import (
	"fmt"
	"net/http"
	"github.com/mls93/crosszeros/controller"
	"github.com/gorilla/mux"


	"github.com/mls93/crosszeros/model"

)


const STATIC_DIR = "./view/build/web"
const PORT = ":8090"
const HTTPS_PORT=":10443"

func main() {

	playerList := *model.GetPlayerList()
	r := mux.NewRouter()
    r.HandleFunc("/newClient",controller.ClientHandler)
    r.HandleFunc("/",controller.MainHandler)
    r.HandleFunc("/logout",controller.LogoutHandler)
	r.HandleFunc("/updatePlayers",controller.UpdatePlayersHandler)
	r.HandleFunc("/confirm", controller.ConfirmHandler)
	r.HandleFunc("/ws", controller.ServeWs)
	r.HandleFunc("/getTable",controller.GetTableHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(STATIC_DIR)))

	fmt.Println(playerList.Len())
	//generate_cert()
	http.Handle("/",r)	
	fmt.Println("Open localhost:8090 to see webpage")
	http.ListenAndServe(PORT,nil)
	//err :=http.ListenAndServeTLS(HTTPS_PORT, "cert.pem", "key.pem", nil)
	//if err != nil {
	//	fmt.Println("ListenAndServe: ", err)
	//}
}







