package model
import (
	"container/list"
	"strings"
	"strconv"

)


var playerList = *list.New()
const FIELD_SIZE = 10


func GetPlayerList() *list.List{
	return &playerList
}



type Field struct {
	Field [FIELD_SIZE][FIELD_SIZE]int
}

func CreateNewField() *Field{
	return &Field{[FIELD_SIZE][FIELD_SIZE]int{}}
}

func ListToArr(playerList *list.List) []Player{
	result := make([]Player, playerList.Len())
	i := 0
	for entry := playerList.Front(); entry != nil; entry = entry.Next(){
		if (entry.Value!=nil) {
			result[i] = entry.Value.(Player)
			i++
		}
	}
	return result
}



func FindPlayer(playerList *list.List, Name string) *list.Element{

	for entry := playerList.Front(); entry != nil; entry = entry.Next() {

		if (entry.Value!=nil) {
			curPlayer := entry.Value.(Player)

			if (curPlayer.Name == Name) {
				return entry
			}
		}
	}
	return nil
}


func GetResultInCell(playerName string) int{
	isCross := GetPlayerSide(playerName)
	result:=1
	if (isCross) {
		result = 2;
	}
	return result
}

func GetIndices(value string) (int,int){
	n := strings.Index(value, "x")
	i,_ :=strconv.Atoi(value[1:n])
	j,_ := strconv.Atoi(value[n+1:])
	return i,j

}