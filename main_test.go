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
	resp, err := http.Get("http://localhost:3000/users/6161b8eea54474fe02d9f96c")
	if err != nil {
		t.Fatal(err)
	}

    body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
    	t.Fatal(err)
   	}

	
	sb := string(body)
	expected := `{"Email":"shivanshsharma2012@gmail.com","Name":"Shivansh","ID":"6161b8eea54474fe02d9f96c"}`
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

func TestCreatePostRoute(t *testing.T){
	postBody , _ := json.Marshal(map[string]string{
		"Caption":"This is a test caption",
    	"ImageURL":"http://www.testimagehere.jpg",
    	"UserID":"6161b8eea54474fe02d9f96c",
	})
	reqBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:3000/posts", "application/json", reqBody)
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

func TestGetPostById(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/posts/6161d48bee1f6959e748faeb")
	if err != nil {
		t.Fatal(err)
	}

    body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
    	t.Fatal(err)
   	}

	
	sb := string(body)
	
	if sb == "Internal error" {
		t.Error("Server is not reachable")
	}
}

func TestGetPostByUserId(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/posts/users/6161b8eea54474fe02d9f96c")
	if err != nil {
		t.Fatal(err)
	}

    body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
    	t.Fatal(err)
   	}

	
	sb := string(body)
	
	if sb == "Internal error" {
		t.Error("Server is not reachable")
	}
}

