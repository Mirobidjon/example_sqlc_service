package service

import (
	"context"
	"database/sql"

	"bitbucket.org/udevs/ur_go_user_service/genproto/position_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/helper"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"bitbucket.org/udevs/ur_go_user_service/storage/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type positionService struct {
	db     *sqlc.Queries
	logger logger.Logger
	db_tx  *sqlc.PositionWithTx
}

func NewPositionService(db *sqlc.Queries, log logger.Logger, db_tx *sql.DB) *positionService {
	return &positionService{
		db:     db,
		logger: log,
		db_tx:  sqlc.NewPosition(db_tx),
	}
}

func (s *positionService) Create(ctx context.Context, req *position_service.CreatePositionRequest) (*position_service.PositionId, error) {
	var position sqlc.CreatePositionRequest
	err := helper.ParseToStruct(&position, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, create position", req, codes.Internal)
	}

	position.ID = uuid.New().String()
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while generating new uuid", req, codes.Internal)
	}

	resp, err := s.db_tx.CreatePositionTx(ctx, position)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating position", req, codes.Internal)
	}

	return &position_service.PositionId{
		Id: resp,
	}, nil
}

func (s *positionService) Update(ctx context.Context, req *position_service.Position) (*position_service.PositionId, error) {
	var position sqlc.CreatePositionRequest
	err := helper.ParseToStruct(&position, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, update position", req, codes.Internal)
	}

	err = s.db_tx.UpdatePositionTx(ctx, position)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating position", req, codes.Internal)
	}

	return &position_service.PositionId{
		Id: req.Id,
	}, nil
}

func (s *positionService) Get(ctx context.Context, req *position_service.PositionId) (*position_service.Position, error) {
	var (
		resp     position_service.Position
		response sqlc.GetPosition
	)
	position, err := s.db.GetPosition(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting position", req, codes.Internal)
	}

	err = helper.ParseToJson(&response, position)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to json, get position", req, codes.Internal)
	}

	response.PositionAttributes, err = s.db.GetPositionAttributes(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting position attributes", req, codes.Internal)
	}

	err = helper.ParseToProto(&resp, response)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to proto, get position", req, codes.Internal)
	}

	return &resp, err
}

func (s *positionService) GetAll(ctx context.Context, req *position_service.GetAllPositionRequest) (*position_service.GetAllPositionResponse, error) {
	var (
		params    sqlc.GetPositionsParams
		positions position_service.GetAllPositionResponse
		response  sqlc.GetAllPositions
	)
	err := helper.ParseToStruct(&params, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to struct, get all position", req, codes.Internal)
	}

	resp, err := s.db.GetPositions(ctx, params)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all position", req, codes.Internal)
	}

	count, err := s.db.GetPositionsCount(ctx, params.Search)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all position count", req, codes.Internal)
	}

	err = helper.ParseToJson(&response, struct {
		Positions []sqlc.Position `json:"positions"`
	}{
		Positions: resp,
	})

	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to json, get all position", req, codes.Internal)
	}

	for _, position := range response.Positions {
		position.PositionAttributes, err = s.db.GetPositionAttributes(ctx, position.ID)
		if err != nil {
			return nil, helper.HandleError(s.logger, err, "error while getting all position attributes", req, codes.Internal)
		}
	}

	response.Count = int32(count)

	err = helper.ParseToProto(&positions, response)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while parsing to proto, get all position", req, codes.Internal)
	}

	return &positions, nil
}

func (s *positionService) Delete(ctx context.Context, req *position_service.PositionId) (*emptypb.Empty, error) {
	err := s.db.DeletePosition(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting position", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
