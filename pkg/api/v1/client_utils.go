package apiv1

import (
	"context"
	"errors"
	"fmt"
	apiv1specs "github.com/jkuettner/go-audioserve-client/pkg/api/v1/specs"
	"io/ioutil"
	"net/http"
)

func (c *Client) getToken(ctx context.Context) error {
	token, err := c.client.GetToken(ctx, c.opts.SharedSecret)
	if err != nil {
		return err
	}
	c.token = *token
	return nil
}

func (c *Client) parseBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) withTokenHeader(ctx context.Context, req *http.Request) error {
	if c.token == "" {
		if err := c.getToken(ctx); err != nil {
			return err
		}
	}

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}
	return nil
}

func (c *Client) withAudioResponse() apiv1specs.RequestEditorFn {
	return c.withAdditionalHeaderResponse(http.Header{
		"accept": []string{
			"audio:*",
			"audio/*",
		},
	})
}

func (c *Client) withPlainTextResponse() apiv1specs.RequestEditorFn {
	return c.withAdditionalHeaderResponse(http.Header{
		"accept": []string{"text/plain"},
	})
}

func (c *Client) withAnyTextResponse() apiv1specs.RequestEditorFn {
	return c.withAdditionalHeaderResponse(http.Header{
		"accept": []string{"text/*"},
	})
}

func (c *Client) withImageResponse() apiv1specs.RequestEditorFn {
	return c.withAdditionalHeaderResponse(http.Header{
		"accept": []string{"image/*"},
	})
}

func (c *Client) withApplicationZipResponse() apiv1specs.RequestEditorFn {
	return c.withAdditionalHeaderResponse(http.Header{
		"accept": []string{"application/zip"},
	})
}

func (c *Client) withAdditionalHeaderResponse(header http.Header) apiv1specs.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		for k, v := range header {
			req.Header[k] = v
		}
		return nil
	}
}

func (c *Client) callStringReturning(ctx context.Context, colId apiv1specs.CollectionId, path apiv1specs.Path, f genericColIdPathMethod, reqEditors ...apiv1specs.RequestEditorFn) (*string, error) {
	resp, err := f(ctx, colId, path, reqEditors...)
	if err != nil {
		return nil, err
	}

	body, err := c.parseBody(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	result := string(body)

	return &result, nil
}
