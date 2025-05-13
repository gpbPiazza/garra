package minuta

type DataFinder interface {
	StartKey() string
	EndKey() string
	Find(text string) string
}
