package advertisments

import (
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Advertisment struct {
	ID        uuid.UUID `json:"ad_id"`
	CreatedAt time.Time `json:"-"`
	CreatedBy uuid.UUID `json:"-"`
	ExpiresAt time.Time `json:"-"`
	Tags      []string  `json:"tags"`
}

func NewAdvertisment(tags []string) *Advertisment {
	return &Advertisment{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Tags:      tags,
	}
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
	ad := NewAdvertisment(tags)
	id := ad.ID
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
