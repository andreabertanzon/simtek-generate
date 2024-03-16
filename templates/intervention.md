# {{ .Title }}
## {{ .Subtitle }}

&nbsp;  
&nbsp;
## Lavoratori
| Lavoratore | ore spese |
|---|---|
{{- range $index, $worker := .Workers }}
| {{ $worker.Name }} | {{ $worker.Hours }} |
{{- end }}

&nbsp;  
&nbsp;
## Descrizione
{{- range $descItem := .Description }}
- {{ $descItem }}
{{- end }}

&nbsp;  
&nbsp;
## Materiali
| Materiale | u.m. | qta |
|---|---|---|
{{- range $index, $material := .Materials }}
| {{ $material.Material }} | {{ $material.Umeasure }} | {{ $material.Quantity }}
{{- end }}

&nbsp;  
&nbsp;
### Note
{{ .Notes }}
