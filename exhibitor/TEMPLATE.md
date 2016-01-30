{{ range . }}{{ range $i, $t := .Titles }}{{ title $i $t }}{{ end }}
{{ range .Entries }}{{ $d := .TmplData }}  * [{{.Username}}/{{.Reponame}}](https://github.com/{{.Username}}/{{.Reponame}}) - :star:{{$d.Stars}}, {{$d.CommitsLastMonth}} commits last month
{{ end }}
{{ end }}
