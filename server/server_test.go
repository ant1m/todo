package server

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGetTasks(t *testing.T) {
	RunServer()
	time.Sleep(1 * time.Second)
	r, err := http.Get("http://localhost:8080/tasks")
	if err != nil {
		t.Fatal("Can't get task list:", err)
	}
	defer r.Body.Close()
	if status := r.StatusCode; status != 200 {
		t.Error("Response status should be 200 but is ", status)
	}
	if json, _ := ioutil.ReadAll(r.Body); string(json) != "[]\n" {
		t.Error("Tasks should be an empty json array, [], but was", string(json))
	}

}
