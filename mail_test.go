package mail

import (
	"strings"
	"testing"
)

// Converts all newlines to CRLFs.
func crlf(s string) string {
	return strings.Replace(s, "\n", "\r\n", -1)
}

type getHeadersTest struct {
	orig string
	hdrs []string
	body string
}

var getHeadersTests = []getHeadersTest{
	{
		`a: b`,
		[]string{`a: b`},
		``,
	},
	{
		crlf(`a: b
`),
		[]string{`a: b`},
		``,
	},
	{
		crlf(`a: b

`),
		[]string{`a: b`},
		``,
	},
	{
		crlf(`a: b
c: d
e: f

Body goes
here.
`),
		[]string{`a: b`, `c: d`, `e: f`},
		crlf(`Body goes
here.
`),
	},
	{
		crlf(`a: b
 c`),
		[]string{crlf(`a: b
 c`)},
		``,
	},
}

func TestGetHeaders(t *testing.T) {
	for i, ht := range getHeadersTests {
		hs, b := getHeaders(ht.orig)
		if len(hs) != len(ht.hdrs) {
			t.Errorf(`%d. getHeaders returned %d headers, wanted %d`,
				i, len(hs), len(ht.hdrs))
		} else {
			for j, h := range hs {
				if h != ht.hdrs[j] {
					t.Errorf(`%d. getHeaders [%d] gave "%s", wanted "%s"`,
						i, j, h, ht.hdrs[j])
				}
			}
		}
		if b != ht.body {
			t.Errorf(`%d. getHeaders [body] gave "%s", wanted "%s"`,
				i, b, ht.body)
		}
	}
}

type splitHeadersTest struct {
	orig, key, val string
}

var splitHeadersTests = []splitHeadersTest{
	{
		`a: b`,
		`a`, `b`,
	},
	{
		`A1: cD`,
		`A1`, `cD`,
	},
	{
		crlf(`ab: cd
 ef`),
		`ab`, `cd ef`,
	},
	{
		crlf(`ab: cd
	ef
	dh`),
		`ab`, `cd	ef	dh`,
	},
}

func TestSplitHeader(t *testing.T) {
	for i, ht := range splitHeadersTests {
		k, v := splitHeader(ht.orig)
		if k != ht.key || v != ht.val {
			t.Errorf(`%d. splitHeader gave ("%s", "%s"), wanted ("%s", "%s")`,
				i, k, v, ht.key, ht.val)
		}
	}
}
