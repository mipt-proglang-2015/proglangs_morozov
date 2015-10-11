{{define "T"}}
<script type="text/javascript">
    $(document).ready(function(){
     $(".players_free").each(function(i){
        $(this).click(function(){
            var clName  =  $.cookie("clientName");
            var content = $(this).find("a").html();
            console.log(content)
            socket.send(content+"%"+clName+":clientToApprove")

        })


     })



    })

</script>
{{range .List}}
<li class='players  {{if ne $.Val .Name | and .IsFree }}players_free' >
    <a href='#'>{{.Name}}</a>
    {{else}}'>{{.Name}}{{end}}</li>
{{end}}
{{end}}