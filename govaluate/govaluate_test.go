package govaluate

import (
	"fmt"
	"testing"

	"github.com/Knetic/govaluate"
)

func TestExp(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("(mem_used/1total_mem) * 100")
	fmt.Println(err)
	parameters := make(map[string]interface{}, 8)
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512

	result, err := expression.Evaluate(parameters)
	fmt.Println(result, err)
	// result is now set to "50.0", the float64 value.
}
