package domain

type RateLimiterRepository interface {
	Increment(key string, window int) (int64, error)
}

type RateLimiterUsecase interface {
	IsAllowed(key string, limit int, window int) (bool, error)
}
