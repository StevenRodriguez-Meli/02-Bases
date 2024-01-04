package taxes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTax(t *testing.T) {
	t.Run("success - case 01: salary below 50000", func(t *testing.T) {
		// arrange
		// ...

		// act
		salary := 40000.0
		expectedTax := 0.0

		result := CalculateTax(salary)

		// assert
		require.Equal(t, expectedTax, result)
	})

	t.Run("success - case 01: salary above 50000", func(t *testing.T) {
		// arrange
		// ...

		// act
		salary := 60000.0
		expectedTax := 10200.0

		result := CalculateTax(salary)

		// assert
		require.Equal(t, expectedTax, result)
	})

	t.Run("success - case 01: salary above 150000", func(t *testing.T) {
		// arrange
		// ...

		// act
		salary := 180000.0
		expectedTax := 48600.0

		result := CalculateTax(salary)

		// assert
		require.Equal(t, expectedTax, result)
	})
}
