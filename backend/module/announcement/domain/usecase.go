package domain

import "context"

// AnnouncementUseCase — implement this in usecase/usecase.go. Do not change
// this interface.
type AnnouncementUseCase interface {
	// GetList returns the page of announcements, the total record count, and the
	// total number of pages given filter.Size (0 total pages if filter.Size <= 0).
	GetList(ctx context.Context, filter AnnouncementFilter) (announcements []*Announcement, total int64, totalPages int, err error)

	// Create validates req and persists a new Announcement. It must reject:
	//   - empty/whitespace-only Title, or Title longer than 200 characters
	//   - TargetRole not in ValidTargetRoles
	Create(ctx context.Context, req CreateAnnouncementRequest) (*Announcement, error)
}
