package model

import "time"

type Campaign struct {
	ID               int
	UserID           string
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdateAt         time.Time
	CampaignImages   []CampaignImage
	User             User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdateAt   time.Time
}
