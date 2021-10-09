package main

import (
	"net/http"
	"io/ioutil"
	"testing"
)

func TestGetUserById(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/users/61615b5ea1bc982a4eae9545")
	if err != nil {
		t.Fatal(err)
	}

    body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
    	t.Fatal(err)
   	}

	
	sb := string(body)
	expected := `{"Email":"shivanshsharma2012@gmail.com","Name":"Shivansh","ID":"61615b5ea1bc982a4eae9545"}`
	if sb != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",sb, expected)
	}
}

