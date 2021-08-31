package retrieve

type Retriever interface {
	Get(string) string
}
