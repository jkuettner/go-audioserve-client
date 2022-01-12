package apiv1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	apiv1specs "github.com/jkuettner/go-audioserve-client/pkg/api/v1/specs"
)

func (c *Client) GetPositionsGroup(ctx context.Context, group apiv1specs.GroupInPath, params *apiv1specs.GetPositionsGroupParams) (*apiv1specs.Position, error) {
	resp, err := c.client.GetPositionsGroup(ctx, group, params, c.withTokenHeader)
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

	var result *apiv1specs.Position
	return result, json.Unmarshal(body, &result)
}

func (c *Client) PostPositionsGroup(ctx context.Context, group apiv1specs.GroupInPath, body apiv1specs.PostPositionsGroupJSONRequestBody) error {
	resp, err := c.client.PostPositionsGroup(ctx, group, body, c.withTokenHeader)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("failed to post positions: %s", err.Error())
	}

	return nil
}

func (c *Client) GetPositionsGroupLast(ctx context.Context, group apiv1specs.GroupInPath) (*apiv1specs.Position, error) {
	resp, err := c.client.GetPositionsGroupLast(ctx, group, c.withTokenHeader)
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

	var result *apiv1specs.Position
	return result, json.Unmarshal(body, &result)
}

func (c *Client) GetPositionsGroupColIdPath(ctx context.Context, group apiv1specs.GroupInPath, colId apiv1specs.CollectionId, path apiv1specs.Path, params *apiv1specs.GetPositionsGroupColIdPathParams) (*apiv1specs.Position, error) {
	resp, err := c.client.GetPositionsGroupColIdPath(ctx, group, colId, path, params, c.withTokenHeader)
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

	var result *apiv1specs.Position
	return result, json.Unmarshal(body, &result)
}
