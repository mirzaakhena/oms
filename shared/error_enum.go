package shared

const (
	MenuItemCodeMustNotEmptyError          ErrorType = "ER1042 Menu Item Code Must Not Empty"
	QuantityMustNotEmptyError              ErrorType = "ER1042 Quantity Must Not Empty"
	OrderIDLengthMust4Char                 ErrorType = "ER1042 OrderID must 4 fix char"
	SequenceMustGreaterThanZero            ErrorType = "ER1042 SequenceMustGreaterThanZero"
	SequenceOutOfBound                     ErrorType = "ER1042 SequenceOutOfBound"
	OrderlineMustNotEmptyError             ErrorType = "ER1042 Orderline must not empty"
	UnrecognizedPaymentMethod              ErrorType = "ER1042 Unrecognized payment method"
	OrderStateMustNotEmptyError            ErrorType = "ER1042 Order State Must Not Empty"
	NotAllowedOrderStateTransitionError    ErrorType = "ER1042 Not Allowed Order State Transition from %s to %s"
	PaymentStatusMustNotEmptyError         ErrorType = "ER1042 Status Must Not Empty"
	NotAllowedPaymentStatusTransitionError ErrorType = "ER1042 Not Allowed Payment Status Transition"
	PhoneNumberMustNotEmptyError           ErrorType = "ER1042 Phone Number Must Not Empty"
	OrderIDMustNotEmptyError               ErrorType = "ER1042 Order ID Must Not Empty"
	InvalidDateError                       ErrorType = "ER1042 Invalid Date"
	UserIsNotActive                        ErrorType = "ER1042 User is not active"
	UserIsNotPremium                       ErrorType = "ER1042 User is not premium"
	AmountMustGreaterThanZeroError         ErrorType = "ER1042 Amount Must Greater Than Zero"
	BalanceIsNotEnoughError                ErrorType = "ER1042 Balance Is Not Enough"
)
