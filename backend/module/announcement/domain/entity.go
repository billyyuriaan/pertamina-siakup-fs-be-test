package domain

import "time"

// Announcement represents a broadcast message shown to students, lecturers, or everyone.
type Announcement struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	TargetRole  string     `json:"target_role"` // "student", "lecturer", or "all"
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Valid target_role values. Anything outside this set must be rejected at the usecase layer.
var ValidTargetRoles = map[string]bool{
	"student":  true,
	"lecturer": true,
	"all":      true,
}
