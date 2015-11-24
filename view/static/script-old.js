var socket;
if($.cookie("clientName")) socket = openSocket();
var isCross
var isTurn
var cross = "X"
var zero = "O"
var madestepName =":MadeStep"
var upd_interval = 0
var act_rows =".active-rows"
var onclick_str = "onclick"

var your_turn = "Your turn"
var opposite_turn = "Your opposite`s turn"

var turn_div ="#which_turn"

function quit(){
/*var enemyName =$.cookie("enemyName")
		console.log(enemyName)
		if (enemyName){
			$("#play").removeAttr("style")
			$("")
		}
*/
	var enemyName =$.cookie("enemyName")
	$.removeCookie('enemyName', { path: '/' });
	socket.send(enemyName+":QuitGame")
	$("#field").html("")
	$("#play").removeAttr("style")
}

$(document).ready(function() {
	function getClientName(){var clName =$.cookie("clientName")
		console.log(clName)
		if (clName){
			console.log(clName+" Helll")
			$("#registr").css("display","none")
			$("#play").removeAttr("style")
			$("#login_name").html($.cookie("clientName"))
			$("#logout").removeAttr("style")
			upd_interval = setInterval("updPlayers()",2000);
		}
		else{
			$("#play").css("display","none")
			$("#registr").removeAttr("style")
			$("#login_name").html("")
			$("#logout").css("display","none")
		}
	}

	function getEnemyName(){
		var enemyName =$.cookie("enemyName")
		console.log(enemyName)
		if (enemyName){

			$("#play").css("display","none")
			$("#field").removeAttr("style")
			updateTable()
			

		}

	}

	$("#logout").click(function(event) {
		event.preventDefault()
		el = $(this)	
		url = el.find("a").attr('href')
		console.log(url)		
		var posting = $.get( url);
		posting.done(function( data ) {
			getClientName()
			console.log($.cookie("clientName"))
			clearInterval(upd_interval)
			socket.close()
		});
	})


	$("#newClientForm").submit(function(event) {
		event.preventDefault();
	
		var $form = $( this ),
		url = $form.attr( 'action' );
		var name = $form.find("input#clientNameField").val()
		
		var posting = $.post( url, {name:name} );
		posting.done(function( data ) {
			getClientName()
			console.log($.cookie("clientName"))
			socket = openSocket()
		});
	})
	getClientName()
	getEnemyName()

})


function updateTable(){

  $.ajax({
   type: "GET",
   url: "/getTable",
   success: function(msg){
     $("#field").removeAttr("style").html(msg)
   }
 });
}

function updPlayers() { 
  $.ajax({
   type: "GET",
   url: "/updatePlayers",
   success: function(msg){
     $("#active_players").html(msg)
   }
 });
}



function openSocket(){
	var socket = new WebSocket("ws://localhost:8090/ws")


	socket.onopen = function() {
	  var enemyName =$.cookie("enemyName")
	  if (enemyName){
		  socket.send(":CheckVal")
	  }
	  console.log("Соединение установлено.");
	};

	socket.onclose = function(event) {
	  if (event.wasClean) {
		console.log('Соединение закрыто чисто');
	  } else {
		console.log('Обрыв соединения'); // например, "убит" процесс сервера
	  }
	  console.log('Код: ' + event.code + ' причина: ' + event.reason);
	};

	socket.onmessage = function(event) {
	  //console.log("Получены данные " + event.data);
	  //$("#active_players").append("<li class='players' id='player"+event.data+"'>"+event.data+"</li>")
		var msg = event.data
		
		console.log(event.data)
		var generateTable =function (enemy){
			var table = "<div>Your enemy:"+enemy+"</div><div id=\"which_turn\">"
			console.log(isTurn)
			var which_turn = (isTurn)? "Your turn":"Your opposite`s turn"
			table += which_turn
			table += "</div><div onclick=\"quit()\"><a href=\"#\">Quit</a></div><div id=\"won_cond\"></div><table><tbody>";
			for (i = 0;i < 10;i++){
				table+="<tr>";
				var madestep = function(){
					if (isCross){
						$(this).html("X");
					}
					else{
						$(this).html("O");
					}
					socket.send(i+"x"+j+":MadeStep");

				}
				for (j = 0; j<10;j++){
					table+="<td class=\"active-rows table_rows\" "
					if (isTurn){
						table+="onclick=\" if (isCross) $(this).html(cross);else $(this).html(zero);$(this).removeClass(act_rows);$(act_rows).removeAttr(onclick_str);isTurn = !isTurn;$(turn_div).html((isTurn)? your_turn:opposite_turn);socket.send($(this)[0].id+madestepName); \""
					}
					table+=" id=\""+i+"x"+j+"\" ></td>"
				}
				table+="</tr>";
			}
			table+="</tbody></table>";
			return table;

		}
		if (msg.split(":")[0]=="ConfirmClient"){
			var enemy = msg.split(":")[1]

			var onclick_function = "var generateTable="+generateTable+";var posting =$.post(\"/confirm\", {enemy:\""+ enemy+"\"});"
			onclick_function+= "posting.done(function(data){socket.send(\""+enemy+":Accepted\"); $(\"#field\").removeAttr(\"style\");"
			onclick_function+= "$.cookie(\"enemyName\",\""+ enemy+"\"); $(\"#play_dialog\").html(\"\");socket.send(\":CheckVal\")})"
			$("#play_dialog").html("<div class='play_message'><div style='margin:10px;text-align:center'>user <b>"+enemy+"</b> wants to play against you</div><div class='confirm_buttons'><input type='button' class='btn btn-primary' value='submit' onclick='"+onclick_function+"'><input type='button' style='margin-left:5px' class='btn' value='cancel' onclick='$(\"#play_dialog\").html(\"\")'></div></div>")

		}
		if (msg.split(":")[0]=="Confirmed"){
			var enemy = msg.split(":")[1]
			$("#field").removeAttr("style");
			socket.send(":CheckVal")
			$.cookie('enemyName', enemy)

			//act_rows = $(".active-rows")
		}
		if (msg.split(":")[0]=="IsCross"){
			var side_part = msg.split(";")[0]
			var cross = side_part.split(":")[1]
			isCross = (cross === 'true')

			var turn_part = msg.split(";")[1]
			var turn = turn_part.split(":")[1]
			console.log("turn="+turn)
			isTurn = (turn === 'true');
			$("#field").html(generateTable($.cookie('enemyName')))
			$("#which_turn").html((isTurn)? "Your turn":"Your opposite`s turn")

		}
		if (msg.split(":")[0]=="MadeStep"){

			var id_val = msg.split(":")[1]
			var value = (isCross)?"O":"X"
			console.log(value)
			$("#"+id_val).html(value)

			$('.active-rows').attr("onclick","if (isCross) $(this).html(cross);else $(this).html(zero);$(this).removeClass(act_rows);$(act_rows).removeAttr(onclick_str);isTurn=!isTurn;$(turn_div).html((isTurn)? your_turn:opposite_turn);socket.send($(this)[0].id+madestepName)");
			$("#"+id_val).removeAttr("onclick")
			$("#"+id_val).removeClass("active-rows")
			isTurn = !isTurn
			$("#which_turn").html((isTurn)? your_turn:opposite_turn)
		}

		if(msg.split(":")[0]=="YouWon"){
			displayResult(msg)
			$("#won_cond").html("You won!")
			$(".active-rows").removeAttr("onclick")

		}

		if(msg.split(":")[0]=="YouLose"){
			displayResult(msg)
			$("#won_cond").html("You lose")
			$(".active-rows").removeAttr("onclick")
		}

		if(msg.split(":")[0]=="EnemyQuit"){
			$("#won_cond").html("Your opponent left this game")
			$(".active-rows").removeAttr("onclick")
			$.removeCookie('enemyName', { path: '/' });
		}
	};
	return socket
}


function displayResult(msg){
		var field_vals_str = msg.split(":")[1]
		var field_vals = field_vals_str.split(";")
		for (k=0;k<5;k++){
			var i_j = field_vals[k].split(",")
			var i=i_j[0]
			var j=i_j[1]
			$("#"+i+"x"+j).css("background-color","red")
		}


}
