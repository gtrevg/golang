// Copyright (c) 2019 The searKing authors. All Rights Reserved.
//
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file in the root of the source
// tree. An additional intellectual property rights grant can be found
// in the file PATENTS.  All contributing project authors may
// be found in the AUTHORS file in the root of the source tree.

package cgo

//go:generate make rebuild
//go:generate bash ../../../../tools/scripts/cgo_include_gen.sh -p "github.com/searKing/golang/go/os/signal/cgo/include" "./include"