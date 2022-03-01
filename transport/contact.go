package transport

import (
	"context"

	"github.com/annbelievable/go-microservice-gokit/endpoint"
	"github.com/annbelievable/go-microservice-gokit/proto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getOrganisationContacts grpc.Handler
}

func NewGrpcServer(endpoints endpoint.Endpoints, logger log.Logger) proto.ContactServiceServer {
	return &grpcServer{
		getOrganisationContacts: grpc.NewServer(
			endpoints.GetOrganisationContacts,
			decodeOrganisationContactsRequest,
			encodeContactsResponse,
		),
	}
}

func (s *grpcServer) GetOrganisationContacts(ctx context.Context, req *proto.OrganisationContactRequest) (*proto.ContactResponse, error) {
	_, resp, err := s.getOrganisationContacts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.ContactResponse), nil
}

func decodeOrganisationContactsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.OrganisationContactRequest)
	return endpoint.OrganisationContactRequest{OrganisationId: req.OrganisationId}, nil
}

func encodeContactsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*proto.ContactResponse_Contact)
	return resp, nil
}
