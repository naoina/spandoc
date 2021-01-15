package spandoc

import "text/template"

var TemplateMarkdown = template.Must(template.New("markdown").Parse(`
{{- define "options" -}}
{{- if .AllowCommitTimestamp -}}ALLOW_COMMIT_TIMESTAMP{{- end -}}
{{- end -}}

{{- define "boolOrEmpty" -}}
{{- if . -}}{{ . }}{{- end -}}
{{- end -}}

{{- range .Tables -}}
# {{ .Name }}
{{ range .Comments -}}
{{ . }}
{{ end }}
COLUMN | TYPE | NOT NULL | PRIMARY | OPTIONS | DESCRIPTION
------ | ---- | ---------| ------- | ------- | -----------
{{- range .Columns }}
{{ .Name }} | {{ .Type }} | {{ template "boolOrEmpty" .NotNull }} | {{ template "boolOrEmpty" .PrimaryKey }} | {{ template "options" . }} | {{ range .Comments }}{{ . }}{{ end -}}
{{ end }}

{{ end -}}
`))
