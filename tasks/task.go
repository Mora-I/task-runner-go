// provides a system to manage tasks
package tasks

//Task represents a task with an identifier, name, action and a completed state
type Task struct {
	Id        int    //unique identifier
	Action    func() //task action
	completed bool   //completed state
}

//NewTask builds and returns a task given an id, name and action
func NewTask(id int, action func()) Task {
	return Task{Id: id, Action: action, completed: false}
}

//ExecuteTask resolves the given task, calling its action and changing its state to completed
func (t *Task) ExecuteTask() bool {
	t.Action()
	t.completed = true
	return t.completed
}
