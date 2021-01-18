package goginrestapi

import "testing"

func TestAdd(t *testing.T) {
	restApi := New()
	restApi.Add(Item{})

	if len(restApi.Items) != 1 {
		t.Errorf("Item was not added")
	}


}

func TestGetAll(t *testing.T) {
	restApi := New()
	restApi.Add(Item{})
	results := restApi.GetAll()
	if len(results) != 1{
		t.Errorf("Item was not added")
	}

}