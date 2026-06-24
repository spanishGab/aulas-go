package budget

import (
	"testing"
)

type StateTransitionTest struct {
	name          string
	budget        Budget
	expectedState State
	expectedErr   bool
}

type AmountChangeTest struct {
	name           string
	budget         Budget
	expectedBudget Budget
	amountToChange int
	expectedErr    bool
}

func TestActivate(t *testing.T) {
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
				amount:    0,
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
			name: "should make the budget EMPTY when it is CREATED and its amount is equal to 0",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			expectedState: EmptyState,
			expectedErr:   false,
		},
		{
			name: "should make the budget EMPTY when it is CREATED and its amount is less than 0",
			budget: Budget{
				amount:    -200,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			expectedState: EmptyState,
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
			assertError(t, test.expectedErr, result)
			assertBudgetStateTransition(t, test)
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
			assertError(t, test.expectedErr, result)
			assertBudgetStateTransition(t, test)
		})
	}
}

func TestFinish(t *testing.T) {
	tests := []StateTransitionTest{
		{
			name: "should finish an EMPTY budget successfully",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			expectedState: FinishedState,
			expectedErr:   false,
		},
		{
			name: "should not be able to finish an CREATED budget",
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
			name: "should not be able to finish an ACTIVE budget",
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
			name: "should not be able to finish a PAUSED budget",
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
			name: "should not be able to finish a FINISHED budget",
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
			result := test.budget.Finish()
			assertError(t, test.expectedErr, result)
			assertBudgetStateTransition(t, test)
		})
	}
}

func TestSpend(t *testing.T) {
	tests := []AmountChangeTest{
		{
			name: "should spend from a positive budget successfully",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedBudget: Budget{
				amount:    500,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			amountToChange: 500,
			expectedErr:    false,
		},
		{
			name: "should spend from a positive budget successfully and transition its state to EMPTY if its whole amount was spent",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedBudget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			amountToChange: 1000,
			expectedErr:    false,
		},
		{
			name: "should not be able to spend from an EMPTY budget",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			expectedBudget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			amountToChange: 1,
			expectedErr:    true,
		},
		{
			name: "should not be able to spend from a CREATED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			amountToChange: 1000,
			expectedErr:    true,
		},
		{
			name: "should not be able to spend from a PAUSED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			amountToChange: 1000,
			expectedErr:    true,
		},
		{
			name: "should not be able to spend from a FINISHED budget",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     FinishedState,
			},
			expectedBudget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     FinishedState,
			},
			amountToChange: 1000,
			expectedErr:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.budget.Spend(test.amountToChange)
			assertError(t, test.expectedErr, result)
			assertBudgetAmountChange(t, test)
		})
	}
}

func TestRecharge(t *testing.T) {
	tests := []AmountChangeTest{
		{
			name: "should be able to recharge an ACTIVE budget successfully",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedBudget: Budget{
				amount:    1500,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			amountToChange: 500,
			expectedErr:    false,
		},
		{
			name: "should be able to recharge an EMPTY budget successfully",
			budget: Budget{
				amount:    0,
				maxAmount: 5000,
				currency:  "R$",
				state:     EmptyState,
			},
			expectedBudget: Budget{
				amount:    1500,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			amountToChange: 1500,
			expectedErr:    false,
		},
		{
			name: "should not be able to recharge above the maximum amount allowed",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			amountToChange: 5001,
			expectedErr:    true,
		},
		{
			name: "should not be able to recharge a zero amount",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			amountToChange: 0,
			expectedErr:    true,
		},
		{
			name: "should not be able to recharge a negative amount",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			amountToChange: -500,
			expectedErr:    true,
		},
		{
			name: "should not be able to recharge a CREATED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     CreatedState,
			},
			amountToChange: 500,
			expectedErr:    true,
		},
		{
			name: "should not be able to recharge a PAUSED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			amountToChange: 500,
			expectedErr:    true,
		},
		{
			name: "should not be able to recharge a FINISHED budget",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			expectedBudget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     PausedState,
			},
			amountToChange: 500,
			expectedErr:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.budget.Recharge(test.amountToChange)
			assertError(t, test.expectedErr, result)
			assertBudgetAmountChange(t, test)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name           string
		budget         Budget
		expectedResult string
	}{
		{
			name: "should return a budget string representation successfully",
			budget: Budget{
				amount:    1000,
				maxAmount: 5000,
				currency:  "R$",
				state:     ActiveState,
			},
			expectedResult: "Budget amount: R$ 10.00; maxAmount: R$ 50.00; state: ACTIVE",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.budget.String()
			if result != test.expectedResult {
				t.Errorf("expected: %q, got: %q", test.expectedResult, result)
			}
		})
	}
}

func assertError(t testing.TB, expectedErr bool, result error) {
	t.Helper()
	if expectedErr && result == nil {
		t.Errorf("expected an error to occur, got no error")
	}
}

func assertBudgetStateTransition(t testing.TB, test StateTransitionTest) {
	t.Helper()
	if test.expectedState != test.budget.state {
		t.Errorf("expected budget state to be %q, but it's %q", test.expectedState, test.budget.state)
	}
}

func assertBudgetAmountChange(t testing.TB, test AmountChangeTest) {
	t.Helper()
	if test.expectedBudget.state != test.budget.state {
		t.Errorf("expected budget state to be %q, but it's %q", test.expectedBudget.state, test.budget.state)
	}
	if test.expectedBudget.amount != test.budget.amount {
		t.Errorf("expected budget amount to be %d, but it's %d", test.expectedBudget.amount, test.budget.amount)
	}
}
