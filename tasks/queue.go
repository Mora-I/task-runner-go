package tasks

import (
	"errors"
	"fmt"
)

//Queue represents a task queue
type Queue struct {
	Tasks []Task //Stores the list of tasks
}

//GetQueue returns a new instance of an empty queue 
func GetQueue() Queue {
	return Queue{
		Tasks: []Task{},
	}
}

//AddTask appends a given task to the queue
func (q *Queue) AddTask(t Task) Task {
	q.Tasks = append(q.Tasks, t)
	return t
}

//GetTask extracts and returns the first task from the queue
func (q *Queue) GetTask() (Task, error) {
	if len(q.Tasks) == 0 {
		return Task{}, errors.New("can't return task from an empty queue")
	}
	t := q.Tasks[0]
	q.Tasks = q.Tasks[1:]
	return t, nil
}

//RemoveTaskByID deletes a task with the given id in the queue
func (q *Queue) RemoveTaskByID(id int) bool {
	for i, task := range q.Tasks {
		if task.Id == id {
			q.Tasks = append(q.Tasks[:i], q.Tasks[i+1:]...)
			fmt.Printf("task %d removed successfully\n", id)
			return true
		}
	}
	//No match found for the given id
	fmt.Println("ID not found")
	return false
}
