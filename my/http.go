package my

import (
	"fmt"
	"net/http"
	"time"
)

// Get is http.Get with retry
func Get(url string) (resp *http.Response, err error) {
	const errorFormat = `"%s" returns http error status %d`
	const max = 3
	for i := 1; ; i++ {
		resp, err = http.Get(url)
		if err != nil {
			return
		}
		sc := resp.StatusCode
		if sc == 200 {
			return
		} else if sc >= 500 {
			if i < max {
				time.Sleep(time.Second) // retry after 1 second
			} else {
				return nil, fmt.Errorf(errorFormat, url, sc)
			}
		} else if sc >= 400 {
			return nil, fmt.Errorf(errorFormat, url, sc)
		}
	}
}
