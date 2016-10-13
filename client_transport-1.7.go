// +build go1.7

package getstream

import "net/http"

var GETSTREAM_TRANSPORT = &http.Transport{
	MaxIdleConns:        5,
	MaxIdleConnsPerHost: 5,
	IdleConnTimeout:     60,
	DisableKeepAlives:   false,
}
