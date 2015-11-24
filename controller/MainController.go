package controller
import (
	"net/http"
	"fmt"
	"strconv"
	"html/template"
	"github.com/mls93/crosszeros/model"
)

const MAIN_TEMPLATE = "view/main.html"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://dartpad.dartlang.org")
	playerList := model.GetPlayerList()
	fmt.Println("Hello!"+strconv.Itoa(playerList.Len()))

	if (playerList.Len()==0){
		fmt.Println("Hello!"+strconv.Itoa(playerList.Len()))
		DeleteCookie(w,r,CLIENT_COOKIE)
		DeleteCookie(w,r,OPPOSITE_COOKIE)
	}
	fmt.Println(GetCookie(r,CLIENT_COOKIE))
	t,_ := template.ParseFiles(MAIN_TEMPLATE)

	if (playerList.Len()>0){
		resultArr := model.ListToArr(playerList)
		t.Execute(w,resultArr)
	}else{
		t.Execute(w,0)

	}
}
