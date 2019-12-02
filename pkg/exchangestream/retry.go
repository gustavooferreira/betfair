package exchangestream

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"time"
)

// PolicyExponential defines an executer for the exponential retry policy.
type PolicyExponential struct {
	// Retries specifies the number of retries allowed.
	// -1 means infinite number of retries
	// 0 means to never retry
	// Any other number greater or equal to 1 means retry X amount of times
	Retries int
	// MaximunBackoff specifies the maximun waiting time between actions (in seconds)
	MaximumBackoff uint
	currentCount   uint
	threshhold     bool
}

// NewPolicyExponential creates a new PolicyExponential.
// Accepts a PolicyExponential struct as the first argument and the initialCount as second.
// For most use cases you will want to set initialCount to zero.
func NewPolicyExponential(retries int, maximunBackoff uint, initialCount uint) PolicyExponential {
	return PolicyExponential{Retries: retries, MaximumBackoff: maximunBackoff, currentCount: initialCount}
}

// BackOffTime returns the time to wait (in milliseconds) until next retry.
// If returned value is -1 it means no more retries.
func (pee *PolicyExponential) BackOffTime() int {

	if pee.Retries >= 0 {
		if pee.Retries == 0 || int(pee.currentCount) >= pee.Retries {
			return -1
		}
	}

	if pee.threshhold {
		pee.currentCount++
		return int(pee.MaximumBackoff * 1000)
	}

	calcedWaitTime := (math.Pow(2, float64(pee.currentCount)) + rand.Float64()) * 1000
	pee.currentCount++

	if calcedWaitTime > float64(pee.MaximumBackoff)*1000 {
		pee.threshhold = true
		return int(pee.MaximumBackoff * 1000)
	}

	return int(calcedWaitTime)
}

// WaitBackOff waits for a certain time until next retry.
// It accepts a context so the caller can cancel the waiting time and get back control.
// In case of a cancellation, the first returned value will be true, false otherwise.
// If no more retries should happen, this function will return immediately with an error.
func (pee *PolicyExponential) WaitBackOff(ctx context.Context) (bool, error) {
	result := pee.BackOffTime()

	if result == -1 {
		return false, errors.New("No more retries allowed")
	}

	duration := time.Duration(result) * time.Millisecond
	cancelled := sleepCanBreak(ctx, duration)

	return cancelled, nil
}

// RetryAllowed returns a boolean reflecting whether it's still allowed to retry or not.
func (pee PolicyExponential) RetryAllowed() bool {
	if pee.Retries >= 0 {
		if pee.Retries == 0 || int(pee.currentCount) >= pee.Retries {
			return false
		}
	}
	return true
}

// RetryCount returns the number of retry attempts so far.
func (pee PolicyExponential) RetryCount() uint {
	return pee.currentCount
}

// sleepCanBreak is an helper function that sleeps for a specified duration and can be stopped via the context passed in.
func sleepCanBreak(ctx context.Context, sleep time.Duration) (isBreak bool) {
	select {
	case <-ctx.Done():
		isBreak = true
	case <-time.After(sleep):
		isBreak = false
	}
	return
}
