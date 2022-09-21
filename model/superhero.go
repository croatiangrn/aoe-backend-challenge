package model

import (
	"encoding/json"
	"fmt"
	"github.com/croatiangrn/aoe-backend-challenge/app"
	"github.com/pkg/errors"
	"io/ioutil"
	"time"
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

type SuperHero struct {
	Name        string            `json:"name"`
	Identity    SuperHeroIdentity `json:"identity"`
	Birthday    time.Time         `json:"birthday"`
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

func (i *SuperHero) Save() error {
	jsonFileBytes, err := ioutil.ReadFile(app.Config.SuperHeroesJSONPath())
	if err != nil {
		return errors.WithStack(err)
	}

	superHeroes := make([]SuperHero, 0)
	if err := json.Unmarshal(jsonFileBytes, &superHeroes); err != nil {
		return errors.WithStack(err)
	}

	superHeroes = append(superHeroes, *i)

	superHeroesBytes, err := json.Marshal(superHeroes)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := ioutil.WriteFile(app.Config.SuperHeroesJSONPath(), superHeroesBytes, 0644); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type SuperHeroIdentity struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
