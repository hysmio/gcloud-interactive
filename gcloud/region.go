package gcloud

// GetActiveRegion gets active gcloud region
func GetActiveRegion() (string, error) {
	return GetConfig("compute/zone")
}

// SetActiveRegion sets gcloud active compute region
func SetActiveRegion(region string) error {
	return SetConfig("compute/zone", region)
}
