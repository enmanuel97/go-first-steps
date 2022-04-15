package main

import "fmt"

// Lista de tareas
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

// Tareas
type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar curso de Go este mes",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de HTML",
		desc:      "Completar curso de HTML esta semana",
		completed: true,
	}

	t3 := Task{
		name:      "Curso de CSS",
		desc:      "Completar curso de CSS esta semana",
		completed: false,
	}

	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)
	lista.appendTask(&t3)

	lista.removeTask(1)

	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}
	/*t1.toPrint()
	fmt.Println("-------------------------")
	t2.toPrint()*/
}
