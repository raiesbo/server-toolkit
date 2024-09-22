# Server Toolkit

Collections of tools to spinnup a Go server using exclusively the standard library.

## Features

1. Go template renderer
2. Go template cache // Coming soon
3. Logger middleware // Coming soon

```go
tk := toolkit.Toolkit{
	TmplsDir:    "./ui/html/pages/", // Tmpl pages directory
	BaseTmplDir: "./ui/html/partials/layout.tmpl", // Tmpl base layout directory
}

// Call inside a HTTP handler
tk.RenderTmpl(w, r, "home.tmpl", http.StatusOK, data)
```

## Instalation

```shell
go get github.com/raiesbo/server-toolkit@latest
```
