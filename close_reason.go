package pool

type CloseReason string

func (c CloseReason) String() string {
	return string(c)
}

const (
	CloseReasonPoolClosing          CloseReason = "pool_closed"
	CloseReasonError                CloseReason = "error"
	CloseReasonMaxAgeExceeded       CloseReason = "max_age_exceeded"
	CloseReasonIdleTimeoutExceeded  CloseReason = "idle_timeout_exceeded"
	CloseReasonPoolCapacityExceeded CloseReason = "capacity_exceeded"
)
