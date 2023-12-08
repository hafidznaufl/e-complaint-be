package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
	"time"
)

func NewsCreateRequestToNewsDomain(request web.NewsCreateRequest) *domain.News {
	return &domain.News{
		Admin_ID: request.Admin_ID,
		Title:    request.Title,
		Content:  request.Content,
		Date:     time.Now(),
	}
}

func NewsUpdateRequestToNewsDomain(request web.NewsUpdateRequest) *domain.News {
	return &domain.News{
		Admin_ID: request.Admin_ID,
		Title:    request.Title,
		Content:  request.Content,
		Date:     request.Date,
	}
}

func NewsDomaintoNewsSchema(request domain.News) *schema.News {
	return &schema.News{
		Admin_ID: request.Admin_ID,
		Title:    request.Title,
		Content:  request.Content,
		Date:     request.Date,
	}
}
