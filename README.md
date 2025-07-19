# EDITORJS GO LIST

`editorjs-go-list` is an official plugin for [`editorjs-go`](https://github.com/christiandenisi/editorjs-go) that provides rendering support for Editor.js `list` blocks into safe and semantic HTML output.

## ✨ Features

- ✅ Supports both unordered (`<ul>`) and ordered (`<ol>`) lists
- ✅ Escapes HTML entities for each list item
- ✅ Simple integration with `editorjs-go`
- 🧩 Zero dependencies and clean output

---

## 📦 Installation

Install the plugin along with `editorjs-go`:

```bash
go get github.com/christiandenisi/editorjs-go
go get github.com/christiandenisi/editorjs-go-list
```

---

## 🚀 Usage

### Register the plugin in your converter and convert

```go
package main

import (
    "fmt"
    "github.com/christiandenisi/editorjs-go"
    "github.com/christiandenisi/editorjs-go-plugin-list/list"
)

func main() {
    jsonData := []byte(`{
        "time": 1234567890,
        "version": "2.27.0",
        "blocks": [
            {
                "type": "list",
                "data": {
                    "style": "unordered",
                    "items": ["First", "Second"]
                }
            }
        ]
    }`)

    conv := editorjs.New()
    editorjs.Register(conv, "list", list.RenderList)

    html, err := conv.Convert(jsonData)
    if err != nil {
        panic(err)
    }

    fmt.Println(html)
}
```

---

## 📌 Notes

- Escaping is handled using `html.EscapeString(...)`
- If the list is empty, it returns an empty string (`""`)

---

## 🧱 Compatibility

This plugin is compatible with:

- `editorjs-go` version `1.x`

---

## 👤 License

MIT – free to use, modify, and distribute.