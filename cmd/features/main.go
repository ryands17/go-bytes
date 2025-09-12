package features

var features []string

func init() {
	features = []string{
		"sets",
		"branded types",
	}
}

func AvailableFeatures() []string {
	return features
}
