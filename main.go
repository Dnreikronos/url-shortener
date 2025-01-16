var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)
	shorten := flag.String("Shorten", "", "Url to shorten")
	expand := flag.String("Expand", "", "Shortened URL to Expand")

	flag.Parse()
