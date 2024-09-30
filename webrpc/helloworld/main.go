package helloworld

import (
	"context"
	"fmt"

	"github.com/matthiasbruns/go-server-comparison/webrpc/server"
)

type HelloWorldV1Service struct{}

func (h *HelloWorldV1Service) Simple(ctx context.Context) (string, error) {
	return "Hello World", nil
}

func (h *HelloWorldV1Service) Greet(ctx context.Context, user *server.User) (string, error) {
	output := fmt.Sprintf("Hello %s. You are %d years old and your sex is %q.", user.Name, user.Age, user.Gender)
	if user.Address != nil {
		city := user.Address.City
		if user.Address.PostalCode != nil {
			city += " " + *user.Address.PostalCode
		}
		output += fmt.Sprintf("You live in %s, %s / %s", user.Address.Street, city, user.Address.Country)
	}
	return output, nil
}
