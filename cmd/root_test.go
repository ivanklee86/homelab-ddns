package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestRoot(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello world!", string(out))

}
