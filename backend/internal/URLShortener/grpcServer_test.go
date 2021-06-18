package URLShortener_test

import (
	"URLShortener/internal/URLShortener"
	"URLShortener/internal/api"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"testing"
	"time"
)

const (
	bindAddr = ":8080"
)

func TestMain(m *testing.M) {
	c := URLShortener.NewConfig()
	c.Server.BindAddr = bindAddr
	srv := URLShortener.NewGRPCServer(c)
	go srv.Start()
	time.Sleep(time.Second * 1)
	os.Exit(m.Run())
}

var tests = []struct {
	url         string
	expectedErr bool
	isMessage   bool
}{
	{
		url:         "https://job.ozon.ru",
		expectedErr: false,
		isMessage:   true,
	},
	{
		url:         "http:/yandex.ru",
		expectedErr: true,
		isMessage:   false,
	},
	{
		url:         "ftp://192.168.1.1:30",
		expectedErr: false,
		isMessage:   true,
	},
	{
		url:         "fmt.Println(https://yandex.ru)",
		expectedErr: true,
		isMessage:   false,
	},
}

func TestGRPCServer_CreateGet(t *testing.T) {
	conn, err := grpc.Dial(bindAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatal("can not connect to gRPC server:", err)
	}
	defer conn.Close()
	c := api.NewURLShortenerClient(conn)
	urlmap := make(map[string]string, 0)
	for _, test := range tests {
		shortCode, err := c.Create(context.Background(), &api.OriginalURL{Url: test.url})
		if test.expectedErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		if test.isMessage {
			urlmap[test.url] = shortCode.ShortURLCode
			assert.NotEmpty(t, shortCode.ShortURLCode)
		} else {
			assert.Nil(t, shortCode)
		}
	}
	for _, test := range tests {
		url, err := c.Get(context.Background(), &api.ShortCode{ShortURLCode: urlmap[test.url]})
		if test.expectedErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		if test.isMessage {
			assert.Equal(t, url.Url, test.url)
		} else {
			assert.Nil(t, url)
		}
	}
}
