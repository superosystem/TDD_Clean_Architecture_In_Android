package transactions

type TransactionUseCase struct {
	transactionRepository Repository
}

func NewTransactionUseCase(tr Repository) UseCase {
	return &TransactionUseCase{
		transactionRepository: tr,
	}
}

func (t TransactionUseCase) Create(domain *Domain) (*Domain, error) {
	transaction, err := t.transactionRepository.Create(domain)
	if err != nil {
		return nil, err
	}

	return transaction, err
}

func (t TransactionUseCase) GetAll() *[]Domain {
	//TODO implement me
	panic("implement me")
}
