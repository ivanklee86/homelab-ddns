package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"version"})
	rootCmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Regexp(t, regexp.MustCompile("v\\d+\\.\\d+\\.\\d+"), string(out))

}
