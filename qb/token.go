// Copyright (C) 2017 ScyllaDB
// Use of this source code is governed by a ALv2-style
// license that can be found in the LICENSE file.

package qb

import (
	"fmt"
	"strings"
)

// TokenBuilder helps implement pagination using token function.
type TokenBuilder []string

// Token creates a new TokenBuilder.
func Token(columns ...string) TokenBuilder {
	return TokenBuilder(columns)
}

// Eq produces token(column)=token(?).
func (t TokenBuilder) Eq() Cmp {
	return t.cmp(eq, nil)
}

// EqNamed produces token(column)=token(?) with a custom parameter name.
func (t TokenBuilder) EqNamed(names ...string) Cmp {
	return t.cmp(eq, names)
}

// Lt produces token(column)<token(?).
func (t TokenBuilder) Lt() Cmp {
	return t.cmp(lt, nil)
}

// LtNamed produces token(column)<token(?) with a custom parameter name.
func (t TokenBuilder) LtNamed(names ...string) Cmp {
	return t.cmp(lt, names)
}

// LtOrEq produces token(column)<=token(?).
func (t TokenBuilder) LtOrEq() Cmp {
	return t.cmp(leq, nil)
}

// LtOrEqNamed produces token(column)<=token(?) with a custom parameter name.
func (t TokenBuilder) LtOrEqNamed(names ...string) Cmp {
	return t.cmp(leq, names)
}

// Gt produces token(column)>token(?).
func (t TokenBuilder) Gt() Cmp {
	return t.cmp(gt, nil)
}

// GtNamed produces token(column)>token(?) with a custom parameter name.
func (t TokenBuilder) GtNamed(names ...string) Cmp {
	return t.cmp(gt, names)
}

// GtOrEq produces token(column)>=token(?).
func (t TokenBuilder) GtOrEq() Cmp {
	return t.cmp(geq, nil)
}

// GtOrEqNamed produces token(column)>=token(?) with a custom parameter name.
func (t TokenBuilder) GtOrEqNamed(names ...string) Cmp {
	return t.cmp(geq, names)
}

func (t TokenBuilder) cmp(op op, names []string) Cmp {
	s := names
	if s == nil {
		s = t
	}
	return Cmp{
		op:     op,
		column: fmt.Sprint("token(", strings.Join(t, ","), ")"),
		value:  Fn("token", s...),
	}
}

// TokenExactBuilder helps implement pagination using exact comparison to token value.
type TokenExactBuilder []string

// TokenExact creates a new TokenExactBuilder.
func TokenExact(columns ...string) TokenExactBuilder {
	return TokenExactBuilder(columns)
}

// Eq produces token(column)=?.
func (t TokenExactBuilder) Eq() Cmp {
	return t.cmp(eq, "")
}

// EqNamed produces token(column)=? with a custom parameter name.
func (t TokenExactBuilder) EqNamed(name string) Cmp {
	return t.cmp(eq, name)
}

// Lt produces token(column)<?.
func (t TokenExactBuilder) Lt() Cmp {
	return t.cmp(lt, "")
}

// LtNamed produces token(column)<? with a custom parameter name.
func (t TokenExactBuilder) LtNamed(name string) Cmp {
	return t.cmp(lt, name)
}

// LtOrEq produces token(column)<=?.
func (t TokenExactBuilder) LtOrEq() Cmp {
	return t.cmp(leq, "")
}

// LtOrEqNamed produces token(column)<=? with a custom parameter name.
func (t TokenExactBuilder) LtOrEqNamed(name string) Cmp {
	return t.cmp(leq, name)
}

// Gt produces token(column)>?.
func (t TokenExactBuilder) Gt() Cmp {
	return t.cmp(gt, "")
}

// GtNamed produces token(column)>? with a custom parameter name.
func (t TokenExactBuilder) GtNamed(name string) Cmp {
	return t.cmp(gt, name)
}

// GtOrEq produces token(column)>=?.
func (t TokenExactBuilder) GtOrEq() Cmp {
	return t.cmp(geq, "")
}

// GtOrEqNamed produces token(column)>=? with a custom parameter name.
func (t TokenExactBuilder) GtOrEqNamed(name string) Cmp {
	return t.cmp(geq, name)
}

func (t TokenExactBuilder) cmp(op op, name string) Cmp {
	s := name
	if s == "" {
		s = "token"
	}
	return Cmp{
		op:     op,
		column: fmt.Sprint("token(", strings.Join(t, ","), ")"),
		value:  param(s),
	}
}
