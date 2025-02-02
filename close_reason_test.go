package pool

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCloseReason(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		cases := []struct {
			reason CloseReason
			want   string
		}{
			{CloseReasonError, "error"},
			{CloseReasonPoolClosing, "pool_closed"},
			{CloseReasonMaxAgeExceeded, "max_age_exceeded"},
			{CloseReasonIdleTimeoutExceeded, "idle_timeout_exceeded"},
			{CloseReasonPoolCapacityExceeded, "capacity_exceeded"},
		}

		for _, tc := range cases {
			tc := tc

			t.Run(tc.want, func(t *testing.T) {
				// act
				got := tc.reason.String()

				// assert
				require.Equal(t, tc.want, got)
			})
		}
	})
}
