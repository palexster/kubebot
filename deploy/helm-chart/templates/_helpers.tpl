{{/*
Expand the name of the chart.
*/}}
{{- define "name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified zerokit-config name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "zerokit-configmap.fullname" -}}
{{- if .Values.zerokit.enabled -}}
{{- printf "%s-zerokit" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- .Values.zerokit.configMap -}}
{{- end -}}
{{- end -}}

{{/*
Create a default fully qualified zerokit-secret name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "zerokit-secret.fullname" -}}
{{- if .Values.zerokit.enabled -}}
{{- printf "%s-zerokit" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- .Values.zerokit.secret -}}
{{- end -}}
{{- end -}}

{{/*
Create a default fully qualified zerokit-secret name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "postgresql.fullname" -}}
{{- if .Values.postgresql.enabled -}}
{{- printf "%s-postgresql" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- .Values.postgresql.externalService -}}
{{- end -}}
{{- end -}}