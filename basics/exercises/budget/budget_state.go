package budget

import "fmt"

const (
	CreatedState  State = "CREATED"
	ActiveState   State = "ACTIVE"
	PausedState   State = "PAUSED"
	EmptyState    State = "EMPTY"
	FinishedState State = "FINISHED"
)

type State string

type stateTransitions map[State]struct{}

type stateTransitionManager struct {
	allowedTransitions map[State]stateTransitions
}

func newStateTransitionManager() *stateTransitionManager {
	return &stateTransitionManager{
		allowedTransitions: map[State]stateTransitions{
			CreatedState: {ActiveState: {}, EmptyState: {}},
			ActiveState:  {PausedState: {}, EmptyState: {}},
			PausedState:  {ActiveState: {}, EmptyState: {}},
			EmptyState:   {PausedState: {}, ActiveState: {}, FinishedState: {}},
		},
	}
}

func (s *stateTransitionManager) checkTransition(from, to State) error {
	if _, ok := s.allowedTransitions[from]; !ok {
		return fmt.Errorf("invalid state %q", from)
	}
	if _, ok := s.allowedTransitions[from][to]; !ok {
		return fmt.Errorf("cannot transition from %q to %q", from, to)
	}
	return nil
}
