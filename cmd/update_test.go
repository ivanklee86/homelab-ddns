package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestUpdate(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"update"})
	rootCmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "WIP", string(out))

}
