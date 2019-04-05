package model

type {{.Table}} struct {
    {{ range $key, $value := .Column }}
        {{ $key }} : {{ $value }}
    {{ end }}
}