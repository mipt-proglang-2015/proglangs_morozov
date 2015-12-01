package controller
import (
	"strings"
	"net/http"
	"strconv"
	"github.com/mls93/crosszeros/model"
	"fmt"

)

func askInvitation(names string){
	n := strings.Index(names, "%")
	enemyName := names[:n]
	clientName := names[n+1:]

	sendMessage(enemyName,"ConfirmClient:"+clientName)
}



func acceptInvitation(names string,colons int,client string,w http.ResponseWriter){
	enemyName := names[:colons]
	//fmt.Println("enemy="+enemyName)


	model.SetPlayerSide(enemyName,CROSS)
	model.SetPlayerSide(client,ZEROS)

	model.SetPlayerTurn(enemyName,TURN_NOW)
	model.SetPlayerTurn(client,TURN_AFTER)

	model.SetPlayerStatus(client,OCCUPIED)
	model.SetPlayerStatus(enemyName,OCCUPIED)


	field := model.CreateNewField()

	model.SetPlayerField(client,field)
	model.SetPlayerField(enemyName,field)
	model.SetPlayerOpposite(client,enemyName)
	model.SetPlayerOpposite(enemyName,client)


	sendMessage(enemyName,"Confirmed:"+client)
}

func getSideValue(client string){

	side := model.GetPlayerSide(client)
	turn := model.GetPlayerTurn(client)


	resultIsCross := strconv.FormatBool(side)
	isHisTurn := strconv.FormatBool(turn)



	sendMessage(client,"IsCross:"+resultIsCross+";IsHisSide:"+isHisTurn)

}


func madeStep(client string,value string){

	i,j := model.GetIndices(value)

	result := model.GetResultInCell(client)

	enemy := model.GetPlayerOpposite(client)

	//fmt.Println(client+" now turn:"+strconv.FormatBool(findPlayer(playerList,client).Value.(Player).isHisTurn))
	model.SetPlayerFieldValue(client,i,j,result)
	model.SetPlayerTurn(client,TURN_AFTER)
	model.SetPlayerTurn(enemy,TURN_NOW)

	isVictory,q := model.CalculatePossibleVictory(client,i,j)

	fmt.Println(isVictory)
	fmt.Println(q.ToString())
	//fmt.Println(client+" now turn:"+strconv.FormatBool(findPlayer(playerList,client).Value.(Player).isHisTurn))
	model.ResizePlayerField(client,i,j)

	if (isVictory){
		sendMessage(client,"YouWon:"+q.ToString())
		sendMessage(enemy,"YouLose:"+q.ToString())
	} else {
		sendMessage(client,"StepDone:");
		sendMessage(enemy,"MadeStep:"+value)
	}
}


func quitGame (w http.ResponseWriter, r *http.Request,enemy string){
	player := GetCookie(r,CLIENT_COOKIE)
	DeleteCookie(w,r,OPPOSITE_COOKIE)
	fmt.Println(enemy)
	model.SetPlayerOpposite(player,"")
	model.SetPlayerOpposite(enemy,"")

	model.SetPlayerStatus(player,FREE)
	model.SetPlayerStatus(enemy,FREE)

	model.SetPlayerField(player, model.CreateNewField())
	model.SetPlayerField(enemy, model.CreateNewField())
	sendMessage(enemy,"EnemyQuit:")

}
