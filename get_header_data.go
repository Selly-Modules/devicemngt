package devicemngt

import "net/http"

// getHeaderData ...
func getHeaderData(headers http.Header) HeaderData {
	return HeaderData{
		UserAgent:      headers.Get("User-Agent"),
		DeviceID:       headers.Get("Deviceid"),
		AppVersion:     headers.Get("App-Version"),
		AppVersionCode: headers.Get("App-Version-Code"),
		OSName:         headers.Get("Os-Name"),
		OSVersion:      headers.Get("Os-Version"),
	}
}
