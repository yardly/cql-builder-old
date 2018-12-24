package query

const InsertTemplate = `
{{- $table := .Table }}
{{- $fields := .Fields }}
{{- $where := .Where }}
{{- $ifnotexist := .IfNotExist }}
INSERT INTO {{ $table }}
(
  {{- range $index, $field := $fields }}
  {{ $field.Key }}
  {{- if lt $index (minus (len $fields) 1) }},{{- end }}
  {{- end }}
) VALUES (
  {{- range $index, $field := $fields }}
  ?
  {{- if lt $index (minus (len $fields) 1) }},{{- end }}
  {{- end}}
){{ if $ifnotexist }} IF NOT EXISTS{{ end }};
`
