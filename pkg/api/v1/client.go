package apiv1

import apiv1specs "github.com/jkuettner/go-audioserve-client/pkg/api/v1/specs"

func NewClient(opts *ClientOpts) (*Client, error) {
	httpHandler := newHTTPHandler(opts.RequestTimeout)
	client, err := apiv1specs.NewClient(opts.ServerURL, apiv1specs.WithHTTPClient(httpHandler))
	if err != nil {
		return nil, err
	}

	return &Client{
		opts:   opts,
		client: client,
		token:  "",
	}, nil
}
