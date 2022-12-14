package nut_test

import (
	"context"
	"log"
	"net"
	"nut"
	"nut/gen/proto"
	"nut/internal"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var client *nut.Client

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func init() {
	testDBName := "nut_test.db"
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	nutService := &internal.NutService{}
	nutService.Init(&testDBName, internal.CreateLogger())
	nutService.Cleanup()
	proto.RegisterNutServiceServer(s, nutService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	conn, _ := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	client, _ = nut.New("locahost:8999", "ash")
	client.ForceConnect(conn)
}

func Test_BuilderWithoutNamespace(t *testing.T) {
	_, err := client.Build("one").
		WithExpression("0 30 * * * * *").
		Nudge()
	if err == nil {
		t.Error("Error should be thrown if namespace is not given")
	}
}

func Test_BuilderWithoutName(t *testing.T) {
	_, err := client.Build("").
		WithExpression("0 30 * * * * *").
		Nudge()
	if err == nil {
		t.Error("Error should be thrown if task name is not given")
	}
}

func Test_BuilderWithoutAnySchedule(t *testing.T) {
	_, err := client.Build("one").
		Nudge()
	if err == nil {
		t.Error("Error should be thrown if schedule is not given")
	}
}

func Test_BuilderWithWrongSchedule(t *testing.T) {
	_, err := client.Build("one").
		WithExpression("0 30 * * * a *").
		Nudge()
	if err == nil {
		t.Error("Cron parser not throwing error on invalid cron expression")
	}
}

func Test_BuilderSuccess(t *testing.T) {
	resp, err := client.Build("one").
		WithExpression("0 30 * * * * *").
		Target("localhost:4000").
		Nudge()
	t.Log(resp)
	if err != nil {
		t.Error(err)
	}
}

func Test_DuplicateTaskInNamespace(t *testing.T) {
	resp, err := client.Build("one").
		WithExpression("0 30 * * * * *").
		Target("localhost:4000").
		Nudge()
	t.Log(resp)
	if err == nil {
		t.Error("should have raised unique constraint error")
	}
}
