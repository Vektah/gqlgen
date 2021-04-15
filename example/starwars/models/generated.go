// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Character interface {
	IsCharacter()
}

type SearchResult interface {
	IsSearchResult()
}

type FriendsEdge struct {
	Cursor string    `json:"cursor"`
	Node   Character `json:"node,omitempty"`
}

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type Starship struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Length  float64 `json:"length"`
	History [][]int `json:"history"`
}

func (Starship) IsSearchResult() {}

type Episode string

const (
	EpisodeNewhope Episode = "NEWHOPE"
	EpisodeEmpire  Episode = "EMPIRE"
	EpisodeJedi    Episode = "JEDI"
)

var AllEpisode = []Episode{
	EpisodeNewhope,
	EpisodeEmpire,
	EpisodeJedi,
}

func (e Episode) IsValid() bool {
	switch e {
	case EpisodeNewhope, EpisodeEmpire, EpisodeJedi:
		return true
	}
	return false
}

func (e Episode) String() string {
	return string(e)
}

func (e *Episode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Episode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Episode", str)
	}
	return nil
}

func (e Episode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LengthUnit string

const (
	LengthUnitMeter LengthUnit = "METER"
	LengthUnitFoot  LengthUnit = "FOOT"
)

var AllLengthUnit = []LengthUnit{
	LengthUnitMeter,
	LengthUnitFoot,
}

func (e LengthUnit) IsValid() bool {
	switch e {
	case LengthUnitMeter, LengthUnitFoot:
		return true
	}
	return false
}

func (e LengthUnit) String() string {
	return string(e)
}

func (e *LengthUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LengthUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LengthUnit", str)
	}
	return nil
}

func (e LengthUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
