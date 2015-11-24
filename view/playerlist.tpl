{{define "T"}}

<table class="table table-striped table-hover table-bordered">
    <thead><th>Player nickname</th></thead>
    <tbody>
    {{range .List}}
        <tr><td class='players  {{if ne $.Val .Name | and .IsFree }}players_free' >
        <a href='#'>{{.Name}}</a>
        {{else}}'>{{.Name}}{{end}}</td>
        </tr>
    {{end}}
    </tbody>
</table>


{{end}}