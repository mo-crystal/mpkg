package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mo-crystal/mpkg/config"
)

func Get(path string) Response {
	resp, err := http.Get(config.Server + path)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ret := Response{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		panic(err)
	}

	return ret
}
