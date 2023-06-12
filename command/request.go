package command

import (
	"backapper-client/params"
	"backapper-client/profile"
	"net/http"
)

type Request struct {
	Params  *params.Params
	Profile *profile.AppProfile
	Client  *http.Client
}

func NewRequest(params *params.Params, profile *profile.AppProfile, client *http.Client) *Request {
	return &Request{Params: params, Profile: profile, Client: client}
}
