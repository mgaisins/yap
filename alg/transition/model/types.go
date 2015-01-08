package model

import (
	. "chukuparser/alg/featurevector"
	"chukuparser/alg/perceptron"
	. "chukuparser/alg/transition"
	// "fmt"
	// "sort"
	// "strings"
)

// func (f *FeaturesList) String() string {
// 	cur := f
// 	results := make([]string, 0, 10)
// 	for cur != nil && cur.Previous != nil {
// 		results = append(results, fmt.Sprintf("Trans,Size: %v, %v", cur.Transition, len(cur.Features)))
// 		cur = cur.Previous
// 	}
// 	// sort.Reverse(sort.StringSlice(results))
// 	return strings.Join(results, "\n")
// }

type Interface interface {
	perceptron.Model
	TransitionScore(transition Transition, features []Feature) float64
}

func MakeFeature(transition, i int, feat interface{}) interface{} {
	return [3]interface{}{transition, i, feat}
}
