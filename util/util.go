package util

import (
	"fmt"
	"net/url"
)

func PullFromQuery(field string, vs url.Values) (string, error) {
	if valueArr, ok := vs[field]; ok && len(valueArr[0]) > 1 {
		return valueArr[0], nil
	}
	return "", fmt.Errorf("Invalid or missing %s.", field)
}
