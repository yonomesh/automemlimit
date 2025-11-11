package automemlimit

import (
	"log/slog"

	"github.com/yonomesh/automemlimit/memlimit"
)

func init() {
	memlimit.SetGoMemLimitWithOpts(
		memlimit.WithLogger(slog.Default()),
	)
}
