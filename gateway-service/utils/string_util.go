package utils

import "slices"

func ContainsString(slice []string, target string) bool {
	return slices.Contains(slice, target)
}
