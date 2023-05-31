---
title: "{{.Title}}"
author: "{{.AuthorName}}"
description: "{{.Description}}"
date: "{{.Date}}"
{{- if .LastMod}}
lastmod: "{{.LastMod}}"{{- else -}}
{{end}}
{{- if .Tags }}
tags: {{.Tags}}{{- else -}}
{{ end }}

{{- if .Categories }}
categories: {{.Categories}}{{- else -}}
{{ end }}
{{- if .Keyword }}
keywords: {{.Keyword}}{{- else -}}
{{ end }}
{{- if .Next }}
next: {{.Next}}{{- else -}}
{{ end }}
{{- if .Prev }}
prev: {{.Prev}}{{- else -}}
{{ end }}
{{- if eq .Status 1}}
draft: false{{else}}
draft: true
{{ end }}
---
{{.Content}}