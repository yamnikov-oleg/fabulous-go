# Fabulous Go

This application parses github repos from [Awesome Go](https://github.com/avelino/awesome-go) and populates the lists with additional data, like stars count and commit activity.

To learn how it works and how to refresh the list, see [`BUILDING.md`](https://github.com/yamnikov-oleg/fabulous-go/blob/master/BUILDING.md).

**clm** stands for 'commits last month'.

## Contents
{{ range . }}{{ range $i, $t := .Titles }}{{if and $t (lt $i 6)}}{{times $i "  "}}* [{{$t}}]({{header_link $t}})
{{ end }}{{ end }}{{ end }}

{{ range . }}{{ range $i, $t := .Titles }}{{ title $i $t }}{{ end }}
{{ range .Entries }}{{ $d := .TmplData }}  * [{{.Username}}/**{{.Reponame}}**](https://github.com/{{.Username}}/{{.Reponame}}) - {{$d.Description}} *(:star:{{$d.Stars}} | {{$d.CommitsLastMonth}}clm)*
{{ end }}
{{ end }}
