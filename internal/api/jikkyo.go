package api

import (
	"encoding/xml"
	"fmt"

	"github.com/hogecode/JikkyoUtil/internal/config"
	"github.com/hogecode/JikkyoUtil/internal/models"
)

// GetJikkyoComments fetches comments from Jikkyo API
func (c *Client) GetJikkyoComments(jikkyoID string, startTime, endTime int64) (*models.JikkyoResponse, error) {
	url := fmt.Sprintf("%s/%s", config.JikkyoBaseURL, jikkyoID)

	resp, err := c.R().
		SetQueryParams(map[string]string{
			"starttime": fmt.Sprintf("%d", startTime),
			"endtime":   fmt.Sprintf("%d", endTime),
			"format":    "json",
		}).
		SetResult(&models.JikkyoResponse{}).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to call Jikkyo API: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("Jikkyo API returned status %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*models.JikkyoResponse)
	if !ok {
		return nil, fmt.Errorf("failed to parse Jikkyo response")
	}

	return result, nil
}

// GetJikkyoCommentsXML fetches comments from Jikkyo API in XML format
func (c *Client) GetJikkyoCommentsXML(jikkyoID string, startTime, endTime int64) (*models.Packet, error) {
	url := fmt.Sprintf("%s/%s", config.JikkyoBaseURL, jikkyoID)

	resp, err := c.R().
		SetQueryParams(map[string]string{
			"starttime": fmt.Sprintf("%d", startTime),
			"endtime":   fmt.Sprintf("%d", endTime),
			"format":    "xml",
		}).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to call Jikkyo API: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("Jikkyo API returned status %d", resp.StatusCode())
	}

	var result models.Packet
	err = xml.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Jikkyo XML response: %w", err)
	}

	return &result, nil
}
