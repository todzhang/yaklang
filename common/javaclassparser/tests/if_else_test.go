package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaklang/yaklang/common/javaclassparser"
	"github.com/yaklang/yaklang/common/javaclassparser/classes"
	"testing"
)

func TestIfElse(t *testing.T) {
	classesContent, err := classes.FS.ReadFile("test/IfTest.class")
	if err != nil {
		t.Fatal(err)
	}
	expectSource, err := classes.FS.ReadFile("test/IfTest.java")
	if err != nil {
		t.Fatal(err)
	}
	cf, err := javaclassparser.Parse(classesContent)
	if err != nil {
		t.Fatal(err)
	}
	source, err := cf.Dump()
	if err != nil {
		t.Fatal(err)
	}
	println(source)
	assert.Equal(t, string(expectSource), source)
}

func TestLoopIf(t *testing.T) {
	classesContent, err := classes.FS.ReadFile("test/LoopTest.class")
	if err != nil {
		t.Fatal(err)
	}
	cf, err := javaclassparser.Parse(classesContent)
	if err != nil {
		t.Fatal(err)
	}
	source, err := cf.Dump()
	if err != nil {
		t.Fatal(err)
	}
	println(source)
	assert.Equal(t, string(loopIdExpect), source)
}

const (
	loopIdExpect = `package org.benf.cfr.reader;

public class LoopTest {
	public LoopTest() {

	}
	public void main() {
		int var1 = 0;
		if ((var1) < (3)){
			var1 = 100
			(var1) += (1)
			goto: 4
		}
		var2 = new java.util.ArrayList();
		java.util.Iterator var3 = var2.iterator();
		if (var3.hasNext()){
			Object var4 = var3.next()
			goto: 18
		}
		var5 = 0;
		if ((var5) < (10)){
			(var2) += (1)
			goto: 27
		}
		goto: 30;
		if ((var5) >= (10)){

		}else{
			(var2) += (1)
			goto: 33
		}
		return;
	}
}`
)