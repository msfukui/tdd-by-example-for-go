# TDD by Example for Go

書籍「テスト駆動開発」を Golang で読み進めてみるレポジトリです。

## Setup

for mac. (macports)

```
$ sudo port install go
--->  Cleaning go
--->  Scanning binaries for linking errors
...
```

## A test sample

sample_test.go
```
package sample

import "testing"

func TestSay(t *testing.T) {
	const in, out = "world", "Hello world."

	if x := Say(in); x != out {
		t.Errorf("Say(%v) = %v, want %v", in, x, out)
	}
}
```

sample.go
```
package sample

// Say returns a greeting message.
func Say(message string) string {
	return "Hello " + message + "."
}
```

## Test

```
$ go test
PASS
ok  	github.com/msfukui/tdd-by-example-for-go	0.007s
```

## References
