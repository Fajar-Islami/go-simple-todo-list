package dto

import (
	"fmt"
	"time"
)

type ActivitiesFilter struct {
	Short string `query:"short_by" enums:"title-asc,title-desc,date-asc,date-desc"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

type ActivitiesReqCreate struct {
	Title string `json:"title" validate:"required" error:"title is required"`
	Email string `json:"email" validate:"email,required" error:"email is required"`
}

type ActivitiesReqUpdate struct {
	Title string `json:"title,omitempty" validate:"omitempty,required"`
	Email string `json:"email,omitempty" validate:"omitempty,email,required"`
}

func (a ActivitiesReqUpdate) Validate() error {
	if a.Title == "" && a.Email == "" {
		return fmt.Errorf("title cannot be null")
	}
	return nil
}

type ActivitiesResp struct {
	ActivityID int64     `json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
