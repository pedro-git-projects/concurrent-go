package pkg

import "reflect"

func Contains[T any](s []T, el T) bool {
	for _, e := range s {
		if reflect.DeepEqual(e, el) {
			return true
		}
	}
	return false
}
