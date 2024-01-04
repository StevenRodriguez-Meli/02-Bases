package grades

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateAverage(t *testing.T) {

	t.Run("PositiveGrades", func(t *testing.T) {
		gradesList := []int{90, 85, 95, 100}
		resultExpect := 92.5
		result := CalculateAverage(gradesList...)
		require.Equal(t, resultExpect, result)
	})

	t.Run("MixedGrades", func(t *testing.T) {
		gradesList := []int{70, 60, 80, -1, 75}
		resultExpect := 71.25
		result := CalculateAverage(gradesList...)
		require.Equal(t, resultExpect, result)
	})

	t.Run("EmptyGrades", func(t *testing.T) {
		resultExpect := 0.0
		result := CalculateAverage()
		require.Equal(t, resultExpect, result)
	})

	t.Run("AllNegativeGrades", func(t *testing.T) {
		gradesList := []int{-10, -20, -30, -40}
		resultExpect := 0.0
		result := CalculateAverage(gradesList...)
		require.Equal(t, resultExpect, result)
	})
}
