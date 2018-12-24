package query

const SelectTemplate = `
{{- $table := .Table }}
{{- $fields := .Fields }}
{{- $where := .Where }}
{{- $limit := .Limit }}
SELECT
  {{- range $index, $field := $fields }}
  {{ $field.Key }}
  {{- if lt $index (minus (len $fields) 1) }},{{- end }}
  {{- end }}
FROM {{ $table }}
{{- if gt (len $where) 0 }}
WHERE
{{- range $index, $field := $where }}
  {{- if $field.IsEq }}
  {{ $field.Key }} = {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsLt }}
  {{ $field.Key }} < {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsLte }}
  {{ $field.Key }} <= {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsGt }}
  {{ $field.Key }} > {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsGte }}
  {{ $field.Key }} >= {{ $field.Parameter }}
  {{- end }}
  {{- if $field.IsIn }}
  {{ $field.Key }} IN {{ $field.Parameter }}
  {{- end }}
  {{- if lt $index (minus (len $where) 1) }} AND {{ end }}
{{- end }}
{{- end }}
{{- if gt $limit 0 }}
LIMIT {{ $limit }}
{{- end }};
`
