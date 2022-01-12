package apiv1

import (
	"context"
	"encoding/json"
	"errors"
	apiv1specs "github.com/jkuettner/go-audioserve-client/pkg/api/v1/specs"
)

func (c *Client) GetCollections(ctx context.Context) (*apiv1specs.CollectionsInfo, error) {
	resp, err := c.client.GetCollections(ctx, c.withTokenHeader)
	if err != nil {
		return nil, err
	}

	body, err := c.parseBody(resp)
	if err != nil {
		return nil, err
	}

	var result apiv1specs.CollectionsInfo

	return &result, json.Unmarshal(body, &result)
}

func (c *Client) GetTranscodings(ctx context.Context) (*apiv1specs.TranscodingsInfo, error) {
	resp, err := c.client.GetTranscodings(ctx, c.withTokenHeader)
	if err != nil {
		return nil, err
	}

	body, err := c.parseBody(resp)
	if err != nil {
		return nil, err
	}

	var result apiv1specs.TranscodingsInfo
	return &result, json.Unmarshal(body, &result)
}

func (c *Client) GetColIdCoverPath(ctx context.Context, colId apiv1specs.CollectionId, path apiv1specs.Path) (*string, error) {
	return c.callStringReturning(ctx, colId, path, c.client.GetColIdCoverPath, c.withImageResponse(), c.withTokenHeader)
}

func (c *Client) GetColIdDescPath(ctx context.Context, colId apiv1specs.CollectionId, path apiv1specs.Path) (*string, error) {
	return c.callStringReturning(ctx, colId, path, c.client.GetColIdDescPath, c.withAnyTextResponse(), c.withTokenHeader)
}

func (c *Client) GetFolder(ctx context.Context, colId apiv1specs.CollectionId, params *apiv1specs.GetColIdFolderParams) (*apiv1specs.AudioFolder, error) {
	resp, err := c.client.GetColIdFolder(ctx, colId, params, c.withTokenHeader)
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

	var result apiv1specs.AudioFolder
	return &result, json.Unmarshal(body, &result)
}

func (c *Client) GetFolderPath(ctx context.Context, colId apiv1specs.CollectionId, path apiv1specs.Path, params *apiv1specs.GetColIdFolderPathParams) (*apiv1specs.AudioFolder, error) {
	resp, err := c.client.GetColIdFolderPath(ctx, colId, path, params, c.withTokenHeader)
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

	var result apiv1specs.AudioFolder
	return &result, json.Unmarshal(body, &result)
}

func (c *Client) GetColIdAudioPath(ctx context.Context, colId apiv1specs.CollectionId, path apiv1specs.Path, params *apiv1specs.GetColIdAudioPathParams) (*apiv1specs.AudioResponse, error) {
	resp, err := c.client.GetColIdAudioPath(ctx, colId, path, params, c.withAudioResponse(), c.withTokenHeader)
	if err != nil {
		return nil, err
	}

	transcoding, _ := resp.Header["x-transcode"]

	body, err := c.parseBody(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	return &apiv1specs.AudioResponse{
		Data:        body,
		Transcoding: transcoding,
	}, nil
}

func (c *Client) GetColIdSearch(ctx context.Context, colId apiv1specs.CollectionId, params *apiv1specs.GetColIdSearchParams) (*apiv1specs.SearchResult, error) {
	resp, err := c.client.GetColIdSearch(ctx, colId, params, c.withTokenHeader)
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

	var result apiv1specs.SearchResult
	return &result, json.Unmarshal(body, &result)
}

func (c *Client) GetColIdRecent(ctx context.Context, colId apiv1specs.CollectionId) (*apiv1specs.AudioFolder, error) {
	resp, err := c.client.GetColIdRecent(ctx, colId, c.withTokenHeader)
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

	var result apiv1specs.AudioFolder
	return &result, json.Unmarshal(body, &result)
}

func (c *Client) GetColIdDownloadPath(ctx context.Context, colId apiv1specs.CollectionId, path apiv1specs.Path, params *apiv1specs.GetColIdDownloadPathParams) ([]byte, error) {
	resp, err := c.client.GetColIdDownloadPath(ctx, colId, path, params, c.withApplicationZipResponse(), c.withTokenHeader)
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

	return body, nil
}
