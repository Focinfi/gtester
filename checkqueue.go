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
