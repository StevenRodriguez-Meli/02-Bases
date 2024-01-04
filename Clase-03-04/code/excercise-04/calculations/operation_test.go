package calculations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimumFunc(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := MinimumFunc()
		resultExpect := 0
		require.Equal(t, resultExpect, result)
	})

	t.Run("PositiveNumbers", func(t *testing.T) {
		result := MinimumFunc(10, 20, 30, 40, 50)
		resultExpect := 10
		require.Equal(t, resultExpect, result)
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		result := MinimumFunc(-5, -10, -15, -20)
		resultExpect := -20
		require.Equal(t, resultExpect, result)
	})
}

func TestMaximumFunc(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := MaximumFunc()
		resultExpect := 0
		require.Equal(t, resultExpect, result)
	})

	t.Run("PositiveNumbers", func(t *testing.T) {
		result := MaximumFunc(10, 20, 30, 40, 50)
		resultExpect := 50
		require.Equal(t, resultExpect, result)
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		result := MaximumFunc(-5, -10, -15, -20)
		resultExpect := -5
		require.Equal(t, resultExpect, result)
	})
}

func TestAverageFunc(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		result := AverageFunc()
		resultExpect := 0
		require.Equal(t, resultExpect, result)
	})

	t.Run("PositiveNumbers", func(t *testing.T) {
		result := AverageFunc(10, 20, 30, 40, 50)
		resultExpect := 30
		require.Equal(t, resultExpect, result)
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		result := AverageFunc(-5, -10, -15, -20)
		resultExpect := -12
		require.Equal(t, resultExpect, result)
	})
}
