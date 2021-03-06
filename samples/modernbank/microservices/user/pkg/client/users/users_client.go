// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateUser creates a new user identity for a customer

Creates a new user
*/
func (a *Client) CreateUser(params *CreateUserParams) (*CreateUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserCreated), nil

}

/*
GetUserByUserName gets user by user name
*/
func (a *Client) GetUserByUserName(params *GetUserByUserNameParams) (*GetUserByUserNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserByUserNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserByUserName",
		Method:             "GET",
		PathPattern:        "/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserByUserNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserByUserNameOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
