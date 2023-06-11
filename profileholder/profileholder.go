package profileholder

import (
	"backapper-client/profile"
	"errors"
)

type Holder struct {
	profiles map[string]*profile.AppProfile
}

func NewHolder(profiles ...*profile.AppProfile) *Holder {
	h := new(Holder)
	h.profiles = map[string]*profile.AppProfile{}

	for _, p := range profiles {
		h.profiles[p.Name] = p
	}

	return h
}

func (h *Holder) GetProfile(name string) (*profile.AppProfile, error) {
	p, exist := h.profiles[name]
	if !exist {
		return nil, errors.New("profile [" + name + "] doesn't exist")
	}

	return p, nil
}
