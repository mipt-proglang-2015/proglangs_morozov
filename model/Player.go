package model
import (
	"github.com/gorilla/websocket"

	"fmt"
)


type Player struct {
	Name string
	WS *websocket.Conn
	IsFree bool
	IsCross bool
	IsHisTurn bool
	Field *Field
	EnemyName string
}




func SetPlayerSocket(playerName string, ws *websocket.Conn){
	thisPlayer := FindPlayer(&playerList,playerName)
	player := thisPlayer.Value.(Player)
	player.WS = ws
	thisPlayer.Value = player
}

func GetPlayerSocket(playerName string)  *websocket.Conn{


	thisPlayer := FindPlayer(&playerList, playerName)

	thisPlayerEnt :=  thisPlayer.Value.(Player)
	return thisPlayerEnt.WS

}



func SetPlayerTurn(playerName string, turn bool){
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	thisPlayerEnt.IsHisTurn = turn
	thisPlayer.Value = thisPlayerEnt
}

func GetPlayerTurn(playerName string) bool{
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	return thisPlayerEnt.IsHisTurn

}




func SetPlayerStatus(playerName string, isFree bool){
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	thisPlayerEnt.IsFree = isFree
	thisPlayer.Value = thisPlayerEnt
}



func SetPlayerSide(playerName string, side bool){
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	thisPlayerEnt.IsCross = side
	thisPlayer.Value = thisPlayerEnt
}

func GetPlayerSide(playerName string) bool{
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	return thisPlayerEnt.IsCross

}




func SetPlayerField(playerName string,field *Field){
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	thisPlayerEnt.Field = field
	thisPlayer.Value = thisPlayerEnt
}

func GetPlayerField(playerName string) *Field{
	thisPlayer := FindPlayer(&playerList, playerName)

	thisPlayerEnt :=  thisPlayer.Value.(Player)
	fmt.Println(thisPlayerEnt)
	fmt.Println(&thisPlayerEnt.Field)
	fmt.Println(&thisPlayerEnt.Field)
	return thisPlayerEnt.Field
}



func SetPlayerFieldValue(playerName string,i int,j int,value int){
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	thisPlayerEnt.Field.Field[i][j] = value
	thisPlayer.Value = thisPlayerEnt
}




func GetPlayerOpposite(playerName string) string{
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	return thisPlayerEnt.EnemyName
}

func SetPlayerOpposite(playerName string,opposite string){
	thisPlayer := FindPlayer(&playerList, playerName)
	thisPlayerEnt :=  thisPlayer.Value.(Player)
	thisPlayerEnt.EnemyName = opposite
	thisPlayer.Value = thisPlayerEnt
}