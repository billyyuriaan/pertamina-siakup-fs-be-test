package memory

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"siakup-fs-test-starter/module/announcement/domain"
)

// Repository is a thread-safe in-memory store. It is already fully
// implemented — you do not need to change this file.
type Repository struct {
	mu     sync.Mutex
	data   map[string]*domain.Announcement
	nextID int
}

func NewRepository() *Repository {
	return &Repository{
		data: make(map[string]*domain.Announcement),
	}
}

func (r *Repository) FindAll(ctx context.Context, filter domain.AnnouncementFilter) ([]*domain.Announcement, int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	matched := make([]*domain.Announcement, 0, len(r.data))
	for _, a := range r.data {
		if filter.TargetRole != "" && a.TargetRole != filter.TargetRole {
			continue
		}
		if filter.Search != "" && !strings.Contains(strings.ToLower(a.Title), strings.ToLower(filter.Search)) {
			continue
		}
		matched = append(matched, a)
	}

	// newest first
	for i := 0; i < len(matched); i++ {
		for j := i + 1; j < len(matched); j++ {
			if matched[j].CreatedAt.After(matched[i].CreatedAt) {
				matched[i], matched[j] = matched[j], matched[i]
			}
		}
	}

	total := int64(len(matched))

	limit := filter.GetLimit()
	if limit <= 0 {
		return matched, total, nil
	}
	offset := filter.GetOffset()
	if offset >= len(matched) {
		return []*domain.Announcement{}, total, nil
	}
	end := offset + limit
	if end > len(matched) {
		end = len(matched)
	}
	return matched[offset:end], total, nil
}

func (r *Repository) Store(ctx context.Context, announcement *domain.Announcement) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.nextID++
	announcement.Id = strconv.Itoa(r.nextID)
	now := time.Now()
	announcement.CreatedAt = now
	announcement.UpdatedAt = now

	r.data[announcement.Id] = announcement
	return nil
}
