
import 'dart:html';
import 'dart:core';
import 'dart:async';

WebSocket socket;
Map<String,StreamSubscription> listeners = new Map<String,StreamSubscription>(); 
Timer timer;
bool isTurn=false;
bool isCross=true;
String cross='X';
String zero='O';
bool isUpdated = false;
void main() {
  
  querySelector("#newClientForm").onSubmit.listen(register);
  querySelector("#logout").onClick.listen(logout);

  getClientName();
	getEnemyName();
	print(readCookie("clientName"));
  
}


void register(Event e){
  e.preventDefault();
  FormElement form = e.target as FormElement;
  String url = form.attributes['action'];
  print(url);
  String name = (form.querySelector('input#clientNameField') as InputElement).value;
  
  var data = { 'name' : name };
  print(url);
	print(url);
  HttpRequest.postFormData(url, data).then((HttpRequest request) {
    
    createCookie('clientName',name,1);
    print(name);
		getClientName();

		socket = openSocket();
});
}


void createCookie(String name, String value, int days) {
  String expires;
  if (days != null)  {
    
    DateTime date = new DateTime.now().add(new Duration(days:days));
    expires = '; expires=' + date.toString();    
  } else {
    DateTime then = new DateTime.now();
    expires = '; expires=' + then.toString();
  }
  document.cookie = name + '=' + value + expires + '; path=/';
}

String readCookie(String name) {
  String nameEQ = name + '=';
  List<String> ca = document.cookie.split(';');
  for (int i = 0; i < ca.length; i++) {
    String c = ca[i];
    c = c.trim();
    if (c.indexOf(nameEQ) == 0) {
      return c.substring(nameEQ.length);
    }
  }
  return null;  
}

void eraseCookie(String name) {
  createCookie(name, '', null);
}


	void getClientName(){
    String clName =readCookie("clientName");
		if (clName!=null){
			querySelector("#registr").style.display="none";
			querySelector("#play").setAttribute("style","");
			querySelector("#login_name").innerHtml = clName;
			querySelector("#logout").setAttribute("style","");
			timer = new Timer.periodic(const Duration(seconds: 2), 								updatePlayers);
			print(timer.isActive);
      socket=openSocket();
    }
		else{
			querySelector("#play").style.display="none";
			querySelector("#registr").setAttribute("style","");
			querySelector("#login_name").innerHtml="";
			querySelector("#logout").style.display="none";
		}
	}

	void getEnemyName(){
		String enemyName =readCookie("enemyName");
		print(readCookie("gjName"));
		print(enemyName);
		if (enemyName!=null && enemyName!="" ){
			querySelector("#play").style.display="none";
			querySelector("#field").setAttribute("style","");
			print('upd+enemyName');
			updateTable(false,"");
		}
		else{
			querySelector("#field").style.display="none";
			querySelector("#play").setAttribute("style","");
			print('no null');
		}
	}

	void logout(Event event){
		event.preventDefault();
    AnchorElement el = event.target as AnchorElement;	
		print(el.innerHtml);
    String url = el.getAttribute('href');
		
  	HttpRequest.getString(url).then((String resp) {
  		getClientName();
      print(url);
      timer.cancel();
			print(url);
      socket.close();
	});
  }
    
void updatePlayers(Timer timer) { 
  String url = "/updatePlayers";
  void setClickable(){
    querySelectorAll(".players_free").forEach((playerButton){
      void requestToPlay(Event e){
        String clName  =  readCookie("clientName");
        String content = playerButton.querySelector("a").innerHtml;
        socket.send(content+"%"+clName+":clientToApprove");
      }
      playerButton.onClick.listen(requestToPlay);
    });
  };
  HttpRequest.getString(url).then((String resp) {
 		
    	querySelector("#active_players").innerHtml = resp;
    	setClickable();
	});
}



WebSocket openSocket(){
	//WebSocket socket = new WebSocket("wss://93.175.9.66:10443/ws");
	WebSocket socket = new WebSocket("ws://localhost:8090/ws");
	socket.onOpen.listen ((e) {
	  String enemyName =readCookie("enemyName");
	  if (enemyName!=null && enemyName!=""){
		  socket.send(":CheckVal");
	  }
	  print("Соединение установлено.");
	});

	socket.onClose.listen((CloseEvent event) {
	  if (event.wasClean) {
		print('Соединение закрыто чисто');
	  } else {
		print('Обрыв соединения'); // например, "убит" процесс сервера
	  }
	  print('Код: ' + (event.code).toString() + ' причина: ' + event.reason);
	});

	socket.onMessage.listen ((event) {

    
		String msg = event.data;
		print(event.data);
		String generateTable (enemy){
			String table = "<div>Your enemy:"+enemy+"</div><div id=\"which_turn\">";
			print(isTurn);
			String which_turn = (isTurn)? "Your turn":"Your opposite`s turn";
			table += which_turn;
			table += "</div><div id='quit_button'><a href=\"#\">Quit</a></div><div id=\"won_cond\"></div><table><tbody>";
			for (int i = 0;i < 10;i++){
				table+="<tr>";
				
				for (int j = 0; j<10;j++){
					table+="<td class=\"active-rows table_rows\" ";

					table+=" id=\"c"+i.toString()+"x"+j.toString()+"\" ></td>";
				}
				table+="</tr>";
			}
			table+="</tbody></table>";
			
      return table;

		}
		if (msg.split(":")[0]=="ConfirmClient"){
			var enemy = msg.split(":")[1];

			querySelector("#play_dialog").innerHtml="<div class='play_message'><div style='margin:10px;text-align:center'>user <b>"+enemy+"</b> wants to play against you</div><div class='confirm_buttons'><input type='button' class='btn btn-primary confirm_inputs' value='submit'><input type='button' class='cancel_inputs' style='margin-left:5px' class='btn' value='cancel'></div></div>";
			querySelectorAll('.confirm_inputs').onClick.listen((MouseEvent e)			{
        var data = {enemy:enemy};
       	String url = //"https://93.175.9.66:10443"+
		     "/confirm";
        HttpRequest.postFormData(url, data).then((HttpRequest request){				 		socket.send(enemy+":Accepted"); 						               						 querySelector("#field").setAttribute("style","");
					 createCookie("enemyName",enemy,1); 																		 		 querySelector("#play_dialog").innerHtml="";                            socket.send(":CheckVal");
          
        }                                        
        ); 
      }
     	);
			querySelectorAll('.cancel_inputs').onClick.listen((MouseEvent e)
             {	
               querySelector("#play_dialog").innerHtml="";

             }
      );
    
    }
    
    
		if (msg.split(":")[0]=="Confirmed"){
			String enemy = msg.split(":")[1];
			querySelector("#field").setAttribute("style","");
			socket.send(":CheckVal");
			createCookie('enemyName', enemy,1);
		}


		if (msg.split(":")[0]=="IsCross"){
			String side_part = msg.split(";")[0];
			String cross = side_part.split(":")[1];
			isCross = (cross == 'true');

			String turn_part = msg.split(";")[1];
			String turn = turn_part.split(":")[1];
			print("turn="+turn);
			isTurn = (turn == 'true');
				updateTable(0,"");
				/*	querySelector("#field").innerHtml=generateTable(readCookie('enemyName'));
      querySelector("#quit_button").onClick.listen(quit);
      void madestep (Event e){
					HtmlElement el = (e.target as HtmlElement);
          socket.send(el.id+":MadeStep");

				}
			print("very important:="+isTurn.toString());
			if (isTurn){

				querySelectorAll(".active-rows").forEach((cell){
				print(cell);
				ElementStream onclick = cell.onClick;
        listeners.putIfAbsent(cell.id,()=>onclick.listen(madestep));
      }
        );
      }
			querySelector("#which_turn").innerHtml=(isTurn)? "Your turn":"Your opposite`s turn";*/

		}
		if (msg.split(":")[0]=="MadeStep"){

			//var id_val = msg.split(":")[1];
			var value = (isCross)?"O":"X";
			print(value);
			/*querySelector("#"+id_val).innerHtml = value;

			querySelectorAll('.active-rows').forEach((row){
        void drawCellInnerElement(Event e){
          row.innerHtml=(isCross)?cross:zero;
          row.classes.remove("active-rows");
          isTurn=!isTurn;
          querySelector("#which_turn").innerHtml=(isTurn)?"Your Turn": "Opposite's turn";
					listeners.forEach((numb,listener)=>listener.cancel());
					listeners = new Map<String,StreamSubscription>();


					socket.send(row.id+":MadeStep");
        }
      listeners.putIfAbsent(row.id,()=>row.onClick.listen(drawCellInnerElement));
      
      }
      );
      
          querySelector("#"+id_val).classes.remove("active-rows");
         isTurn=!isTurn;
          querySelector("#which_turn").innerHtml=(isTurn)?"Your Turn": "Opposite's turn";

					var listener = listeners[id_val];
					print(listener.cancel());
					print(listeners.remove(id_val));
*/
			isTurn=!isTurn;
			updateTable(0,"");
		}

		if(msg.split(":")[0]=="YouWon"){
			displayResult(msg);
			isTurn=false;

      listeners.forEach((numb,listener)=>listener.cancel());
      listeners = new Map<String,StreamSubscription>();
			updateTable(2,msg);
		}

		if(msg.split(":")[0]=="YouLose"){
			displayResult(msg);
			isTurn=false;


			listeners.forEach((numb,listener)=>listener.cancel());
      listeners = new Map<String,StreamSubscription>();
			updateTable(1,msg);
		}

		if(msg.split(":")[0]=="EnemyQuit"){
			querySelector("#won_cond").innerHtml="Your opponent left this game";
			print('Quitted');
			listeners.forEach((numb,listener)=>listener.cancel());
      listeners = new Map<String,StreamSubscription>();
			eraseCookie('enemyName');

			socket=openSocket();
			print('new socket should be opened');
		}
		if(msg.split(":")[0]=="StepDone"){

			isTurn=!isTurn;
			querySelector("#which_turn").innerHtml=(isTurn)?"Your Turn": "Opposite's turn";
			listeners.forEach((numb,listener)=>listener.cancel());
			listeners = new Map<String,StreamSubscription>();
			updateTable(0,"");
		}
	});
	return socket;
}

void displayResult(msg){
		var field_vals_str = msg.split(":")[1];
		var field_vals = field_vals_str.split(";");
		for (int k=0;k<5;k++){
			var i_j = field_vals[k].split(",");
			var i=i_j[0];
			var j=i_j[1];
			querySelector("#c"+i+"x"+j).style.backgroundColor="red";
		}
  
}


void quit(Event e){
	String enemyName =readCookie("enemyName");
	if (enemyName!=null && enemyName!="") {

		eraseCookie('enemyName');
		print(readCookie('enemyName'));
		socket.send(enemyName + ":QuitGame");

		socket = openSocket();
	}
	querySelector("#field").innerHtml="";
	querySelector("#play").setAttribute("style",'');
}


void updateTable(int is_end,String msg){
	String url = "/getTable";

	HttpRequest.getString(url).then((String resp) {
    	querySelector("#field").setAttribute('style','');
			querySelector("#field").innerHtml = resp;
    	querySelector("#quit_button").onClick.listen(quit);


			print("His turn:"+isTurn.toString());
			if (isTurn){
      querySelectorAll(".active-rows").forEach((cell){

				void drawCellInnerElement(Event e){
					//cell.innerHtml=(isCross)?cross:zero;
					//cell.classes.remove("active-rows");
					//isTurn=!isTurn;
					//querySelector("#which_turn").innerHtml=(isTurn)?"Your Turn": "Opposite's turn";
					listeners.forEach((numb,listener)=>listener.cancel());
					listeners = new Map<String,StreamSubscription>();


					socket.send(cell.id+":MadeStep");
				}
				listeners.putIfAbsent(cell.id,()=>cell.onClick.listen(drawCellInnerElement));
				});
		}
			if(is_end==0) {
				querySelector("#which_turn").innerHtml = (isTurn) ? "Your turn" : "Your opposite`s turn";
			}
			else{
				querySelector("#which_turn").innerHtml = "Game was ended";
				displayResult(msg);
				if (is_end==2){
					querySelector("#won_cond").innerHtml="You won!";
				} else{
					querySelector("#won_cond").innerHtml="You lose";
				}

			}
});
}
  
