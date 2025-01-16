var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)
