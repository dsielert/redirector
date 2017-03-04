package store

import (
	"fmt"
	"github.com/tboerger/redirects/model"
	// "golang.org/x/crypto/acme/autocert"
)

var (
	// ErrRedirectNotFound gets returned if a redirect can't be found on the store.
	ErrRedirectNotFound = fmt.Errorf("Failed to find a redirect")

	// ErrRedirectSourceExists gets returned if a redirect s source already exists.
	ErrRedirectSourceExists = fmt.Errorf("Source already exists")
)

// Store implements all required data-layer functions for Redirects.
type Store interface {
	// GetRedirects retrieves all redirects from the store.
	GetRedirects() ([]*model.Redirect, error)

	// GetRedirect retrieves a specific redirect from the store.
	GetRedirect(string) (*model.Redirect, error)

	// DeleteRedirect deletes a redirect from the store.
	DeleteRedirect(string) error

	// CreateRedirect creates a redirect on the store.
	CreateRedirect(*model.Redirect) error

	// UpdateRedirect updates a redirect on the store.
	UpdateRedirect(*model.Redirect) error

	// CertCache returns the cert cache implementation.
	// CertCache() autocert.Cache
}
