package cafe

import "errors"

var (
	ErrOrderLimitReached = errors.New("order limit reached")
	ErrInvalidIndex      = errors.New("invalid index")
	ErrAlreadyReady      = errors.New("order is already ready")
	ErrInvalidItem       = errors.New("invalid item in order")
)
