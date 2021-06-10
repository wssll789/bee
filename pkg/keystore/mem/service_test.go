// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/wssll789/bee/pkg/keystore/mem"
	"github.com/wssll789/bee/pkg/keystore/test"
)

func TestService(t *testing.T) {
	test.Service(t, mem.New())
}
