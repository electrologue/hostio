// Package hostio contains a host.io API client.
package hostio

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

const defaultBaseURL = "https://host.io/api"

type Pager struct {
	Limit int
	Page  int
}

// Client is a host.io API client.
type Client struct {
	token   string
	baseURL *url.URL

	HTTPClient *http.Client
}

// NewClient creates a new Client.
func NewClient(token string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		token:      token,
		baseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
}

// Web Metadata scraped from a domain homepage.
// https://host.io/docs#apiwebdomain
func (c Client) Web(ctx context.Context, domain string) (*WebResponse, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "web", domain))
	if err != nil {
		return nil, err
	}

	var apiResp WebResponse
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// DNS Get all the DNS records stored for a domain.
// https://host.io/docs#apidnsdomain
func (c Client) DNS(ctx context.Context, domain string) (*DNSResponse, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "dns", domain))
	if err != nil {
		return nil, err
	}

	var apiResp DNSResponse
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// Related Get a count of the number of related domains for all supported lookups we offer.
// https://host.io/docs#apirelateddomain
func (c Client) Related(ctx context.Context, domain string) (*RelatedResponse, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "related", domain))
	if err != nil {
		return nil, err
	}

	var apiResp RelatedResponse
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// Full A single endpoint that includes the data from /api/web, /api/dns, /api/related and IPinfo.
// https://host.io/docs#apifulldomain
func (c Client) Full(ctx context.Context, domain string) (*FullResponse, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "full", domain))
	if err != nil {
		return nil, err
	}

	var apiResp FullResponse
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

// Domains Get all domains associated with Field, and a count of the total.
// https://host.io/docs#apidomainsfieldvalue
func (c Client) Domains(ctx context.Context, field Field, value string, pager *Pager) (*DomainsResponse, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "domains", string(field), value))
	if err != nil {
		return nil, err
	}

	if pager != nil {
		query := endpoint.Query()
		query.Set("limit", strconv.Itoa(pager.Limit))
		query.Set("page", strconv.Itoa(pager.Page))
		endpoint.RawQuery = query.Encode()
	}

	var apiResp DomainsResponse
	err = c.do(ctx, endpoint, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

func (c Client) do(ctx context.Context, endpoint *url.URL, data interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(data))
	}

	return json.NewDecoder(resp.Body).Decode(data)
}
