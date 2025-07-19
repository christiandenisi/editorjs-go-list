package list

import (
	"html"
	"strings"

	"github.com/christiandenisi/editorjs-go"
)

type ListData struct {
	Style string   `json:"style"` // "ordered" or "unordered"
	Items []string `json:"items"`
}

// RenderList is the renderer function for the "list" block.
func RenderList(b editorjs.Block[ListData], ctx *editorjs.Context) (string, error) {
	if len(b.Data.Items) == 0 {
		return "", nil
	}

	tag := "ul"
	if b.Data.Style == "ordered" {
		tag = "ol"
	}

	var sb strings.Builder
	sb.WriteString("<" + tag + ">")
	for _, item := range b.Data.Items {
		sb.WriteString("<li>")
		sb.WriteString(html.EscapeString(item))
		sb.WriteString("</li>")
	}
	sb.WriteString("</" + tag + ">")

	return sb.String(), nil
}
