package controller
import (
	"net/http"
	"strconv"
	"github.com/mls93/crosszeros/model"
)

const CELL_ONCLICK = `
	if (isCross)
		$(this).html(cross);
	else $(this).html(zero);
	$(this).removeClass(act_rows);
	$(act_rows).removeAttr(onclick_str);
	isTurn = !isTurn;
	$(turn_div).html((isTurn)? your_turn:opposite_turn);
	socket.send($(this)[0].id+madestepName);
`





func GetTableHandler(w http.ResponseWriter, r *http.Request){

	client := GetCookie(r,CLIENT_COOKIE)
	enemy := GetCookie(r,OPPOSITE_COOKIE)

	field := model.GetPlayerField(client)
	isTurn := model.GetPlayerTurn(client)
	result := drawTable(enemy,isTurn,field)
	w.Write([]byte(result))
}


func drawTable(enemy string, isTurn bool,fld_wrap *model.Field ) string {
	field := fld_wrap.Field
	turn_string :="Your turn"
	if (!isTurn){
		turn_string = "Your opposite`s turn"
	}
	result:= "<div>Your enemy:"+enemy+"</div><div id='which_turn'>"+turn_string+"</div><div id='won_cond'></div><div id='quit_button'><a href=\"#\">Quit</a></div><table><tbody>";

	for i:=0;i<model.FIELD_SIZE;i++{
		result+="<tr>"
		for j:=0;j<model.FIELD_SIZE;j++{
			k := strconv.Itoa(i)
			l := strconv.Itoa(j)

			if (field[i][j]<1){
				result += "<td class='table_rows active-rows' id='"+k+"x"+l+"'></td>"
			} else {
				if (field[i][j]==1){
					result += "<td  class='table_rows' id='"+k+"x"+l+"'>O</td>"
				} else{
					result += "<td  class='table_rows' id='"+k+"x"+l+"'>X</td>"
				}
			}

		}
		result+="</tr>"
	}
	result+="</tbody></table>";
	return result
}
