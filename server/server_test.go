package server

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/ant1m/todo/tasks"
)

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

func TestGetTasks(t *testing.T) {
	tmanager := tasks.NewTaskManager()
	tmanager.Save(tasks.NewTask("task1"))
	tmanager.Save(tasks.NewTask("task2"))
	context := &AppContext{
		tmanager,
	}
	fr := &fakeRespWriter{}
	req, _ := http.NewRequest("GET", "bla", bytes.NewReader([]byte("bla")))
	handler := tasksHandler(context)
	handler(fr, req)
	json := string(fr.spy)
	if json != "[]\n" {
		t.Error("Tasks should contains tasks, but contain", json)
	}
}
