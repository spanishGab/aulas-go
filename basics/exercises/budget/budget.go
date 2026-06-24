package budget

import (
	"fmt"
)

var stateManager = newStateTransitionManager()

type Budget struct {
	amount    int
	maxAmount int
	currency  string
	state     State
}

func NewBudget(currency string, amount, maxAmount int) *Budget {
	state := CreatedState
	if amount <= 0 {
		state = EmptyState
	}
	return &Budget{
		amount:    amount,
		maxAmount: maxAmount,
		currency:  currency,
		state:     state,
	}
}

func (b *Budget) Activate() error {
	nextState := EmptyState
	if b.amount > 0 {
		nextState = ActiveState
	}

	if err := stateManager.checkTransition(b.state, nextState); err != nil {
		return err
	}
	b.state = nextState
	return nil
}

func (b *Budget) Pause() error {
	if err := stateManager.checkTransition(b.state, PausedState); err != nil {
		return err
	}
	b.state = PausedState
	return nil
}

func (b *Budget) Spend(amount int) error {
	if !b.IsActive() {
		return fmt.Errorf("budget is not active")
	}

	spentAmount := b.amount - amount
	if spentAmount < 0 {
		return fmt.Errorf("cannot spend more than %d, trying to spend %d", b.amount, amount)
	}

	if spentAmount == 0 {
		b.state = EmptyState
	}
	b.amount = spentAmount
	return nil
}

func (b *Budget) Recharge(amount int) error {
	if !b.IsRechargeable() {
		return fmt.Errorf("budget is not rechargeable")
	}

	if amount <= 0 {
		return fmt.Errorf("recharge amount must be greater than 0")
	}

	rechargedAmount := amount + b.amount
	if rechargedAmount > b.maxAmount {
		return fmt.Errorf("cannot recharge above %d, trying to recharge %d", b.maxAmount, amount)
	}
	b.amount = rechargedAmount
	b.state = ActiveState
	return nil
}

func (b *Budget) Finish() error {
	if err := stateManager.checkTransition(b.state, FinishedState); err != nil {
		return err
	}
	b.state = FinishedState
	return nil
}

func (b *Budget) String() string {
	amount := float64(b.amount) / 100
	maxAmount := float64(b.maxAmount) / 100
	return fmt.Sprintf("Budget amount: %[1]s %.2[2]f; maxAmount: %[1]s %.2[3]f; state: %[4]s", b.currency, amount, maxAmount, b.state.String())
}

func (b *Budget) IsActive() bool {
	return b.state == ActiveState
}

func (b *Budget) IsRechargeable() bool {
	return b.state == ActiveState || b.state == EmptyState
}
