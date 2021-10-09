{{- define "fullname" }}
{{ printf "a: %s" .Chart.Name|indent 2 }}
{{- end }}

{{/*Chart Namespace configuration*/}}
{{- define "chart.Namespace" }}
{{- if eq .Release.Namespace "default" -}}
  {{ .Values.namespace }}
  {{- else -}}
  {{ .Release.Namespace }}
  {{- end }}
{{- end }}
