// +build !debug

package logx_test

import (
	"testing"
	"bytes"
	"github.com/tamtam-im/logx"
	"github.com/stretchr/testify/assert"
)

func TestStandaloneLogger_Debug(t *testing.T) {
	w := &bytes.Buffer{}
	l := logx.NewLog(logx.NewSimpleAppender(w, 0), "")
	l.Debug("test")
	assert.Equal(t, "", w.String())
}
