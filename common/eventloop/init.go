package eventloop

import "widesocks/common/slog"

func init() {
	if err := epollInit(); err != nil {
		slog.Emergency(err)
	}
}
