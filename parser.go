package main

import (
    "fmt"
    "strings"
)

// Extended parsing functions

func parseItalic(line string) string {
    for strings.Contains(line, "*") && !strings.Contains(line, "**") {
        start := strings.Index(line, "*")
        if start == -1 {
            break
        }
        
        end := strings.Index(line[start+1:], "*")
        if end == -1 {
            break
        }
        
        before := line[:start]
        content := line[start+1 : start+1+end]
        after := line[start+1+end+1:]
        
        line = before + "<em>" + content + "</em>" + after
    }
    return line
}

func parseCode(line string) string {
    for strings.Contains(line, "`") {
        start := strings.Index(line, "`")
        if start == -1 {
            break
        }
        
        end := strings.Index(line[start+1:], "`")
        if end == -1 {
            break
        }
        
        before := line[:start]
        content := line[start+1 : start+1+end]
        after := line[start+1+end+1:]
        
        line = before + "<code>" + content + "</code>" + after
    }
    return line
}

func parseLinks(line string) string {
    // Simple link parsing [text](url)
    for strings.Contains(line, "[") && strings.Contains(line, "](") {
        start := strings.Index(line, "[")
        if start == -1 {
            break
        }
        
        textEnd := strings.Index(line[start:], "](")
        if textEnd == -1 {
            break
        }
        
        urlEnd := strings.Index(line[start+textEnd+2:], ")")
        if urlEnd == -1 {
            break
        }
        
        before := line[:start]
        text := line[start+1 : start+textEnd]
        url := line[start+textEnd+2 : start+textEnd+2+urlEnd]
        after := line[start+textEnd+2+urlEnd+1:]
        
        line = before + fmt.Sprintf("<a href=\"%s\">%s</a>", url, text) + after
    }
    return line
}