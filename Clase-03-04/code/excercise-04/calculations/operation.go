package calculations

type OperationFunc func(...int) int

const (
	Minimum = "minimum"
	Maximum = "maximum"
	Average = "average"
)

func Operation(op string) (OperationFunc, string) {
	switch op {
	case Minimum:
		return MinimumFunc, ""
	case Maximum:
		return MaximumFunc, ""
	case Average:
		return AverageFunc, ""
	default:
		return nil, "Operation not found"
	}
}
