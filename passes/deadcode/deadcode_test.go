package deadcode

import (
	"testing"

	"github.com/krageon/php/ast"
	"github.com/krageon/php/parser"
)

func TestDeadFunctions(t *testing.T) {
	src := `<?php
	$var1 = "a";
	function simple() {
		$var2 = "b";
		$var3 = "c";
	}

	class fizz {
		const buzz = "fizzbuzz";

		static function notsimple() {
			$var4 = "d";
		}

		function other() {}
	}

	fizz::notsimple();
	`

	p := parser.NewParser()
	if _, err := p.Parse("test.php", src); err != nil {
		t.Fatal(err)
	}

	var shouldBeDead = map[string]struct{}{
		"simple": {},
		"other":  {},
	}

	dead := DeadFunctions(p.FileSet, []string{"test.php"})

	for _, deadFunc := range dead {

		fnName := deadFunc.(*ast.FunctionStmt).Name
		if _, ok := shouldBeDead[fnName]; !ok {
			t.Error("%q was found dead, but shouldn't have been", fnName)
		}
		delete(shouldBeDead, fnName)
	}

	for fugitive := range shouldBeDead {
		t.Error("%q should have been found dead, but wasn't", fugitive)
	}
}
