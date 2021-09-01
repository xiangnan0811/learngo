package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(s string) string {
	return r.Contents
}
