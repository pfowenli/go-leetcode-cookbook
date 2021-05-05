package leetcode

import (
    "testing"
)

func TestTrailingZeroesI(t *testing.T) {
    n := int(3)
    expected := int(0)

    if got := trailingZeroes(n); got != expected {
        t.Errorf("trailingZeroes() = %d, expected = %d", got, expected)
    }
}

func TestTrailingZeroesII(t *testing.T) {
    n := int(5)
    expected := int(1)

    if got := trailingZeroes(n); got != expected {
        t.Errorf("trailingZeroes() = %d, expected = %d", got, expected)
    }
}

func TestTrailingZeroesIII(t *testing.T) {
    n := int(0)
    expected := int(0)

    if got := trailingZeroes(n); got != expected {
        t.Errorf("trailingZeroes() = %d, expected = %d", got, expected)
    }
}
