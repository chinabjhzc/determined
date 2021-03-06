package scheduler

import "github.com/determined-ai/determined/master/pkg/actor"

// Group-related cluster level messages.
// TODO: Consider explicit AddGroup messages.
type (
	// groupStopped notifies that the group actor is stopped.
	groupStopped struct {
		Ref *actor.Ref
	}
	// SetMaxSlots sets the maximum number of slots that a group can consume in the cluster.
	SetMaxSlots struct{ MaxSlots *int }
	// SetWeight sets the weight of a group in the fair share scheduler.
	SetWeight struct{ Weight float64 }
)

// group manages the common state for a set of tasks that all share the same scheduling restrictions
// (e.g. max slots or fair share weight).
type group struct {
	handler  *actor.Ref
	maxSlots *int
	weight   float64
}
