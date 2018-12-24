package query

const UpdateTemplate = `
{{- $table := .Table }}
{{- $fields := .Fields }}
{{- $where := .Where }}
{{- $ifnotexist := .IfNotExist }}
UPDATE {{ $table }} SET
  {{- range $index, $field := $fields }}
  {{- if $field.IsCounter }}
  {{- if $field.IsIncrement }}
  {{ $field.Key }} = {{ $field.Key }} + ?
  {{- end }}
  {{- if not $field.IsIncrement }}
  {{ $field.Key }} = {{ $field.Key }} - ?
  {{- end }}
  {{- end }}
  {{- if not $field.IsCounter }}
  {{ $field.Key }} = ?
  {{- end }}
  {{- if lt $index (minus (len $fields) 1) }},{{- end }}
  {{- end }}
{{- if gt (len $where) 0 }}
WHERE {{ range $index, $field := $where }}
  {{- if $field.IsEq }}
  {{- $field.Key }} = {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsLt }}
  {{- $field.Key }} < {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsLte }}
  {{- $field.Key }} <= {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsGt }}
  {{- $field.Key }} > {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsGte }}
  {{- $field.Key }} >= {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsIn }}
  {{- $field.Key }} IN {{ $field.Parameter }}
  {{- end }}
  {{- if lt $index (minus (len $where) 1) }} AND {{ end }}
{{- end }}
{{- end }}
{{- if $ifnotexist }} IF NOT EXISTS{{ end }};
`
