package p6

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestNilPointer(t *testing.T) {
	if uintptr(unsafe.Pointer(nil)) != 0 {
		t.Fatal()
	}
}

func TestAdd(t *testing.T) {
	arr := []int{1, 2, 3}

	var ll XORLinkedList

	for _, v := range arr {
		ll.Add(v)
	}

	for i := 0; i < len(arr); i++ {
		v, err := ll.Get(i)
		if err != nil {
			t.Errorf("error not expected err: %v", err.Error())
		}
		if !reflect.DeepEqual(arr[i], v) {
			t.Fatal()
		}
	}

	_, err := ll.Get(0)
	if err != nil {
		t.Errorf("error not expected err: %v", err.Error())
	}

	_, err = ll.Get(2)
	if err != nil {
		t.Errorf("error not expected err: %v", err.Error())
	}

	_, err = ll.Get(3)
	if err == nil {
		t.Errorf("not error")
	}
	if _, ok := err.(ErrIndexOutOfRange); !ok {
		t.Errorf("not valid error")
	} else {
		t.Logf("expected error: %v", err.Error())
	}

	_, err = ll.Get(-1)
	if err == nil {
		t.Errorf("not error")
	}
	if _, ok := err.(ErrIndexOutOfRange); !ok {
		t.Errorf("not valid error")
	}

	ll.Add(20)

	v, err := ll.Get(3)
	if err != nil {
		t.Errorf("error not expected err: %v", err.Error())
	}
	if v.(int) != 20 {
		t.Fail()
	}
}
