package endpoint

import (
	"context"

	"github.com/annbelievable/go-microservice-gokit/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetOrganisationContacts endpoint.Endpoint
}

type OrganisationContactRequest struct {
	OrganisationId uint32
}

type ContactResponse struct {
	Contacts []*service.Contact
}

func MakeContactEndpoints(s service.ContactService) Endpoints {
	return Endpoints{
		GetOrganisationContacts: makeGetOrganisationContactEnpoints(s),
	}
}

func makeGetOrganisationContactEnpoints(s service.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(OrganisationContactRequest)
		result, _ := s.GetOrganisationContacts(ctx, req.OrganisationId)
		return ContactResponse{Contacts: result}, nil
	}
}
