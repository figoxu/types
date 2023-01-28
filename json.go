package types

import "encoding/json"

type Json string

func JsonString(obj interface{}) Json {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return Json(b)
}
