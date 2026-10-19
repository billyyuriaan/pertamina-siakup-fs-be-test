package usecase

import (
	"context"

	"siakup-fs-test-starter/module/announcement/domain"
	"siakup-fs-test-starter/module/announcement/repository"
)

type UseCase struct {
	repository domain.AnnouncementRepository
}

func NewUseCase(repository domain.AnnouncementRepository) *UseCase {
	return &UseCase{repository: repository}
}

// TODO(candidate): implement per the contract documented on
// domain.AnnouncementUseCase.GetList. Compute totalPages from total and
// filter.Size — see the pagination logic question in the written test if
// you're unsure how to handle filter.Size <= 0.
func (u *UseCase) GetList(ctx context.Context, filter domain.AnnouncementFilter) ([]*domain.Announcement, int64, int, error) {
	announcements, total, err := u.repository.FindAll(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := 1
	if filter.Size > 0 {
		totalPages = int((total + int64(filter.Size) - 1) / int64(filter.Size))
	}

	return announcements, total, totalPages, nil
}

// TODO(candidate): implement per the contract documented on
// domain.AnnouncementUseCase.Create. Validation belongs here, NOT in the HTTP
// handler and NOT in the repository.
func (u *UseCase) Create(ctx context.Context, req domain.CreateAnnouncementRequest) (*domain.Announcement, error) {
	if req.title == "" {
		return nil, fmt.Errorf("title is required")
	}

	if req.content == "" {
		return nil, fmt.Errorf("content is required")
	}

	if len(req.content) < 200 {
		return nil, fmt.Errorf("content must be at least 200 characters")
	}

	if req.target_role != "student" && req.target_role != "lecturer" && req.target_role != "all" {
		return nil, fmt.Errorf("invalid target role")
	}

	announcement := &domain.Announcement{
		Title:      req.title,
		Content:    req.content,
		TargetRole: req.target_role,
	}

	return u.repository.Create(ctx, announcement)
}
