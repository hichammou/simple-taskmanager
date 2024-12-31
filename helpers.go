package main

func contains[E []S, S comparable](elms E, elm S) bool {
	for _, e := range elms {
		if e == elm {
			return true
		}
	}
	return false
}
