package sailthru

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

// Checksum creates md5 checksum of given json data.
func (c *Client) Checksum(jsonData string) string {
	values := []string{c.Key, DefaultFormat}
	if jsonData != "" {
		values = append(values, jsonData)
	}
	sort.Strings(values)
	s := c.Secret + strings.Join(values, "")
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
