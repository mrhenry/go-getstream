// +build !go1.7

package getstream

import "net/http"

var GETSTREAM_TRANSPORT = &http.Transport{
	MaxIdleConnsPerHost: 5,
	DisableKeepAlives:   false,
}
