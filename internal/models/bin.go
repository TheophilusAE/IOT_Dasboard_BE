package models

type BinType string

const (
	BinOrganic   BinType = "organic"
	BinAnorganic BinType = "anorganic"
)

func (b BinType) Valid() bool {
	return b == BinOrganic || b == BinAnorganic
}

func NormalizeBinType(raw string) (BinType, bool) {
	switch normalizeToken(raw) {
	case "organic":
		return BinOrganic, true
	case "anorganic", "inorganic", "non-organic", "nonorganic":
		return BinAnorganic, true
	default:
		return "", false
	}
}
