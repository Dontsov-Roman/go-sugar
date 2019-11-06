package request

// Condition for Where method
type Condition struct {
	Column         string
	Operator       string
	Value          string
	ConcatOperator string
	Native         bool
}

// NewCond - return a new Condition for Where methods
func (r *Request) NewCond(column string, operator string, value string) *Condition {
	return &Condition{Column: column, Operator: operator, Value: value, ConcatOperator: "OR", Native: false}
}

// NewCondition - return a new Condition for Where methods
func (r *Request) NewCondition(column string, operator string, value string, concatOperator string, native bool) *Condition {
	condition := Condition{Column: column, Operator: operator, Value: value, Native: native}
	if concatOperator != "" {
		condition.ConcatOperator = concatOperator
	} else {
		condition.ConcatOperator = "OR"
	}
	return &condition
}
