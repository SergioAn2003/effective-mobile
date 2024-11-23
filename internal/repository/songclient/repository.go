package songclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
)

type Client struct {
	client  *http.Client
	baseURL string
}

func New(baseURL string) *Client {
	return &Client{
		client: &http.Client{
			Timeout: 10,
		},
		baseURL: baseURL,
	}
}

func (c *Client) GetSongDetails(ctx context.Context, songName, groupName string) (entity.SongDetails, error) {
	url := fmt.Sprintf("%s/info?group=%s&song=%s", c.baseURL, groupName, songName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entity.SongDetails{}, fmt.Errorf("songClient: failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return entity.SongDetails{}, fmt.Errorf("songClient: failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.SongDetails{}, fmt.Errorf("songClient: unexpected status code: %s", resp.Status)
	}

	var songDetails entity.SongDetails

	err = json.NewDecoder(resp.Body).Decode(&songDetails)
	if err != nil {
		return entity.SongDetails{}, fmt.Errorf("songClient: failed to decode response: %w", err)
	}

	return songDetails, nil
}
