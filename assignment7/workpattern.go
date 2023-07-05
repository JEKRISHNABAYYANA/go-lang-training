package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	ID         int
	Name       string
	Salary     int
	sync.Mutex // Embedding sync.Mutex for synchronization
}

type Worker struct {
	ID         int
	EmployeeCh chan *Employee
	QuitCh     chan bool
}

func NewWorker(id int) *Worker {
	return &Worker{
		ID:         id,
		EmployeeCh: make(chan *Employee),
		QuitCh:     make(chan bool),
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			select {
			case employee := <-w.EmployeeCh:
				// Process the employee and update salary
				employee.UpdateSalary(1000)
				fmt.Printf("Worker %d updated salary for employee %s (ID: %d)\n", w.ID, employee.Name, employee.ID)
			case <-w.QuitCh:
				fmt.Printf("Worker %d quitting\n", w.ID)
				return
			}
		}
	}()
}

func (e *Employee) UpdateSalary(increase int) {
	e.Lock()
	defer e.Unlock()
	e.Salary += increase
}

func main() {
	numWorkers := 3
	numEmployees := 5

	// Create workers
	workers := make([]*Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i)
	}

	// Start workers
	var wg sync.WaitGroup
	for _, worker := range workers {
		wg.Add(1)
		worker.Start(&wg)
	}

	// Create employees
	employees := make([]*Employee, numEmployees)
	for i := 0; i < numEmployees; i++ {
		employees[i] = &Employee{
			ID:     i + 1,
			Name:   fmt.Sprintf("Employee %d", i+1),
			Salary: 5000,
		}
	}

	// Send employees to workers for processing
	for _, employee := range employees {
		// Select a random worker
		workerIndex := employee.ID % numWorkers
		workers[workerIndex].EmployeeCh <- employee
	}

	// Signal workers to quit
	for _, worker := range workers {
		worker.QuitCh <- true
	}

	// Wait for workers to finish
	wg.Wait()

	fmt.Println("All workers finished processing employees.")

	// Print final salaries
	for _, employee := range employees {
		fmt.Printf("Employee %s (ID: %d) - Salary: %d\n", employee.Name, employee.ID, employee.Salary)
	}
}
