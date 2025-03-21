package types

import (
	"time"
)

//making the fields unexported so that ItemMetadata cannot be instantiated, almost acting like an abstract type
type ItemMetadata struct {
	id string `json:"id"`
	createddatetime time.Time `json:"createddatetime"`
	title *string `json:"title,omitempty"`  //Optional, hence the *
}

func NewItemMetadata(id string, createdDateTime time.Time, title *string) *ItemMetadata {
	return &ItemMetadata{
		id:              id,
		createdDateTime: createdDateTime,
		title:           title,
	}
}

type WorkoutComponent struct {
	ItemMetadata //embedding item metadata here, which essentially acts like inheritance
	Part string `json:"part"`
	Tags []string `json:"tags"`
	Category string `json:"category"`
	Type ComponentType `json:"type"`

	//the following are optional components
	Name *string `json:"name,omitempty"` 
	CaloriesBurnt *uint32 `json:"calburnt,omitempty"`
	Source *string `json:"source,omitempty"` //a link/youtube link to the workout
	Notes []string `json:"notes,omitempty"`
}

type RepComponent struct {
	ItemMetadata
	Reps uint8 `json:"reps"`
	Sets uint8 `json:"sets"`
	RestDuration string `json:"restduration"`
}

type NonRepComponent struct {
	Duration uint64 `json:"duration"`
}

type ComponentType interface {
	IsComponent() //any type that implements this method will be accepted
}

func (c RepComponent) IsComponent() {}

func (c NonRepComponent) IsComponent() {}


