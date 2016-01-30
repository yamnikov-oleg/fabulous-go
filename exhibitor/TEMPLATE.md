# Fabulous Go

This application parses github repos from [Fabulous Go](https://github.com/avelino/awesome-go) and populates the lists with additional data, like stars count and commit activity.

## Contents
{{ range . }}{{ range $i, $t := .Titles }}{{if and $t (lt $i 6)}}{{times $i "  "}}* [{{$t}}]({{header_link $t}})
{{ end }}{{ end }}{{ end }}

{{ range . }}{{ range $i, $t := .Titles }}{{ title $i $t }}{{ end }}
{{ range .Entries }}{{ $d := .TmplData }}  * [{{.Username}}/{{.Reponame}}](https://github.com/{{.Username}}/{{.Reponame}}) :star:{{$d.Stars}} - {{$d.Description}} *(**{{$d.CommitsLastMonth}}** commits last month)*
{{ end }}
{{ end }}
