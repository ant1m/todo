package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/ant1m/todo/tasks"
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

type fakeRespWriter struct {
	http.ResponseWriter
	spy []byte
}

func (f *fakeRespWriter) Write(b []byte) (int, error) {
	f.spy = b
	return len(b), nil
}

func (f *fakeRespWriter) WriteHeader(h int) {
	// does nothing, just a stub
}

func TestGetTasksHandler(t *testing.T) {
	context := &AppContext{
		tasks.NewTaskManager(),
	}
	fr := &fakeRespWriter{}
	req, _ := http.NewRequest("GET", "bla", bytes.NewReader([]byte("bla")))
	handler := tasksHandler(context)
	handler(fr, req)
	json := string(fr.spy)
	if json != "[]\n" {
		t.Error("Tasks should be an empty json array, [], but was", json)
	}
}
