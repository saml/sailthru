package sailthru

import (
	"fmt"
)

type jsonLeafCollector struct {
	values []string
}

func (c *jsonLeafCollector) collectSlice(l []interface{}) {
	for _, v := range l {
		c.collect(v)
	}
}

func (c *jsonLeafCollector) collectMap(m map[string]interface{}) {
	for _, v := range m {
		c.collect(v)
	}
}

func (c *jsonLeafCollector) collect(v interface{}) {
	switch v.(type) {
	case float64:
		c.values = append(c.values, fmt.Sprintf("%v", v.(float64)))
	case bool:
		if v.(bool) {
			c.values = append(c.values, "1")
		} else {
			c.values = append(c.values, "0")
		}
	case string:
		c.values = append(c.values, v.(string))
	case []interface{}:
		c.collectSlice(v.([]interface{}))
	case map[string]interface{}:
		c.collectMap(v.(map[string]interface{}))
	}
}

// Params is paramters that client takes.
type Params map[string]interface{}

// ExtractParams extracts sailthru parameters given json object.
func ExtractParams(m Params) []string {
	c := &jsonLeafCollector{}
	c.collectMap(m)
	return c.values
}
