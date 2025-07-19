package list

import (
    "testing"

    "github.com/christiandenisi/editorjs-go"
)

func TestRenderList_Unordered(t *testing.T) {
    block := editorjs.Block[ListData]{
        Data: ListData{
            Style: "unordered",
            Items: []string{"Item 1", "Item 2"},
        },
    }

    ctx := &editorjs.Context{}
    out, err := RenderList(block, ctx)
    if err != nil {
        t.Fatalf("RenderList returned an error: %v", err)
    }

    expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"
    if out != expected {
        t.Errorf("unexpected output: got %q, want %q", out, expected)
    }
}

func TestRenderList_Ordered(t *testing.T) {
    block := editorjs.Block[ListData]{
        Data: ListData{
            Style: "ordered",
            Items: []string{"First", "Second"},
        },
    }

    ctx := &editorjs.Context{}
    out, err := RenderList(block, ctx)
    if err != nil {
        t.Fatalf("RenderList returned an error: %v", err)
    }

    expected := "<ol><li>First</li><li>Second</li></ol>"
    if out != expected {
        t.Errorf("unexpected output: got %q, want %q", out, expected)
    }
}

func TestRenderList_Escaping(t *testing.T) {
    block := editorjs.Block[ListData]{
        Data: ListData{
            Style: "unordered",
            Items: []string{"<b>bold</b>", "Hello & Goodbye"},
        },
    }

    ctx := &editorjs.Context{}
    out, err := RenderList(block, ctx)
    if err != nil {
        t.Fatalf("RenderList returned an error: %v", err)
    }

    expected := "<ul><li>&lt;b&gt;bold&lt;/b&gt;</li><li>Hello &amp; Goodbye</li></ul>"
    if out != expected {
        t.Errorf("unexpected output: got %q, want %q", out, expected)
    }
}

func TestRenderList_Empty(t *testing.T) {
    block := editorjs.Block[ListData]{
        Data: ListData{
            Style: "unordered",
            Items: []string{},
        },
    }

    ctx := &editorjs.Context{}
    out, err := RenderList(block, ctx)
    if err != nil {
        t.Fatalf("RenderList returned an error: %v", err)
    }

    if out != "" {
        t.Errorf("expected empty output, got %q", out)
    }
}