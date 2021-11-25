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

type professionService struct {
	db     *sqlc.Queries
	logger logger.Logger
}

func NewProfessionService(db *sqlc.Queries, log logger.Logger) *professionService {
	return &professionService{
		db:     db,
		logger: log,
	}
}

func (s *professionService) Create(ctx context.Context, req *position_service.CreateProfessionRequest) (*position_service.ProfessionId, error) {
	var profession sqlc.CreateProfessionParams
	err := helper.ParseToStruct(&profession, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, create profession", req, codes.Internal)
	}

	profession.ID = uuid.New().String()
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while generating new uuid", req, codes.Internal)
	}

	resp, err := s.db.CreateProfession(ctx, profession)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating profession", req, codes.Internal)
	}

	return &position_service.ProfessionId{
		Id: resp,
	}, nil
}

func (s *professionService) Update(ctx context.Context, req *position_service.Profession) (*position_service.ProfessionId, error) {
	var profession sqlc.UpdateProfessionParams
	err := helper.ParseToStruct(&profession, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, update profession", req, codes.Internal)
	}

	err = s.db.UpdateProfession(ctx, profession)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating profession", req, codes.Internal)
	}

	return &position_service.ProfessionId{
		Id: req.Id,
	}, nil
}

func (s *professionService) Get(ctx context.Context, req *position_service.ProfessionId) (*position_service.Profession, error) {
	var resp position_service.Profession
	profession, err := s.db.GetProfession(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting profession", req, codes.Internal)
	}

	err = helper.ParseToProto(&resp, profession)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to proto, update profession", req, codes.Internal)
	}

	return &resp, err
}

func (s *professionService) GetAll(ctx context.Context, req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error) {
	var (
		params      sqlc.GetProfessionsParams
		professions position_service.GetAllProfessionResponse
	)
	err := helper.ParseToStruct(&params, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, get all profession", req, codes.Internal)
	}

	resp, err := s.db.GetProfessions(ctx, params)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all profession", req, codes.Internal)
	}

	count, err := s.db.GetProfessionsCount(ctx, params.Search)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all profession", req, codes.Internal)
	}

	err = helper.ParseToProto(&professions, struct {
		Professions []sqlc.Profession `json:"professions"`
	}{
		Professions: resp,
	})

	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to proto, get all profession", req, codes.Internal)
	}

	professions.Count = int32(count)
	return &professions, nil
}

func (s *professionService) Delete(ctx context.Context, req *position_service.ProfessionId) (*emptypb.Empty, error) {
	err := s.db.DeleteProfession(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting profession", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
