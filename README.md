# Markdown Parser

A simple markdown to HTML converter written in Go.

## Usage

```bash
go run main.go example.md
```

## Supported Syntax

- Headers: `#`, `##`, `###`
- Bold text: `**text**`
- Paragraphs
- Line breaks

## Example

Input (example.md):
```markdown
# Hello World
This is a **bold** text.

## Subheader
Another paragraph.
```

Output:
```html
<h1>Hello World</h1>
<p>This is a <strong>bold</strong> text.</p>
<br>
<h2>Subheader</h2>
<p>Another paragraph.</p>
```