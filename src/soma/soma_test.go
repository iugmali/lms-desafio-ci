package main

import "testing"

func TestSum(t *testing.T) {
    total := Soma(5, 5)
    if total != 10 {
       t.Errorf("Soma errada, foi: %d, tinha que ser: %d.", total, 10)
    }
}