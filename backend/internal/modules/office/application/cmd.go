package application

import (
	"context"
)

type CmdOfficeService struct {
}

func NewCmdOfficeService() *CmdOfficeService {
	return &CmdOfficeService{}
}

// (*domain.Feed, error)
func (s *QueryOfficeService) Create(ctx context.Context) {
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
