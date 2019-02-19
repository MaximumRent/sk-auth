package util

// EDRIT YA ZAGNUL

const (
	FIRST_LEVEL = 0
)

type EntityOp func(interface{}) error

// Chain of operation which will use for one entity
// For example: prepare user data, save user to db
type EntityOpChain struct {
	operation       EntityOp
	entityType      string
	executionResult error
	nextOp          *EntityOpChain
	chainHead       *EntityOpChain
}

func GetEmptyOpChain(eType string) *EntityOpChain {
	return &EntityOpChain{entityType: eType, operation: nil}
}

func (chain *EntityOpChain) AddOperation(operation EntityOp) *EntityOpChain {
	if chain.operation == nil {
		chain.operation = operation
		chain.chainHead = chain
		return chain
	} else {
		chain.nextOp = &EntityOpChain{
			entityType: chain.entityType,
			operation:  operation,
		}
		return chain.nextOp
	}
}

// input here, is non-immutable object
func (chain *EntityOpChain) Execute(input interface{}) error {
	return chain.ExecuteStartsOnLevel(input, FIRST_LEVEL)
}

// Can be used to skip some execution levels
func (chain *EntityOpChain) ExecuteStartsOnLevel(input interface{}, currentLevel int) error {
	if currentLevel == FIRST_LEVEL {

	} else {
		return chain.chainHead.ExecuteStartsOnLevel(input, currentLevel+1)
	}
	//if chain == chain.chainHead {
	//	err := chain.operation(input)
	//	if err != nil {
	//		return err
	//	}
	//	chain.nextOp.ExecuteStartsOnLevel(input, currentLevel + 1)
	//} else {
	//	if currentLevel != FIRST_LEVEL {
	//		err := chain.operation(input)
	//		if err != nil {
	//			return err
	//		}
	//		chain.nextOp.ExecuteStartsOnLevel(input, currentLevel + 1)
	//	}
	//}
	return nil
}
