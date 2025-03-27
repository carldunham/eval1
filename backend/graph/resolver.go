package graph

type Resolver struct {
	service *Service
}

func NewResolver() (*Resolver, error) {
	service, err := NewService()
	if err != nil {
		return nil, err
	}
	return &Resolver{service: service}, nil
}
