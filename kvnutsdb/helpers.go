package kvnuts

import "reflect"

func CheckItemsArePointers(items ...any) int {
	for ix := range items {
		if items[ix] == nil {
			return ix
		}

		if reflect.ValueOf(items[ix]).Kind() != reflect.Ptr {
			return ix
		}
	}

	return -1
}
