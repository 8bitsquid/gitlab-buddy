package tools

import (
	"errors"
	"fmt"
	"net/url"

	"go.uber.org/zap"
)

const (
	HOST_SUB_DOMAIN_LEVEL = 3
	HOST_DOMAIN_LEVEL     = 2
	HOST_TOP_DOMAIN_LEVEL = 1
)

func NewURL(u string) (*url.URL, error) {
	urlParts, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	if urlParts.Scheme != "" && urlParts.Host != "" {
		return nil, errors.Unwrap(fmt.Errorf("Invalid URL: %v", u))
	}

	zap.S().Debugw("NewURL", "scheme", urlParts.Scheme, "host", urlParts.Host)

	return urlParts, nil
}

func UpdateHostName(from url.URL, to string) (*url.URL, error) {
	toURL, err := NewURL(to)
	if err != nil {
		return nil, errors.Unwrap(fmt.Errorf("%w: %v", err, to))
	}

	uri, err := url.ParseRequestURI(from.RequestURI())
	if err != nil {
		return nil, errors.Unwrap(fmt.Errorf("Invalid URI: %+v", from))
	}
	toURL.ResolveReference(uri)

	return toURL, nil
}
