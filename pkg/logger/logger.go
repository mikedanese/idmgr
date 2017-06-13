package logger

import (
	"flag"

	_ "google.golang.org/grpc/grpclog/glogger"
)

func init() {
	flag.Set("alsologtostderr", "true")
}
