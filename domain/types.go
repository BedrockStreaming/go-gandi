package domain

import (
	"time"

	"github.com/go-gandi/go-gandi/internal/client"
)

// Domain is the API client to the Gandi v5 Domain API
type Domain struct {
	client client.Gandi
}

// Contact represents a contact associated with a domain
type Contact struct {
	Country         string                 `json:"country"`
	Email           string                 `json:"email"`
	FamilyName      string                 `json:"family"`
	GivenName       string                 `json:"given"`
	StreetAddr      string                 `json:"streetaddr"`
	ContactType     int                    `json:"type"`
	BrandNumber     string                 `json:"brand_number,omitempty"`
	City            string                 `json:"city,omitempty"`
	DataObfuscated  *bool                  `json:"data_obfuscated,omitempty"`
	Fax             string                 `json:"fax,omitempty"`
	Language        string                 `json:"lang,omitempty"`
	MailObfuscated  *bool                  `json:"mail_obfuscated,omitempty"`
	Mobile          string                 `json:"mobile,omitempty"`
	OrgName         string                 `json:"orgname,omitempty"`
	Phone           string                 `json:"phone,omitempty"`
	Siren           string                 `json:"siren,omitempty"`
	State           string                 `json:"state,omitempty"`
	Validation      string                 `json:"validation,omitempty"`
	Zip             string                 `json:"zip,omitempty"`
	ExtraParameters map[string]interface{} `json:"extra_parameters,omitempty"`
}

// ResponseDates represents all the dates associated with a domain
type ResponseDates struct {
	RegistryCreatedAt   *time.Time `json:"registry_created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
	AuthInfoExpiresAt   *time.Time `json:"authinfo_expires_at,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	DeletesAt           *time.Time `json:"deletes_at,omitempty"`
	HoldBeginsAt        *time.Time `json:"hold_begins_at,omitempty"`
	HoldEndsAt          *time.Time `json:"hold_ends_at,omitempty"`
	PendingDeleteEndsAt *time.Time `json:"pending_delete_ends_at,omitempty"`
	RegistryEndsAt      *time.Time `json:"registry_ends_at,omitempty"`
	RenewBeginsAt       *time.Time `json:"renew_begins_at,omitempty"`
	RenewEndsAt         *time.Time `json:"renew_ends_at,omitempty"`
}

// NameServerConfig represents the name server configuration for a domain
type NameServerConfig struct {
	Current string   `json:"current"`
	Hosts   []string `json:"hosts,omitempty"`
}

// ListResponse is the response object returned by listing domains
type ListResponse struct {
	AutoRenew   *bool             `json:"autorenew"`
	Dates       *ResponseDates    `json:"dates"`
	DomainOwner string            `json:"domain_owner"`
	FQDN        string            `json:"fqdn"`
	FQDNUnicode string            `json:"fqdn_unicode"`
	Href        string            `json:"href"`
	ID          string            `json:"id"`
	NameServer  *NameServerConfig `json:"nameserver"`
	OrgaOwner   string            `json:"orga_owner"`
	Owner       string            `json:"owner"`
	Status      []string          `json:"status"`
	TLD         string            `json:"tld"`
	SharingID   string            `json:"sharing_id,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
}

// AutoRenew is the auto renewal information for the domain
type AutoRenew struct {
	Href     string       `json:"href,omitempty"`
	Dates    *[]time.Time `json:"dates,omitempty"`
	Duration int          `json:"duration,omitempty"`
	Enabled  *bool        `json:"enabled,omitempty"`
	OrgID    string       `json:"org_id,omitempty"`
}

// SharingSpace defines the Organisation that owns the domain
type SharingSpace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Details describes a single domain
type Details struct {
	AutoRenew    *AutoRenew     `json:"autorenew"`
	CanTLDLock   *bool          `json:"can_tld_lock"`
	Contacts     *Contacts      `json:"contacts"`
	Dates        *ResponseDates `json:"dates"`
	FQDN         string         `json:"fqdn"`
	FQDNUnicode  string         `json:"fqdn_unicode"`
	Href         string         `json:"href"`
	Nameservers  []string       `json:"nameservers,omitempty"`
	Services     []string       `json:"services"`
	SharingSpace *SharingSpace  `json:"sharing_space"`
	Status       []string       `json:"status"`
	TLD          string         `json:"tld"`
	AuthInfo     string         `json:"authinfo,omitempty"`
	ID           string         `json:"id,omitempty"`
	SharingID    string         `json:"sharing_id,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	TrusteeRoles []string       `json:"trustee_roles,omitempty"`
}

// CreateRequest is used to request a new domain
type CreateRequest struct {
	FQDN     string   `json:"fqdn"`
	Owner    *Contact `json:"owner"`
	Admin    *Contact `json:"admin,omitempty"`
	Billing  *Contact `json:"bill,omitempty"`
	Claims   string   `json:"claims,omitempty"`
	Currency string   `json:"currency,omitempty"`
	// Duration in years between 1 and 10
	Duration       int    `json:"duration,omitempty"`
	EnforcePremium bool   `json:"enforce_premium,omitempty"`
	Lang           string `json:"lang,omitempty"`
	// NameserverIPs sets the Glue Records for the domain
	NameserverIPs map[string]string `json:"nameserver_ips,omitempty"`
	Nameservers   []string          `json:"nameservers,omitempty"`
	Price         int               `json:"price,omitempty"`
	ReselleeID    string            `json:"resellee_id,omitempty"`
	// SMD is a Signed Mark Data file; if used, `TLDPeriod` must be "sunrise"
	SMD       string   `json:"smd,omitempty"`
	Tech      *Contact `json:"tech,omitempty"`
	TLDPeriod string   `json:"tld_period,omitempty"`
}

// Contacts is the set of contacts associated with a Domain
type Contacts struct {
	Admin   *Contact `json:"admin,omitempty"`
	Billing *Contact `json:"bill,omitempty"`
	Owner   *Contact `json:"owner,omitempty"`
	Tech    *Contact `json:"tech,omitempty"`
}

// Nameservers represents a list of nameservers
type Nameservers struct {
	Nameservers []string `json:"nameservers,omitempty"`
}

// DNSSECKey represents a DNSSEC key associated with a domain
type DNSSECKey struct {
	Algorithm  int    `json:"algorithm"`
	Digest     string `json:"digest"`
	DigestType int    `json:"digest_type"`
	ID         int    `json:"id"`
	KeyTag     int    `json:"keytag"`
	Type       string `json:"type"`
	PublicKey  string `json:"public_key"`
}

// DNSSECKeyCreateRequest represents a request to create a DNSSEC key for a domain
type DNSSECKeyCreateRequest struct {
	Algorithm int    `json:"algorithm"`
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
}
