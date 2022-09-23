package utils

/**
 * @file utils.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This file parses incoming requests in JSON into corresponding Go objects.
 */

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// The ParseBody() function reads the body of a request and unmarshals JSON into Go objects.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
