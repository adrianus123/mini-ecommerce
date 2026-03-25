package usecase

import "gateway-service/internal/domain"

type RateLimiterUsecase struct {
	repo domain.RateLimiterRepository
}

func NewRateLimiterUsecase(repo domain.RateLimiterRepository) *RateLimiterUsecase {
	return &RateLimiterUsecase{repo: repo}
}

func (u *RateLimiterUsecase) IsAllowed(key string, limit int, window int) (bool, error) {
	count, err := u.repo.Increment(key, window)
	if err != nil {
		return false, err
	}

	if count > int64(limit) {
		return false, nil
	}

	return true, nil
}
