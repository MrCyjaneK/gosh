package gosh_ls_test

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"testing"

	gosh "git.mrcyjanek.net/mrcyjanek/gosh/_core"
)

// This test make sure that the command 'works'
// Take this file as an example, when all tests
// will pass it should mean that the package is
// working exacly as the binary that we try to
// package into gosh work.

// TestCore - Test if the command 'works'
func TestCore(t *testing.T) {
	gosh.Handlecmd([]string{"ls"}, []byte{})
	if gosh.ERRCODE != 0 {
		t.Fail()
	}
}

// TestNoArgAll | Files with '.' at begging should not be displayed
func TestNoArgAll(t *testing.T) {
	var err error
	os.Chdir("..")
	gosh.CWD, err = os.MkdirTemp(os.TempDir(), "gosh_ls_testnoargall")
	os.MkdirAll(path.Join(gosh.CWD, "/.invisible"), 0755)

	if err != nil {
		log.Println(err)
		t.Fail()
	}
	stdout, _ := gosh.Handlecmd([]string{"ls"}, []byte{})
	if gosh.ERRCODE != 0 {
		t.Fail()
	}
	if stdout.Writer.Buffered() != 0 {
		stdoutbyte := make([]byte, stdout.Writer.Buffered()-1)
		stdout.Flush()
		_, err = stdout.Read(stdoutbyte)
		if err != nil {
			log.Println(err)
			t.Fail()
		}
		if strings.Contains(string(stdoutbyte), ".invisible") {
			fmt.Println("'.invisible' directory was visible.")
			t.Fail()
		}
	}
}

// TestArgAll | -a, --all do not ignore entries starting with .
func TestArgAll(t *testing.T) {
	var err error
	os.Chdir("..")
	gosh.CWD, err = os.MkdirTemp(os.TempDir(), "gosh_ls_testnoargall")
	os.MkdirAll(path.Join(gosh.CWD, "/.invisible"), 0755)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	stdout, _ := gosh.Handlecmd([]string{"ls", "--all"}, []byte{})
	if gosh.ERRCODE != 0 {
		t.Fail()
	}
	if stdout.Writer.Buffered() == 0 {
		fmt.Println("No output")
		t.FailNow()
	}
	stdoutbyte := make([]byte, stdout.Writer.Buffered()-1)
	stdout.Flush()
	_, err = stdout.Read(stdoutbyte)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	if !strings.Contains(string(stdoutbyte), ".invisible") {
		fmt.Println("'.invisible' directory was not visible.")
		t.Fail()
	}
}
