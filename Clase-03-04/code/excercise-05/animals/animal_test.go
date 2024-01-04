package animals

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var quantity = 10.0

func TestDog(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := Dog(quantity)
		resultExpect := 100.0
		require.Equal(t, resultExpect, result)
	})
}

func TestCat(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := Cat(quantity)
		resultExpect := 50.0
		require.Equal(t, resultExpect, result)
	})
}

func TestHamster(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := Hamster(quantity)
		resultExpect := 2.5
		require.Equal(t, resultExpect, result)
	})
}

func TestTarantula(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := Tarantula(quantity)
		resultExpect := 1.5
		require.Equal(t, resultExpect, result)
	})
}
