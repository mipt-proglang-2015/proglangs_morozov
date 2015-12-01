package model
import (
	"container/list"
	"strings"
	"strconv"
)


var playerList = *list.New()
const FIELD_SIZE = 10
var WIDTH = FIELD_SIZE
var HEIGHT = FIELD_SIZE
const START =5000
const MAX_FIELD_SIZE=10000

func GetPlayerList() *list.List{
	return &playerList
}



type Field struct {
	Field [][]int
	Buffer [][] int
	StartW int
	StartH int
	Width int
	Height int
}

func CreateNewField() *Field{
	buffer := make([][]int,MAX_FIELD_SIZE)
	for i := range buffer {
		buffer[i] = make([]int,MAX_FIELD_SIZE)
	}
	field := buffer[START:START+FIELD_SIZE]
	for i := range field {
		field[i] = buffer[START+i][START:START+FIELD_SIZE];
	}
	return &Field{field,buffer,START,START,FIELD_SIZE,FIELD_SIZE}
}

func ResizeField(fld *Field,changeW int, changeH int) {
	if (changeH==-1){

		fld.StartH--
		fld.Height++
	}
	if (changeH==+1){
		fld.Height++
	}
	if (changeW==-1){
		fld.StartW--
		fld.Width++
	}
	if (changeW==+1){
		fld.Width++
	}

	fld.Field = fld.Buffer[fld.StartH:fld.StartH+fld.Height]
	for i := range fld.Field {


		buffer := make([]int,MAX_FIELD_SIZE)
		for j:=0;j<WIDTH;j++{
			add:=0
			if (changeW==-1) {
				add=1
			}
			buffer[fld.StartW+j+add]=fld.Buffer[i+fld.StartH][j]
		}

		fld.Field[i]=buffer[fld.StartW:fld.StartW+fld.Width]



	}
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