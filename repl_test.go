package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "gIve mE a break     ",
            expected: []string{"give", "me", "a", "break"},
        },
        {
            input:    "   extra       spaces      ",
            expected: []string{"extra", "spaces"},
        },
    }
    
    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("slice length doesn't match expected")
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("'%s' doesn't match expected '%s'", word, expectedWord)
            }
        }
    }
}