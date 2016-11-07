package llr

import (
    "testing"
)

//  Test the statuses
func TestStatus(t *testing.T) {
    cases := []struct {
        in Status
        want string
    }{
        {ORDERED, "ORDERED"},
        {OTHER, "OTHER"},
        {FOR_SALE_DISCOUNT, "FOR SALE DISCOUNTED"},
    }
    for _, c := range cases {
        got := c.in.String()
        if got != c.want {
            t.Errorf("Status %q returned %q; wanted %q", c.in, got, c.want)
        }
    }
}

//  Test the priorities
func TestPriority(t *testing.T) {
    cases := []struct {
        in Priority
        want string
    }{
        {NONE, "NONE"},
        {LOW, "LOW"},
        {NORMAL, "NORMAL"},
        {HIGH, "HIGH"},
    }
    for _, c := range cases {
        got := c.in.String()
        if got != c.want {
            t.Errorf("Priority %q returned %q; wanted %q", c.in, got, c.want)
        }
    }
}
