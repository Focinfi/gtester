package gtester

type CheckQueue struct {
	Checks []func() error
	Err    error
}

// Add add check func to Checks in CheckQueue
func (cq *CheckQueue) Add(check func() error) *CheckQueue {
	cq.Checks = append(cq.Checks, check)
	return cq
}

// Run invoke check in Checks successively,
// stop checking and return the first error
func (cq *CheckQueue) Run() error {
	for _, check := range cq.Checks {
		if cq.Err = check(); cq.Err != nil {
			return cq.Err
		}
	}

	return nil
}

// NewCheckQueue allocates and return a new CheckQueue
func NewCheckQueue() *CheckQueue {
	return &CheckQueue{Checks: []func() error{}}
}

type Inspector struct {
	checkQueue *CheckQueue
}

// NewInspector allocates and returns a new Inspector
func NewInspector() *Inspector {
	return &Inspector{NewCheckQueue()}
}

// Check receive checking to checkQueue
func (i *Inspector) Check(checking func() error) *Inspector {
	i.checkQueue.Add(checking)
	return i
}

// Then execute the do if all checking passes
func (i *Inspector) Then(do func()) error {
	i.checkQueue.Run()
	if i.checkQueue.Err == nil {
		do()
	}
	return i.checkQueue.Err
}
