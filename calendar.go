package main

import "golang.org/x/oauth2"

type calendar struct {
	ID          string         `json:"id"` // FIXME: @awwalker Can this be combined with Alias?
	ClientID    string         `json:"client_id"`
	Alias       string         `json:"alias"`
	Token       *oauth2.Token  `json:"oauth_token"`
	OAuthConfig *oauth2.Config `json:"oauth_config"`
}
