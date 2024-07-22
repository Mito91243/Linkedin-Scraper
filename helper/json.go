package helperFunction

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	//"strings"
)

func Decoding(res *http.Response) []byte {

	var body []byte
	var err error
	if res.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(res.Body)
		if err != nil {
			fmt.Println("Failed to create gzip reader:", err)
			return nil
		}
		defer gzipReader.Close()
		body, err = io.ReadAll(gzipReader)
		if err != nil {
			fmt.Println("Failed to read gzipped body:", err)
			return nil
		}
	} else {

		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return nil
		}
	}
	return body
}

// For Future API
func EncodeProfiles(profiles []ProfileRes) {
	_, err := json.Marshal(profiles)
	if err != nil {
		return
	}
}

func safeGetString(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}
