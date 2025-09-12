//go:build premium

package features

func init() {
	features = append(features, "caching", "build flags")
}
