package apiv1specs

import (
	"context"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
)

func (c *Client) GetToken(ctx context.Context, sharedSecret string) (*string, error) {
	resp, err := c.PostAuthenticate(ctx, PostAuthenticateJSONRequestBody{
		Secret: c.SaltPassword(sharedSecret),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticate: %s", err.Error())
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get token, got %q", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %s", err.Error())
	}

	token := string(body)

	return &token, nil
}

func (c *Client) SaltPassword(password string) string {
	salt := make([]byte, 32)
	rand.Read(salt)

	hash := sha256.New()
	hash.Write(append([]byte(password), salt...))

	b64Salt := b64.StdEncoding.EncodeToString(salt)
	b64Hash := b64.StdEncoding.EncodeToString(hash.Sum(nil))

	return fmt.Sprintf("%s|%s", b64Salt, b64Hash)
}
