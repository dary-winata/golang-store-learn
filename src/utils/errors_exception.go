package utils

import "log"

func ErrorException(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorMethodInput() {}
