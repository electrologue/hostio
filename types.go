package hostio

import "time"

type Field string

const (
	// IP address (v4 or v6).
	// Example: `8.8.8.8`.
	IP Field = "ip"

	// NS DNS Nameserver record (root domain).
	// Example: `google.com`.
	NS Field = "ns"

	// MX DNS Mailserver record (root domain).
	// Example: `google.com`.
	MX Field = "mx"

	// ASN AS number.
	// Example: `AS15169`.
	ASN Field = "asn"

	// Backlinks Domains that include a link to the domain on their homepage.
	// Example: `google.com`.
	Backlinks Field = "backlinks"

	// Redirects Domains that redirect to the domain from their homepage.
	// Example: `google.com`.
	Redirects Field = "redirects"

	// Adsense Domains that include an adsense ID on their homepage.
	// Example: `pub-1556223355139109`.
	Adsense Field = "adsense"

	// Facebook Domains linking to the Facebook social media handle on their homepage.
	// Example: `spacenewsx`.
	Facebook Field = "facebook"

	// Twitter Domains linking to the Twitter social media handle on their homepage.
	// Example: `elonmusk`.
	Twitter Field = "twitter"

	// Instagram Domains linking to the Instagram social media handle on their homepage.
	// Example: `chadwickboseman`.
	Instagram Field = "instagram"

	// GTM Domains having the google tag manager ID on their homepage.
	// Example: `GTM-544JFM`.
	GTM Field = "gtm"

	// GoogleAnalytics Domains that include a googleanalytics ID on their homepage.
	// Example: `UA-55552418`.
	GoogleAnalytics Field = "googleanalytics"

	// Email Domains that include an email address on their homepage.
	// Example: `admin@google.com`.
	Email Field = "email"
)

type WebResponse struct {
	Domain string `json:"domain,omitempty"`

	// Position in host.io 10M domains ranking, https://host.io/rankings
	Rank int `json:"rank,omitempty"`

	// URL we scraped the data from
	URL string `json:"url,omitempty"`

	// Actual IP we scraped the data from
	IP string `json:"ip,omitempty"`

	// Date we scraped the data
	Date time.Time `json:"date,omitempty"`

	// Length of the HTML content we scraped
	Length int `json:"length,omitempty"`

	// Encoding of the scraped data
	Encoding string `json:"encoding,omitempty"`

	// Scraped copyright notice
	Copyright string `json:"copyright,omitempty"`

	// HTML title
	Title string `json:"title,omitempty"`

	// HTML meta description
	Description string `json:"description,omitempty"`

	// Domains of links on the homepage
	Links []string `json:"links,omitempty"`
}

type DNSResponse struct {
	Domain string   `json:"domain,omitempty"`
	A      []string `json:"a,omitempty"`
	AAAA   []string `json:"aaaa,omitempty"`
	MX     []string `json:"mx,omitempty"`
	NS     []string `json:"ns,omitempty"`
}

type RelatedResponse struct {
	IP        []RelatedDomain `json:"ip,omitempty"`
	Redirects []RelatedDomain `json:"redirects,omitempty"`
	ASN       []RelatedDomain `json:"asn,omitempty"`
	Backlinks []RelatedDomain `json:"backlinks,omitempty"`
	MX        []RelatedDomain `json:"mx,omitempty"`
	NS        []RelatedDomain `json:"ns,omitempty"`
}

type RelatedDomain struct {
	Value string `json:"value,omitempty"`
	Count int    `json:"count,omitempty"`
}

type IPInfo struct {
	City     string  `json:"city,omitempty"`
	Region   string  `json:"region,omitempty"`
	Country  string  `json:"country,omitempty"`
	Loc      string  `json:"loc,omitempty"`
	Postal   string  `json:"postal,omitempty"`
	Timezone string  `json:"timezone,omitempty"`
	ASN      ASNInfo `json:"asn,omitempty"`
}

type ASNInfo struct {
	ASN    string `json:"asn,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Route  string `json:"route,omitempty"`
	Type   string `json:"type,omitempty"`
}

type FullResponse struct {
	Domain  string            `json:"domain,omitempty"`
	DNS     DNSResponse       `json:"dns,omitempty"`
	IPInfo  map[string]IPInfo `json:"ipinfo,omitempty"`
	Web     WebResponse       `json:"web,omitempty"`
	Related RelatedResponse   `json:"related,omitempty"`
}

type DomainsResponse struct {
	// IP address (v4 or v6).
	IP string `json:"ip,omitempty"`

	// NS DNS Nameserver (NS) record (root domain).
	NS string `json:"ns,omitempty"`

	// MX DNS Mailserver (MX) record (root domain).
	MX string `json:"mx,omitempty"`

	// ASN AS Number.
	ASN string `json:"asn,omitempty"`

	// Backlinks domains that include a link to the domain on their homepage.
	Backlinks []string `json:"backlinks,omitempty"`

	// Redirects domains that redirect to the domain from their homepage.
	Redirects []string `json:"redirects,omitempty"`

	// Adsense domains that include an adsense ID on their homepage.
	Adsense []string `json:"adsense,omitempty"`

	// Facebook domains linking to the Facebook social media handle on their homepage.
	Facebook string `json:"facebook,omitempty"`

	// Twitter domains linking to the Twitter social media handle on their homepage.
	Twitter string `json:"twitter,omitempty"`

	// Instagram domains linking to the Instagram social media handle on their homepage.
	Instagram string `json:"instagram,omitempty"`

	// GTM domains having the google tag manager ID on their homepage.
	GTM string `json:"gtm,omitempty"`

	// GoogleAnalytics domains that include a googleanalytics ID on their homepage.
	GoogleAnalytics string `json:"googleanalytics,omitempty"`

	// Email domains that include an email address on their homepage.
	Email string `json:"email,omitempty"`

	Domains []string `json:"domains,omitempty"`

	Page  int `json:"page"`
	Total int `json:"total"`
}
