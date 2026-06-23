package budget

import "testing"

type StateTransitionTest struct {
	name          string
	budget        Budget
	expectedState State
	expectedErr   bool
}

func TestActivate(t *testing.T) {
	tests := []StateTransitionTest{
		{
			name:          "should activate a CREATED budget successfully",
			budget:        *NewBudget(1000, 5000, "R$"),
			expectedState: ActiveState,
			expectedErr:   false,
		},
		{
			name: "should activate a PAUSED budget successfully",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedState: ActiveState,
			expectedErr:   false,
		},
		{
			name: "should not be able to activate an EMPTY budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			expectedState: ActiveState,
			expectedErr:   false,
		},
		{
			name: "should not be able to activate an ACTIVE budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedState: ActiveState,
			expectedErr:   true,
		},
		{
			name: "should not be able to activate a FINISHED budget",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     FinishedState,
			},
			expectedState: FinishedState,
			expectedErr:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.budget.Activate()
			assertBudgetStateTransition(t, test, result)
		})
	}
}

func TestPause(t *testing.T) {
	tests := []StateTransitionTest{
		{
			name: "should pause an ACTIVE budget successfully",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedState: PausedState,
			expectedErr:   false,
		},
		{
			name: "should pause an EMPTY budget successfully",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			expectedState: PausedState,
			expectedErr:   false,
		},
		{
			name: "should not be able to pause a CREATED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			expectedState: CreatedState,
			expectedErr:   true,
		},
		{
			name: "should not be able to pause a PAUSED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedState: PausedState,
			expectedErr:   true,
		},
		{
			name: "should not be able to pause a FINISHED budget",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     FinishedState,
			},
			expectedState: FinishedState,
			expectedErr:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.budget.Pause()
			assertBudgetStateTransition(t, test, result)
		})
	}
}

func TestUnpause(t *testing.T) {
	tests := []StateTransitionTest{
		{
			name: "should make the budget ACTIVE when it is PAUSED and its amount is greater than 0",
			budget: Budget{
				amount:    1,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedState: ActiveState,
			expectedErr:   false,
		},
		{
			name: "should make the budget EMPTY when it is PAUSED and its amount is equal to 0",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedState: EmptyState,
			expectedErr:   false,
		},
		{
			name: "should make the budget EMPTY when it is PAUSED and its amount is less than 0",
			budget: Budget{
				amount:    -5,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedState: EmptyState,
			expectedErr:   false,
		},
		{
			name: "should make the budget ACTIVE when it is CREATED and its amount is greater than 0",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			expectedState: ActiveState,
			expectedErr:   false,
		},
		{
			name: "should not be able to unpause an ACTIVE budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedState: ActiveState,
			expectedErr:   true,
		},
		{
			name: "should not be able to unpause an EMPTY budget",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			expectedState: EmptyState,
			expectedErr:   true,
		},
		{
			name: "should not be able to unpause a FINISHED budget",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     FinishedState,
			},
			expectedState: FinishedState,
			expectedErr:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.budget.Unpause()
			assertBudgetStateTransition(t, test, result)
		})
	}
}

func assertBudgetStateTransition(t testing.TB, test StateTransitionTest, result error) {
	t.Helper()
	if test.expectedErr && result == nil {
		t.Errorf("expected an error to occur, got no error")
	}
	if test.expectedState != test.budget.state {
		t.Errorf("expected budget state to be %q, but it's %q", test.expectedState, test.budget.state)
	}
}
