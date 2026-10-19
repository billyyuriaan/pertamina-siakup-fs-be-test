package domain

import "context"

// AnnouncementRepository is already fully implemented for you in
// repository/memory/repository.go — you should not need to touch it.
type AnnouncementRepository interface {
	FindAll(ctx context.Context, filter AnnouncementFilter) ([]*Announcement, int64, error)
	Store(ctx context.Context, announcement *Announcement) error
}
