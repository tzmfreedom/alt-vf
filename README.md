# Alt Visualforce

Visualforce page transpiler from Go [html/template](https://golang.org/pkg/html/template/).

## Install

```bash
$ go get -u github.com/tzmfreedom/alt-vf
```

## Usage

create template file with [html/template](https://golang.org/pkg/html/template/)
```html
<html>
<head></head>
<body>
  {{if .xxx }}
  <p>foo bar</p>
  {{end}}
  {{if eq .xxx 1 }}
  <p>foo bar</p>
  {{end}}
</body>
</html>
```

execute command
```bash
$ alt-vf sample.template
```

output converted template file
```html
<apex:page>
<html>
<head></head>
<body>

<apex:outputPanel rendered="{!xxx}">
<p>foo bar</p>
</apex:outputPanel>


<apex:outputPanel rendered="{!xxx == 1}">
<p>foo bar</p>
</apex:outputPanel>

</body>
</html>
</apex:page>
```
