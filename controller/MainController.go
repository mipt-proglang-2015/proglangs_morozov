package controller
import (
	"net/http"
	"html/template"
	"github.com/mls93/crosszeros/model"
)

const MAIN_TEMPLATE = "view/build/web/main.html"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	playerList := model.GetPlayerList()
	if (playerList.Len()==0){
		DeleteCookie(w,r,CLIENT_COOKIE)
		DeleteCookie(w,r,OPPOSITE_COOKIE)
	}
	t,_ := template.ParseFiles(MAIN_TEMPLATE)

	if (playerList.Len()>0){
		resultArr := model.ListToArr(playerList)
		t.Execute(w,resultArr)
	}else{
		t.Execute(w,0)

	}
}
