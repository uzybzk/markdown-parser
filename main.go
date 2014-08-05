package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: markdown-parser <file.md>")
        return
    }
    
    filename := os.Args[1]
    
    err := parseMarkdown(filename)
    if err != nil {
        fmt.Printf("Error parsing markdown: %v\n", err)
        return
    }
}

func parseMarkdown(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        html := parseMarkdownLine(line)
        fmt.Println(html)
    }
    
    return scanner.Err()
}

func parseMarkdownLine(line string) string {
    original := line
    line = strings.TrimSpace(line)
    
    // Headers
    if strings.HasPrefix(line, "# ") {
        content := strings.TrimPrefix(line, "# ")
        content = parseInlineElements(content)
        return fmt.Sprintf("<h1>%s</h1>", content)
    }
    if strings.HasPrefix(line, "## ") {
        content := strings.TrimPrefix(line, "## ")
        content = parseInlineElements(content)
        return fmt.Sprintf("<h2>%s</h2>", content)
    }
    if strings.HasPrefix(line, "### ") {
        content := strings.TrimPrefix(line, "### ")
        content = parseInlineElements(content)
        return fmt.Sprintf("<h3>%s</h3>", content)
    }
    
    // Empty line
    if line == "" {
        return "<br>"
    }
    
    // Regular paragraph
    content := parseInlineElements(line)
    return fmt.Sprintf("<p>%s</p>", content)
}

func parseInlineElements(line string) string {
    // Parse bold
    for strings.Contains(line, "**") {
        start := strings.Index(line, "**")
        if start == -1 {
            break
        }
        
        end := strings.Index(line[start+2:], "**")
        if end == -1 {
            break
        }
        
        before := line[:start]
        content := line[start+2 : start+2+end]
        after := line[start+2+end+2:]
        
        line = before + "<strong>" + content + "</strong>" + after
    }
    
    // Parse other inline elements
    line = parseItalic(line)
    line = parseCode(line)
    line = parseLinks(line)
    
    return line
}