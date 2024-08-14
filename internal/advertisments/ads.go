package advertisments

import (
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Advertisment struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []string  `json:"tags"`
}

type AdvertismentInventory struct {
	inventory map[uuid.UUID]*Advertisment
}

func NewAdvertismentInventory() *AdvertismentInventory {
	return &AdvertismentInventory{
		inventory: make(map[uuid.UUID]*Advertisment),
	}
}

func (ai *AdvertismentInventory) CreateAdvertisment(tags []string) (*Advertisment, error) {
	id := uuid.New()
	ad := &Advertisment{
		ID:        id,
		CreatedAt: time.Now(),
		Tags:      tags,
	}

	if _, exists := ai.inventory[id]; !exists {
		ai.inventory[id] = ad

		return ad, nil
	}

	// ad already exists
	return nil, errors.New("Advertisment already exists.")
}

func (ai *AdvertismentInventory) GetAdvertismentByID(id uuid.UUID) (*Advertisment, error) {
	if ad, exists := ai.inventory[id]; exists {
		return ad, nil
	}

	// ad does not exist
	return nil, errors.New("Advertisment doesn't exist.")
}

func (ai *AdvertismentInventory) GetAdvertismentsWithTag(tag string) ([]*Advertisment, error) {
	var ads []*Advertisment

	for _, v := range ai.inventory {
		if slices.Contains(v.Tags, tag) {
			ads = append(ads, v)
		}
	}

	return ads, nil
}

func (ai *AdvertismentInventory) GetAllAdvertisments() ([]*Advertisment, error) {
	var ads []*Advertisment

	for _, v := range ai.inventory {
		ads = append(ads, v)
	}

	return ads, nil
}
