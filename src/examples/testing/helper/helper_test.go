package helper

import "testing"

func TestExportedFunc(t *testing.T) {
	res := ExportedFunc()
	if res != 2 {
		t.Errorf("expect %v, instead got %v", 2, res)
	}
}
