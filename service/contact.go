package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

type contactService struct {
	logger log.Logger
}

type ContactService interface {
	GetOrganisationContacts(ctx context.Context, organisationId uint32) ([]*Contact, error)
}

func NewContactService(logger log.Logger) ContactService {
	return &contactService{
		logger: logger,
	}
}

func (s contactService) GetOrganisationContacts(ctx context.Context, organisationId uint32) ([]*Contact, error) {
	//i will have to filter through the contacts and get the one that matches the organisation id
	return contactList, nil
}

type Contact struct {
	Id           uint32
	Email        string
	Organisation uint32
}

// dummy data
var contactList = []*Contact{
	&Contact{
		Id:           1,
		Email:        "user1@mail.com",
		Organisation: 1,
	},
	&Contact{
		Id:           2,
		Email:        "user2@mail.com",
		Organisation: 1,
	},
	&Contact{
		Id:           3,
		Email:        "user3@mail.com",
		Organisation: 2,
	},
	&Contact{
		Id:           4,
		Email:        "user4@mail.com",
		Organisation: 2,
	},
	&Contact{
		Id:           5,
		Email:        "user5@mail.com",
		Organisation: 3,
	},
	&Contact{
		Id:           6,
		Email:        "user6@mail.com",
		Organisation: 3,
	},
}
