package application

import "context"

//fileID: file_10086
//URL:    https://cdn.example.com/demo.docx

type QueryOfficeService struct {
}

func NewQueryOfficeService() *QueryOfficeService {
	return &QueryOfficeService{}
}

// (*domain.Feed, error)
func (s *QueryOfficeService) OpenWPSFile(ctx context.Context) {
	//content, err := domain.NewFeedContent(rawContent)
	//if err != nil {
	//	return nil, err
	//}
	//feed, err := domain.NewFeed(content)
	//if err != nil {
	//	return nil, err
	//}
	//return s.repo.Create(ctx, feed)
}
