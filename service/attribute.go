package service

import (
	"context"

	"bitbucket.org/udevs/ur_go_user_service/genproto/position_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/helper"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"bitbucket.org/udevs/ur_go_user_service/storage/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type attributeService struct {
	db     *sqlc.Queries
	logger logger.Logger
}

func NewAttributeService(db *sqlc.Queries, log logger.Logger) *attributeService {
	return &attributeService{
		db:     db,
		logger: log,
	}
}

func (s *attributeService) Create(ctx context.Context, req *position_service.CreateAttributeRequest) (*position_service.AttributeId, error) {
	var attribute sqlc.CreateAttributeParams
	err := helper.ParseToStruct(&attribute, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, create attribute", req, codes.Internal)
	}

	attribute.ID = uuid.New().String()
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while generating new uuid", req, codes.Internal)
	}

	resp, err := s.db.CreateAttribute(ctx, attribute)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating attribute", req, codes.Internal)
	}

	return &position_service.AttributeId{
		Id: resp,
	}, nil
}

func (s *attributeService) Update(ctx context.Context, req *position_service.Attribute) (*position_service.AttributeId, error) {
	var attribute sqlc.UpdateAttributeParams
	err := helper.ParseToStruct(&attribute, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, update attribute", req, codes.Internal)
	}

	err = s.db.UpdateAttribute(ctx, attribute)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating attribute", req, codes.Internal)
	}

	return &position_service.AttributeId{
		Id: req.Id,
	}, nil
}

func (s *attributeService) Get(ctx context.Context, req *position_service.AttributeId) (*position_service.Attribute, error) {
	var resp position_service.Attribute
	attribute, err := s.db.GetAttribute(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting attribute", req, codes.Internal)
	}

	err = helper.ParseToProto(&resp, attribute)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to proto, update attribute", req, codes.Internal)
	}

	return &resp, err
}

func (s *attributeService) GetAll(ctx context.Context, req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error) {
	var (
		params     sqlc.GetAttributesParams
		Attributes position_service.GetAllAttributeResponse
	)
	err := helper.ParseToStruct(&params, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, get all attribute", req, codes.Internal)
	}

	resp, err := s.db.GetAttributes(ctx, params)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all attribute", req, codes.Internal)
	}

	count, err := s.db.GetAttributesCount(ctx, params.Search)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all attribute", req, codes.Internal)
	}

	err = helper.ParseToProto(&Attributes, struct {
		Attributes []sqlc.Attribute `json:"attributes"`
	}{
		Attributes: resp,
	})

	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to proto, get all attribute", req, codes.Internal)
	}

	Attributes.Count = int32(count)
	return &Attributes, nil
}

func (s *attributeService) Delete(ctx context.Context, req *position_service.AttributeId) (*emptypb.Empty, error) {
	err := s.db.DeleteAttribute(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting attribute", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
