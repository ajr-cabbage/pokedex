package main

import (
    "strings"
)

func cleanInput(text string) []string {
    cleanedSlice := strings.Fields(strings.ToLower(text))
    
    return cleanedSlice
}