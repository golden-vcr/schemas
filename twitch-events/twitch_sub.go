package etwitch

import "fmt"

func getCreditMultiplierFromTier(tier string) (int, error) {
	// We expect the "tier" value in EventSub messages to be one of the following:
	switch tier {
	case "1000":
		// Tier 1 subs are the baseline at $5; fun points are credited 1x
		return 1, nil
	case "2000":
		// Tier 2 subs cost $10; fun point credits are doubled
		return 2, nil
	case "3000":
		// Tier 3 subs are $25; so subscribers get 5x fun points
		return 5, nil
	}

	// In the event of an unrecognized value, fail
	return 0, fmt.Errorf("unrecognized tier value '%s'", tier)
}
