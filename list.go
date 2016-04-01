package sailthru

import (
	"encoding/json"
	"fmt"
	"time"
)

// List is a mailinglist.
type List struct {
	ID              string    `json:"list_id"`
	Name            string    `json:"name"`
	EmailCount      uint      `json:"email_count"`
	ValidEmailCount uint      `json:"valid_count"`
	Created         time.Time `json:"create_date"`
	Type            string    `json:"type"`
}

type allLists struct {
	Items []*List `json:"lists"`
}

// FetchLists fetches all available mailinglists.
func (c *Client) FetchLists() ([]*List, error) {
	res, err := c.Get(ListURL, nil)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s | %s", res.Request.Method, res.Request.URL.String(), res.Status)
	}

	var lists allLists
	d := json.NewDecoder(res.Body)
	err = d.Decode(&lists)
	if err != nil {
		return nil, err
	}
	return lists.Items, nil
}
