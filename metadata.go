package go_certcentral

import (
	"encoding/json"
	"io"
	"net/http"
)

const metadataURL = "/account/metadata"

type Metadata struct {
	MetadataList []MetadataItem `json:"metadata,omitempty"`
}

type MetadataItem struct {
	ID            int    `json:"id"`
	Label         string `json:"label,omitempty"`
	IsRequired    bool   `json:"is_required,omitempty"`
	IsActive      bool   `json:"is_active,omitempty"`
	DataType      string `json:"data_type,omitempty"`
	ShowInReceipt bool   `json:"show_in_receipt,omitempty"`
}

func (c *Client) QueryMetadata() (*Metadata, error) {
	req, err := http.NewRequest(http.MethodGet, makeURL(metadataURL), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	err = json.Unmarshal(body, &metadata)
	return &metadata, err
}
