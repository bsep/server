package main

import (
	"errors"
	"strings"
)

func splitCreds(userpass string) (error, []string) {
	arr := strings.SplitN(userpass, ":", 2)
	if len(arr) == 1 {
		return errors.New("Credentials require username and password"), nil
	} else {
		return nil, arr
	}
}
