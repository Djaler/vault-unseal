// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"errors"
)

func notify(err error) {
	err = errors.New("error: " + err.Error())
	logger.WithError(err).Error("notify-error")
}
