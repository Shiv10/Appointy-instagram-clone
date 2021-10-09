package main

import (
	"net/http"
	"io/ioutil"
	"testing"
	"encoding/json"
	"bytes"
	// "fmt"
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

func TestCreateUserRoute(t *testing.T){
	postBody , _ := json.Marshal(map[string]string{
		"Username":"Shivansh",
    	"Password":"mango5741",
    	"Email":"shivanshsharma2012@gmail.com",
	})
	reqBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:3000/users", "application/json", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	//Read the response body
  	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	sb := string(body)
	if sb=="Internal error" {
		t.Error("Server is not reachable")
	}
}

