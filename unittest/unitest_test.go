package unittest

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{"niumowang", 400, "niumoquan"}
	monster.Store()
}

func TestRestore(t *testing.T) {
	monster := Monster{}
	monster.ReStore()
	if monster.Name == "niumowang" {
		t.Logf("Stroe and Restroe 执行正确。。。")
	} else {
		t.Fatalf("Stroe and Restroe 执行错误。。。monster.Name = %v", monster.Name)
	}
}
