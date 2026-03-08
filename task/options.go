package task

// AddOptions allows specifying optional parameters when creating a task.
// The pointer-for-optional pattern: if Priority is nil, use the default (PriorityMedium).
type AddOptions struct {
	Priority *Priority
	Tags     []string
}
