// Copyright 2015 The XC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xc // import "modernc.org/xc"

import (
	"go/token"

	"modernc.org/golex/lex"
	"modernc.org/strutil"
)

// Token describes a token.
type Token struct {
	lex.Char     // Token rune and position.
	Val      int // ID of token value in the global Dict variable.
}

// S returns a R/O value of t, ie. Dict.S(t.Val).
func (t Token) S() []byte { return Dict.S(t.Val) }

// String implements fmt.Stringer.
func (t Token) String() string { return strutil.PrettyString(t, "", "", PrintHooks) }

// Position returns position of t.
func (t *Token) Position() token.Position { return FileSet.Position(t.Pos()) }
