package controller
import (
	"net/http"
	"html/template"
	"github.com/mls93/crosszeros/model"

)

const PLAYERS_TEMPLATE = "view/build/web/playerlist.tpl"

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	w.Header().Set("Access-Control-Request-Method", "POST")
	playerList := model.GetPlayerList()
	clientName := r.PostFormValue("name")
	SetCookie(w,CLIENT_COOKIE,clientName)
	client := &model.Player{clientName,nil,true,false,false,nil,""}
	playerList.PushFront(*client)



}

func ConfirmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	enemyName := r.PostFormValue("enemy")

	SetCookie(w,OPPOSITE_COOKIE, enemyName)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	playerList := model.GetPlayerList()
	clientName := DeleteCookie(w,r,CLIENT_COOKIE)

	playerList.Remove(model.FindPlayer(playerList,clientName))


}

func UpdatePlayersHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	playerList := model.GetPlayerList()
	clientName := GetCookie(r,CLIENT_COOKIE)
	type Page1 struct{
		Val string
		List []model.Player
	}

	page := Page1{clientName,model.ListToArr(playerList)}
	template.ParseFiles()
	t, _ := template.ParseFiles(PLAYERS_TEMPLATE)
	//fmt.Println(page.List)

	t.ExecuteTemplate(w, "T",page )

}
