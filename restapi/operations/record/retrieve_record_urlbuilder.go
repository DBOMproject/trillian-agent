// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RetrieveRecordURL generates an URL for the retrieve record operation
type RetrieveRecordURL struct {
	ChannelID string
	RecordID  string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveRecordURL) WithBasePath(bp string) *RetrieveRecordURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveRecordURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RetrieveRecordURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/channels/{channelID}/records/{recordID}"

	channelID := o.ChannelID
	if channelID != "" {
		_path = strings.Replace(_path, "{channelID}", channelID, -1)
	} else {
		return nil, errors.New("channelId is required on RetrieveRecordURL")
	}

	recordID := o.RecordID
	if recordID != "" {
		_path = strings.Replace(_path, "{recordID}", recordID, -1)
	} else {
		return nil, errors.New("recordId is required on RetrieveRecordURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RetrieveRecordURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RetrieveRecordURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RetrieveRecordURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RetrieveRecordURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RetrieveRecordURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *RetrieveRecordURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
