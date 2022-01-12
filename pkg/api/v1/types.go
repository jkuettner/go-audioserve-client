package apiv1

import (
	"context"
	apiv1specs "github.com/jkuettner/go-audioserve-client/pkg/api/v1/specs"
	"net/http"
	"time"
)

type ClientOpts struct {
	ServerURL      string
	SharedSecret   string
	RequestTimeout time.Duration
}

type Client struct {
	opts   *ClientOpts
	client *apiv1specs.Client
	token  string
}

type genericColIdPathMethod func(context.Context, apiv1specs.CollectionId, apiv1specs.Path, ...apiv1specs.RequestEditorFn) (*http.Response, error)
