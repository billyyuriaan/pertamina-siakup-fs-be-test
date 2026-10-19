package domain

// Pagination mirrors the pattern used across SIAKUP services: page is 1-indexed,
// size <= 0 means "no pagination requested" and must be handled explicitly.
type Pagination struct {
	Page int
	Size int
}

func (p Pagination) GetLimit() int {
	if p.Size <= 0 {
		return 0
	}
	return p.Size
}

func (p Pagination) GetOffset() int {
	if p.Size <= 0 || p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

// AnnouncementFilter is passed from the handler down to the repository.
type AnnouncementFilter struct {
	Pagination
	Search     string // case-insensitive match against Title
	TargetRole string // exact match against TargetRole; empty = no filter
}

// CreateAnnouncementRequest is the expected JSON body for POST /announcements.
type CreateAnnouncementRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	TargetRole string `json:"target_role"`
}
