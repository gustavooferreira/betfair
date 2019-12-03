package market

import "github.com/gustavooferreira/betfair/pkg/exchangestream"

type CacheManager struct {
	// caches are not thread-safe
	caches map[string]MarketCache
	// Message ID
	msgID uint32
}

func NewCacheManager() CacheManager {
	caches := make(map[string]MarketCache)
	cm := CacheManager{caches: caches}

	return cm
}

// Process gets a MarketChangeM message and updates the relevant caches accordingly
// Returns a list with the MarketIDs of the caches that were updated
func (cm *CacheManager) Process(mcm exchangestream.MarketChangeM) ([]string, error) {

	// Detect if ID has changed!
	// If msgID is zero, accept the first ID in the message
	// After that lock in all message changes to that ID
	// When a new higher ID shows up, that means there was a subscription change
	// Accept the new ID, update the msgID and reject everything else!

	// If you get a marketID update that wasn't in the cache that might mean we delete it?

	// Support segmentation

	return []string{}, nil
}

// GetCachesAvailable returns the list of MarketIDs available in the caches.
func (cm *CacheManager) GetCachesAvailable(mcm exchangestream.MarketChangeM) ([]string, error) {
	result := []string{}
	for marketID := range cm.caches {
		result = append(result, marketID)
	}

	return result, nil
}
