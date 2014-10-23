package main

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("Nouvelle tâche")
	taskIsNotNil(task, t)
}

func taskIsNotNil(task *Task, t *testing.T) {
	if task == nil {
		t.Fatalf("Can't create new Task")
	}
	if task.Done {
		t.Fatalf("Can't create task with default value")
	}
}

func TestNewTaskEmptyTitle(t *testing.T) {
	task := NewTask("")
	taskIsNotNil(task, t)
	expectedTitle := "No Title"
	if task.Title != expectedTitle {
		t.Errorf("Task title is not as expected: %q", task.Title)
	}
}

func TestSaveOneTask(t *testing.T) {
	task := NewTask("Learn go")
	taskIsNotNil(task, t)
	m := NewTaskManager()
	m.Save(task)

	tl := m.All()
	if len(tl) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tl))
	}
	if *tl[0] != *task {
		t.Errorf("Expected %v, got %v", task, tl[0])
	}
}

func TestSaveTwoTasks(t *testing.T) {
	t1 := NewTask("Learn go")
	t2 := NewTask("Learn go again")

	m := NewTaskManager()
	m.Save(t1)
	m.Save(t2)

	tl := m.All()
	if len(tl) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tl))
	}
	if *tl[0] != *t1 && *tl[1] != *t1 {
		t.Errorf("Could not find %v in manager %v %v", t1, tl[0], tl[1])
	}
	if *tl[0] != *t2 && *tl[1] != *t2 {
		t.Errorf("Could not find %v in manager", t2)
	}
}

func TestTaskIdAreGrowing(t *testing.T) {
	t1 := NewTask("Learn go")
	t2 := NewTask("Learn go again")

	m := NewTaskManager()
	m.Save(t1)
	m.Save(t2)

	if expected := t1.Id + 1; t2.Id != expected {
		t.Errorf("Autoincrement doesn't work")
	}
}

func TestSaveModifyAndGet(t *testing.T) {
	task := NewTask("Learn go")

	m := NewTaskManager()
	m.Save(task)

	task.Title = "Unlearn go"

	if m.All()[0].Title == "Unlearn go" {
		t.Errorf("Saved task should not be updated")
	}
}

func TestCantSaveTwiceSameTask(t *testing.T) {
	task := NewTask("Learn go")

	m := NewTaskManager()
	m.Save(task)
	m.Save(task)

	if len(m.All()) != 1 {
		t.Errorf("Must not save twice the same task %v", task)
	}
}

func TestSaveAndFindById(t *testing.T) {
	task := NewTask("Learn go")
	m := NewTaskManager()
	m.Save(task)

	tmTask := m.Find(task.Id)

	if *task != *tmTask {
		t.Errorf("Can't find task by id")
	}

}
