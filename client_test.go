package hostio

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (*Client, *http.ServeMux) {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client := NewClient("secret")
	client.baseURL, _ = url.Parse(server.URL)
	client.HTTPClient = server.Client()

	return client, mux
}

func testHandler(filename string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(rw, fmt.Sprintf("unsupported method: %s", req.Method), http.StatusMethodNotAllowed)
			return
		}

		file, err := os.Open(filepath.Join("fixtures", filename))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() { _ = file.Close() }()

		_, err = io.Copy(rw, file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TestClient_Web(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/web/example.com", testHandler("web.json"))

	response, err := client.Web(context.Background(), "example.com")
	require.NoError(t, err)

	expected := &WebResponse{
		Domain:      "facebook.com",
		Rank:        2,
		URL:         "https://www.facebook.com/",
		IP:          "157.240.11.35",
		Date:        time.Date(2020, time.August, 26, 17, 39, 17, 981000000, time.UTC),
		Length:      160817,
		Encoding:    "utf8",
		Copyright:   "Facebook © 2020",
		Title:       "Facebook - Log In or Sign Up",
		Description: "Create an account or log into Facebook. Connect with friends, family and other people you know. Share photos and videos, send messages and get updates.",
		Links:       []string{"messenger.com", "oculus.com"},
	}

	assert.Equal(t, expected, response)
}

func TestClient_DNS(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/dns/example.com", testHandler("dns.json"))

	response, err := client.DNS(context.Background(), "example.com")
	require.NoError(t, err)

	expected := &DNSResponse{
		Domain: "facebook.com",
		A:      []string{"157.240.3.35"},
		AAAA:   []string{"2a03:2880:f101:83:face:b00c:0:25de"},
		MX:     []string{"10 smtpin.vvv.facebook.com."},
		NS:     []string{"a.ns.facebook.com.", "b.ns.facebook.com."},
	}

	assert.Equal(t, expected, response)
}

func TestClient_Related(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/related/example.com", testHandler("related.json"))

	response, err := client.Related(context.Background(), "example.com")
	require.NoError(t, err)

	expected := &RelatedResponse{
		IP:        []RelatedDomain{{Value: "172.217.14.238", Count: 293}, {Value: "216.58.193.68", Count: 71}},
		Redirects: []RelatedDomain{{Value: "google.com", Count: 629989}},
		ASN:       []RelatedDomain{{Value: "AS15169", Count: 16219992}},
		Backlinks: []RelatedDomain{{Value: "google.com", Count: 17314912}},
	}

	assert.Equal(t, expected, response)
}

func TestClient_Full(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/full/example.com", testHandler("full.json"))

	response, err := client.Full(context.Background(), "example.com")
	require.NoError(t, err)

	expected := &FullResponse{
		Domain: "google.com",
		DNS: DNSResponse{
			Domain: "google.com",
			A:      []string{"172.217.14.238"},
			AAAA:   []string{"2607:f8b0:400a:803::200e"},
			MX:     []string{"10 aspmx.l.google.com.", "20 alt1.aspmx.l.google.com.", "30 alt2.aspmx.l.google.com.", "40 alt3.aspmx.l.google.com.", "50 alt4.aspmx.l.google.com."},
			NS:     []string{"ns1.google.com.", "ns2.google.com.", "ns3.google.com.", "ns4.google.com."},
		},
		IPInfo: map[string]IPInfo{
			"216.58.193.68": {
				City:     "Seattle",
				Region:   "Washington",
				Country:  "US",
				Loc:      "47.6062,-122.3321",
				Postal:   "98111",
				Timezone: "America/Los_Angeles",
				ASN:      ASNInfo{ASN: "AS15169", Name: "Google LLC", Domain: "google.com", Route: "216.58.192.0/22", Type: "business"},
			},
			"2607:f8b0:400a:803::200e": {
				City:     "Mountain View",
				Region:   "California",
				Country:  "US",
				Loc:      "37.4056,-122.0775",
				Postal:   "94043",
				Timezone: "America/Los_Angeles",
				ASN:      ASNInfo{ASN: "AS15169", Name: "Google LLC", Domain: "google.com", Route: "2607:f8b0:400a::/48", Type: "business"},
			},
		},
		Web: WebResponse{
			Domain:      "google.com",
			Rank:        1,
			URL:         "https://www.google.com/?gws_rd=ssl",
			IP:          "216.58.193.68",
			Date:        time.Date(2019, time.November, 25, 18, 58, 31, 543000000, time.UTC),
			Length:      205694,
			Encoding:    "utf8",
			Copyright:   "",
			Title:       "Google",
			Description: "Search the world's information, including webpages, images, videos and more. Google has many special features to help you find exactly what you're looking for.",
			Links:       []string{"about.google"},
		},
		Related: RelatedResponse{
			IP:        []RelatedDomain{{Value: "172.217.14.238", Count: 293}, {Value: "216.58.193.68", Count: 71}},
			Redirects: []RelatedDomain{{Value: "google.com", Count: 629989}},
			ASN:       []RelatedDomain{{Value: "AS15169", Count: 16219992}},
			Backlinks: []RelatedDomain{{Value: "google.com", Count: 17314912}},
			MX:        []RelatedDomain{{Value: "google.com", Count: 11624298}},
			NS:        []RelatedDomain{{Value: "google.com", Count: 12221}},
		},
	}

	assert.Equal(t, expected, response)
}

func TestClient_Domains(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/domains/ns/google.com", testHandler("domains.json"))

	pager := &Pager{
		Limit: 5,
		Page:  5,
	}

	response, err := client.Domains(context.Background(), NS, "google.com", pager)
	require.NoError(t, err)

	expected := &DomainsResponse{
		NS: "google.com",
		Domains: []string{
			"google.com.eg",
			"google.co.th",
			"google.nl",
			"google.co.ve",
			"google.co.za",
		},
		Page:  5,
		Total: 12221,
	}

	assert.Equal(t, expected, response)
}
