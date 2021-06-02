package employee

import "fmt"

// non exported struct.It is implemented using New constructor
type employee struct {
	firstname   string
	lastname    string
	totalLeaves int
	leavesTaken int
}

// employee's method
func (e employee) LeaveRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstname, e.lastname, (e.totalLeaves - e.leavesTaken))
}

// used as a constructor
func New(firstname, lastname string, totalLeaves, leavesTaken int) employee {
	e := employee{firstname, lastname, totalLeaves, leavesTaken}
	return e
}
