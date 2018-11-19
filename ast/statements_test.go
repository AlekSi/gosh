// Gosh programming language.
// Copyright (c) 2018 Alexey Palazhchenko and contributors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"testing"

	"gosh-lang.org/gosh/tokens"

	"github.com/stretchr/testify/assert"
)

func TestSwitchStatementString(t *testing.T) {
	for expected, ss := range map[string]*SwitchStatement{
		"switch {\n}": {},
		"switch {\ncase 1:\n}": {
			Cases: []CaseClause{
				{
					List: []Expression{
						&IntegerLiteral{
							Token: tokens.Token{Literal: "1"},
						},
					},
				},
			},
		},
		"switch {\ncase 1:\nfoo;\n}": {
			Cases: []CaseClause{
				{
					List: []Expression{
						&IntegerLiteral{
							Token: tokens.Token{Literal: "1"},
						},
					},
					Body: []Statement{
						&ExpressionStatement{
							Expression: &Identifier{
								Token: tokens.Token{Literal: "foo"},
							},
						},
					},
				},
			},
		},
	} {
		t.Run(expected, func(t *testing.T) {
			assert.Equal(t, expected, ss.String())
		})
	}
}
