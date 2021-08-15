package one_direction_list

import "testing"

func TestOneDirectionList_Add(t *testing.T) {
	odlist := New("first")
	_ = odlist.Add("second")
	_ = odlist.Add("third")

	if odlist.Root.Next.Next.Name != "third"{
		t.Error("third  is not the third")
	}
}

func TestOneDirectionList_AddAfter(t *testing.T) {

	odlist := New("first")

	second := odlist.Add("second")

	_ = odlist.AddAfter("fourth", second)

	_ = odlist.AddAfter("third", second)

	if odlist.Root.Next.Next.Name != "third"{
		t.Error("third  is not the third")
	}
	if odlist.Root.Next.Next.Next.Name != "fourth"{
		t.Error("fourth  is not the fourth")
	}
}