package cmd

import (
	"bytes"
	"fmt"
	//"github.com/stretchr/testify/assert"
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

	fmt.Printf("%s", out)
	//assert.Equal(t, "Hello world!", string(out))
}
