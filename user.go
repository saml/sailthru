package sailthru

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// UserID uniquely identifies a user.
type UserID struct {
	SID    string `json:"sid"`
	Cookie string `json:"cookie"`
	Email  string `json:"email"`
}

// SubscribedLists is a map from list name to subscription date string.
type SubscribedLists map[string]string

// User has subscribed lists
type User struct {
	IDs             *UserID         `json:"keys"`
	SubscribedLists SubscribedLists `json:"lists"`
	Engagement      string          `json:"engagement"`
	OptOutEmail     string          `json:"optout_email"`
}

// FetchUser fetches a user with given id.
func (c *Client) FetchUser(userID string) (*User, error) {
	res, err := c.Get(UserURL, Data{
		"id": userID,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var user User
	d := json.NewDecoder(res.Body)
	err = d.Decode(&user)
	if err != nil {
		return nil, err
	}

	// IDs field is required.
	// Sailthru does repspond with 200 status without IDs field. So, don't need to check for status code.
	if user.IDs == nil {
		return nil, fmt.Errorf("No ID can be fetched for user: %v", userID)
	}

	return &user, nil
}

// SubscriptionStatus is a map of mailinglist name to 0 and 1.
// 1 means subscribed. 0 means unsubscribed.
type SubscriptionStatus map[string]uint

// UpdateSubscription subscribes and unsubscribes the user to and from mailinglist.
func (c *Client) UpdateSubscription(userID string, lists SubscriptionStatus) error {
	for listName := range lists {
		if _, ok := c.Lists[listName]; !ok {
			return fmt.Errorf("Unknown list: %v", listName)
		}
	}

	res, err := c.Post(UserURL, Data{
		"id":    userID,
		"lists": lists,
	})
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		s := string(b)
		return fmt.Errorf("Failed to update subscription for %v: %v", userID, s)
	}
	return nil
}
