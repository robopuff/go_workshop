package dictionary

type response struct {
	Entry Entries `json:"entry"`
	Items int     `json:"items_count"`
}

func MapEntries(entries Entries) response {
	return response{entries, len(entries)}
}
