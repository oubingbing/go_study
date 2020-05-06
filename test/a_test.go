package tgo

import (
	"go_study/test/t_string"
	"testing"
	"os/exec"
)


func TestStr(t *testing.T)  {
	if t_string.StringL("abc") != 3 {
		t.Error("测试失败")
	}

	exec.Command("go run C:/Users/Administrator/go/src/go_study/bytes.go ")
}
