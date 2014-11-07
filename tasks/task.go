package tasks

type Task struct {
	Id    int64
	Title string
	Done  bool
}

func NewTask(title string) *Task {
	if title == "" {
		title = "No Title"
	}
	return &Task{Title: title}
}

type TaskManager struct {
	tasks     []*Task
	currentId int64
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:     make([]*Task, 0),
		currentId: 1,
	}
}

func (tm *TaskManager) Save(task *Task) {
	if task.Id == 0 {
		task.Id = tm.currentId
		tm.currentId++
	}
	for _, t := range tm.tasks {
		if *t == *task {
			return
		}
	}
	taskCopy := *task
	tm.tasks = append(tm.tasks, &taskCopy)
}

func (tm TaskManager) All() []*Task {
	return tm.tasks
}

func (tm TaskManager) Find(id int64) *Task {
	for _, task := range tm.tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}
