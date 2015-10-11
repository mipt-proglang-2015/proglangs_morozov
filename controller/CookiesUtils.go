package controller
import (
	"net/http"
	"time"
)

const CLIENT_COOKIE = "clientName"
const OPPOSITE_COOKIE = "enemyName"

func GetCookie(r *http.Request,cookieName string) string{
	clientNameCookie, _ := r.Cookie(cookieName)
	client :=""
	if (clientNameCookie!=nil) {
		client = clientNameCookie.Value
	}
	return client
}

func SetCookie(w http.ResponseWriter,cookieName string, value string){
	expiration := time.Now().Add(time.Hour)
	clientSession := &http.Cookie{Name:cookieName,Value:value,Expires:expiration}
	http.SetCookie(w,clientSession)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request,cookieName string) string{
	cookie, _ := r.Cookie(cookieName)
	val:=""
	if (cookie!=nil){
		val = cookie.Value
		cookie.MaxAge = -1000
		http.SetCookie(w,cookie)
	}
	return val
}