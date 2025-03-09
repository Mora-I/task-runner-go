package tasks_test

import (
	"testing"
	"github.com/Mora-I/task-runner-go/tasks"
)

// TestGetQueue checks if the returned Queue is empty and initialized
func TestGetQueue(t *testing.T) {
	queue := tasks.GetQueue()
	if queue.Tasks == nil {
		t.Error("returned nil queue")
	}
	if len(queue.Tasks) != 0 {
		t.Error("returned queue with elements")
	}
}

// TestAddTask checks if the task is added correctly
func TestAddTask(t *testing.T) {
	queue := tasks.GetQueue()
	task := tasks.NewTask(1, tasks.Download)

	result := queue.AddTask(task)

	if result.Id != task.Id {
		t.Errorf("returned value: %d expected value: %d\n", task.Id, result.Id)
	}
}

// TestGetTask checks for the first task of queue being extracted correctly
func TestGetTask(t *testing.T) {
	queue := tasks.GetQueue()

	//attempt with an empty queue, err expected
	result, err := queue.GetTask()
	if err == nil{
		t.Error("func returning values when queue is empty")
	}

	//attempt with an only one task in queue
	task1 := tasks.NewTask(1, tasks.Download)
	queue.AddTask(task1)
	result, _ = queue.GetTask()

	if result.Id != task1.Id {
		t.Errorf("returned value: %d expected value: %d\n", task1.Id, result.Id)
	}

	//attempt with many tasks in queue
	task2 := tasks.NewTask(2, tasks.ProcessImage)
	task3 := tasks.NewTask(3, tasks.SendEmail)

	queue.AddTask(task3)
	queue.AddTask(task2)
	queue.AddTask(task1)

	result, _ = queue.GetTask()

	if result.Id != task3.Id {
		t.Errorf("returned value: %d expected value: %d\n", task3.Id, result.Id)
	}
}

// TestRemoveTaskByID enzures that passing an unvalid id returns false
func TestRemoveTaskByID(t *testing.T) {
	//testing with a single task in queue
	queue := tasks.GetQueue()
	task1 := tasks.NewTask(1, tasks.Download)

	queue.AddTask(task1)

	//invalid case
	result := queue.RemoveTaskByID(3)
	if result {
		t.Error("returned false expected true")
	}

	//valid case
	result = queue.RemoveTaskByID(1)
	if !result {
		t.Error("returned true expected false")
	}

	//now with multiple tasks in queue
	task2 := tasks.NewTask(2, tasks.ProcessImage)
	task3 := tasks.NewTask(3, tasks.SendEmail)

	queue.AddTask(task1)
	queue.AddTask(task2)
	queue.AddTask(task3)

	//invalid case
	result = queue.RemoveTaskByID(4)
	if result {
		t.Error("returned true expected false")
	}

	//valid case
	result = queue.RemoveTaskByID(2)
	if !result {
		t.Error("returned false expected true")
	}
}
