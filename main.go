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
    line = strings.TrimSpace(line)
    
    // Headers
    if strings.HasPrefix(line, "# ") {
        return fmt.Sprintf("<h1>%s</h1>", strings.TrimPrefix(line, "# "))
    }
    if strings.HasPrefix(line, "## ") {
        return fmt.Sprintf("<h2>%s</h2>", strings.TrimPrefix(line, "## "))
    }
    if strings.HasPrefix(line, "### ") {
        return fmt.Sprintf("<h3>%s</h3>", strings.TrimPrefix(line, "### "))
    }
    
    // Bold
    if strings.Contains(line, "**") {
        line = strings.ReplaceAll(line, "**", "<strong>")
        line = strings.ReplaceAll(line, "</strong>", "</strong>")
    }
    
    // Empty line
    if line == "" {
        return "<br>"
    }
    
    // Regular paragraph
    return fmt.Sprintf("<p>%s</p>", line)
}