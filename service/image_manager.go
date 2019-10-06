package service

// ImageManager represents the contract required to fetch image urls
type ImageManager interface {
	FetchImageUrlsByHashtag(hashtag string) ([]string, error)
}

type instagramDownloader struct {
	fetcherFunc InstagramFetcher
}

// InstagramFetcher represents a closure to abstract away instagram
type InstagramFetcher = func(tag string) ([]string, error)

// NewInstagramDownloader will return an Instagram implementation of an Image Manager
func NewInstagramDownloader(helperfunc InstagramFetcher) ImageManager {
	return &instagramDownloader{
		fetcherFunc: helperfunc,
	}
}

func (insta instagramDownloader) FetchImageUrlsByHashtag(hashtag string) ([]string, error) {
	fetchedImageUrls, err := insta.fetcherFunc(hashtag)
	if err != nil {
		return nil, err
	}
	return fetchedImageUrls, nil
}
