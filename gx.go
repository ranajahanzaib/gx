/**
 * Copyright (c) 2022-present, Rana Jahanzaib
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package gx

import (
	"errors"
	"fmt"
)

// Logs a message to the console
func Log(msg string) error {
	if msg == "" {
		return errors.New("Log expects a non-empty parameter type string")
	}
	fmt.Println(msg)
	return nil
}

func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}
