package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

var ip_regex = ".*\\d{1,}\\.\\d{1,}\\.\\d{1,}\\.\\d{1,}\\."

func TestUpdate(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"update"})
	rootCmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Regexp(t, regexp.MustCompile(ip_regex), string(out))
}
