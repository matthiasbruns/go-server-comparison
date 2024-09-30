package main_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/matthiasbruns/go-server-comparison/webrpc/client"
	"github.com/matthiasbruns/go-server-comparison/webrpc/helloworld"
	"github.com/matthiasbruns/go-server-comparison/webrpc/server"
)

func Test_Server(t *testing.T) {
	sut := server.NewHelloWorldV1Server(&helloworld.HelloWorldV1Service{})
	srv := httptest.NewServer(sut)

	c := client.NewHelloWorldV1Client(srv.URL, http.DefaultClient)
	ctx := context.Background()

	t.Run("Simple", func(t *testing.T) {
		response, err := c.Simple(ctx)
		require.NoError(t, err)
		require.Equal(t, "Hello World", response)
	})

	t.Run("Complex", func(t *testing.T) {
		user := &client.User{
			Name:   "Werner",
			Age:    42,
			Gender: client.Gender_OTHER,
			Address: &client.Address{
				Street:     "Abbey Road",
				PostalCode: nil,
				City:       "London",
				Country:    "England",
			}}
		response, err := c.Greet(ctx, user)
		require.NoError(t, err)
		require.Equal(t, "Hello Werner. You are 42 years old and your sex is \"OTHER\".You live in Abbey Road, London / England", response)
	})
}
