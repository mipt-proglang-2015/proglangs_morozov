package controller
import (
	"net/http"
	"time"
	"log"

	"github.com/mls93/crosszeros/model"
	"github.com/gorilla/websocket"
)

const (
// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

// Maximum message size allowed from peer.
	maxMessageSize = 512
)

const (
	CROSS = true
	ZEROS = false
	FREE = true
	TURN_NOW =true
	OCCUPIED = false
	TURN_AFTER = false

)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWs(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client:= GetCookie(r,CLIENT_COOKIE)

	model.SetPlayerSocket(client,ws)

	go pingSockets(ws)
	responseAndAnswer(ws,client,r,w)

}




