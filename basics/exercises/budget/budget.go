package budget

var stateManager = newStateTransitionManager()

type Budget struct {
	amount    int
	maxAmount int
	currency  string
	state     State
}

func NewBudget(amount int, maxAmount int, currency string) *Budget {
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
	if err := stateManager.checkTransition(b.state, ActiveState); err != nil {
		return err
	}
	b.state = ActiveState
	return nil
}

func (b *Budget) Pause() error {
	if err := stateManager.checkTransition(b.state, PausedState); err != nil {
		return err
	}
	b.state = PausedState
	return nil
}

func (b *Budget) Unpause() error {
	if b.amount > 0 {
		return b.Activate()
	}

	if err := stateManager.checkTransition(b.state, EmptyState); err != nil {
		return err
	}
	b.state = EmptyState
	return nil
}
