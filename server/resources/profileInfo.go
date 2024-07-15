package resources

type OwnProfileInfo struct {
	ContributionsDispersion float32  `json:"contributionsDispersion"`
	Type                    string   `json:"type"`
	Languages               []string `json:"languages"`
}
