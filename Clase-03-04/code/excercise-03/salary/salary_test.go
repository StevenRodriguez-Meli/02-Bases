package salary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateSalary(t *testing.T) {
	minutesWorked := 7200

	t.Run("CategoryC", func(t *testing.T) {
		category := "C"
		expectedSalary := 120000.0
		result := CalculateSalary(minutesWorked, category)
		require.Equal(t, expectedSalary, result)
	})

	t.Run("CategoryB", func(t *testing.T) {
		category := "B"
		expectedSalary := 216000.0
		result := CalculateSalary(minutesWorked, category)
		require.Equal(t, expectedSalary, result)
	})

	t.Run("CategoryA", func(t *testing.T) {
		category := "A"
		expectedSalary := 540000.0
		result := CalculateSalary(minutesWorked, category)
		require.Equal(t, expectedSalary, result)
	})
}
