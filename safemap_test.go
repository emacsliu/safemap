package safemap

import (
	"reflect"
	"sync"
	"testing"
)

var sm *SafeMap

func init() {
	sm = NewSafeMap()
	sm.objects["111"] = 111
	sm.objects["222"] = "222"
	sm.objects["333"] = 333.333
}

func TestSafeMap_GetObject(t *testing.T) {
	value1 := sm.GetObject("111")
	if value1.(int) != 111 {
		t.Error("value1")
	}
	value2 := sm.GetObject("222")
	if value2.(string) != "222" {
		t.Error("value2")
	}
	value3 := sm.GetObject("333")
	if value3.(float64) != 333.333 {
		t.Error("value3")
	}
}

func TestSafeMap_SetObject(t *testing.T) {

}

func TestSafeMap_RemoveObject(t *testing.T) {

}

func TestSafeMap_Values(t *testing.T) {

}

func TestSafeMap_Keys(t *testing.T) {

}

func TestSafeMap_invalidKey(t *testing.T) {

}
