package ratelimiter

import (
	limiter "github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// Setting up rate limiter
func SetupLimiter() *limiter.Limiter {

	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted("10-S")
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()

	// Then, create the limiter instance which takes the store and the rate as arguments.
	// Now, you can give this instance to any supported middleware.
	instance := limiter.New(store, rate)

	return instance
}
