package urlarchive

type URLsArchive interface {
	Add(id string, url string)
	Get(id string) (string, error)
}
