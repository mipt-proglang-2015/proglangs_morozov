package controller
import (
	"time"
	"github.com/gorilla/websocket"
	"net/http"
	"bytes"
	"github.com/mls93/crosszeros/model"

)

func sendMessage(playerName string, message string) error{

	ws :=  model.GetPlayerSocket(playerName)
	ws.SetWriteDeadline(time.Now().Add(writeWait))
	return ws.WriteMessage(websocket.TextMessage,[]byte(message))

}

func ping(ws *websocket.Conn) error{
	ws.SetWriteDeadline(time.Now().Add(writeWait))
	return ws.WriteMessage(websocket.PingMessage, []byte{})
}

func pingSockets(ws *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		ws.Close()
	}()
	for {
		select {

		case <-ticker.C:
			if err := ping(ws); err != nil {
				return
			}
		}
	}
}

// readPump pumps messages from the websocket connection to the hub.
func responseAndAnswer(ws *websocket.Conn, client string,r *http.Request,w http.ResponseWriter) {
	defer func() {
		ws.Close()
	}()
	ws.SetReadLimit(maxMessageSize)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	//enemyFromCookie:=""
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		//enemyNameCookie, _ := r.Cookie("enemyName")
		//if (enemyNameCookie!=nil){
		//	enemyFromCookie =enemyNameCookie.Value
		//}
		//fmt.Println(string(message))
		names := string(message[:bytes.Index(message, []byte(":"))])
		colons := bytes.Index(message, []byte(":"))
		identifier := string(message[colons+1:colons+9])
		//fmt.Println("IDDDDDDDDDDD "+identifier)

		switch(identifier){
		case "clientTo":
			askInvitation(names);break;
		case "Accepted":
			acceptInvitation(names,colons,client,w);break
		case "CheckVal":
			getSideValue(client);break
		case "MadeStep":
			madeStep(client,names)
		case "QuitGame":
			quitGame(w,r,names)
		}


	}
}





