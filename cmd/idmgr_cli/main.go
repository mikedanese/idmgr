package main

import (
	"encoding/json"
	"flag"
	"io"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	api "containers.google.com/idmgr/pkg/api"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	path := "/tmp/idmgr.sock"

	conn, err := grpc.DialContext(ctx, path,
		grpc.WithInsecure(),
		grpc.WithTimeout(100*time.Second),
		grpc.WithDialer(func(address string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", address, timeout)
		}),
	)
	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}
	c := api.NewManagementClient(conn)

	switch os.Args[1] {
	case "init":
		output{Status: "Success"}.write(os.Stdout)
	case "getvolumename":
		output{Status: "Success", VolumeName: "mds-" + RandStringBytes(10)}.write(os.Stdout)
	case "mount":
		r, err := c.CreateMetadata(ctx, &api.CreateMetadataRequest{
			MountPath: os.Args[2],
		})
		if err != nil {
			glog.Fatalf("fail to create mds: %v", err)
		}
		glog.Infof("response: %#v", r)
		output{Status: "Success"}.write(os.Stdout)
	case "unmount":
		r, err := c.DestroyMetadata(ctx, &api.DestroyMetadataRequest{
			MountPath: os.Args[2],
		})
		if err != nil {
			glog.Fatalf("fail to destroy mds: %v", err)
		}
		glog.Infof("response: %#v", r)
		output{Status: "Success"}.write(os.Stdout)
	default:
		output{Status: "Not supported"}.write(os.Stdout)
		os.Exit(1)
	}
}

type output struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	VolumeName string `json:"volumeName"`
	Attached   bool   `json:"attached"`
}

func (o output) write(w io.Writer) {
	b, err := json.Marshal(o)
	if err != nil {
		glog.Fatalf("failed to marshal status: %v", err)
	}
	w.Write(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
