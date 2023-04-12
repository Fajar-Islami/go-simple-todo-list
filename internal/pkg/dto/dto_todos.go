package dto

import (
	"fmt"
	"time"
)

type TodosFilter struct {
	Short           string `query:"short_by" enums:"title-asc,title-desc,date-asc,date-desc,done"`
	Limit           int    `query:"limit"`
	Page            int    `query:"page"`
	ActivityGroupID int    `query:"activity_group_id"`
}

type TodosReqCreate struct {
	ActivityGroupID int    `json:"activity_group_id" validate:"required,numeric" error:"activity_group_id is required"`
	Priority        string `json:"priority" error:"priority is required" enum:"1,2,3,4,5"`
	Title           string `json:"title" validate:"required" error:"title is required" enum:"very-high,high,medium,low,very-low"`
}

func (a TodosReqCreate) Validate() error {
	if a.Title == "" && a.ActivityGroupID == 0 && a.Priority == "" {
		return fmt.Errorf("title cannot be null")
	}
	return nil
}

type TodosReqUpdate struct {
	ActivityGroupID int    `json:"activity_group_id,omitempty"`
	Priority        string `json:"priority,omitempty"`
	Title           string `json:"title,omitempty"`
	IsActive        bool   `json:"is_active,omitempty"`
}

type TodosResp struct {
	TodoID          int64     `json:"id"`
	ActivityGroupID int64     `json:"activity_group_id"`
	Title           string    `json:"title"`
	Priority        string    `json:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
