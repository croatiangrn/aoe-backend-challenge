package model

import (
	"fmt"
)

const (
	SuperPowerStrength        = "strength"
	SuperPowerSpeed           = "speed"
	SuperPowerFlight          = "flight"
	SuperPowerInvulnerability = "invulnerability"
	SuperPowerHealing         = "healing"
)

var (
	allowedSuperPowersMap = map[string]struct{}{SuperPowerStrength: {}, SuperPowerSpeed: {}, SuperPowerFlight: {}, SuperPowerInvulnerability: {}, SuperPowerHealing: {}}
)

type SuperPower string

type SuperHero struct {
	Name        string            `json:"name"`
	Identity    SuperHeroIdentity `json:"identity"`
	Birthday    string            `json:"birthday"`
	Superpowers []string          `json:"superpowers"`
}

func (i *SuperHero) Validate() error {
	for j := range i.Superpowers {
		if _, ok := allowedSuperPowersMap[i.Superpowers[j]]; !ok {
			return fmt.Errorf("superpower %q is not allowed", i.Superpowers[j])
		}
	}
	
	return nil
}

type SuperHeroIdentity struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
