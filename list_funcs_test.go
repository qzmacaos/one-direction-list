package one_direction_list

import "testing"

func TestOneDirectionList_Add(t *testing.T) {

	odlist := New("first")

	second := odlist.Add("second", nil)

	_ = odlist.Add("fourth", second)

	_ = odlist.Add("third", second)

	odlist.Print()
}