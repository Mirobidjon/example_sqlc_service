package sqlc

import (
	"context"
	"database/sql"

	"github.com/xtgo/uuid"
)

type PositionWithTx struct {
	*Queries
	db *sql.DB
}

func NewPosition(db *sql.DB) *PositionWithTx {
	return &PositionWithTx{
		db:      db,
		Queries: New(db),
	}
}

type CreatePositionRequest struct {
	CreatePositionParams
	Attributes []CreatePositionAttributeParams `json:"position_attributes"`
}

type GetPosition struct {
	Position
	PositionAttributes []GetPositionAttributesRow `json:"position_attributes"`
}

type GetAllPositions struct {
	Positions []*GetPosition `json:"positions"`
	Count     int32          `json:"count"`
}

func (r *PositionWithTx) CreatePositionTx(ctx context.Context, req CreatePositionRequest) (string, error) {
	var resp string
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	queries := New(tx)

	resp, err = queries.CreatePosition(ctx, CreatePositionParams{
		ID:           req.ID,
		Name:         req.Name,
		ProfessionID: req.ProfessionID,
	})

	if err != nil {
		return "", err
	}

	for _, attribute := range req.Attributes {
		attribute.ID = uuid.NewRandom().String()
		attribute.PositionID = resp
		_, err = queries.CreatePositionAttribute(ctx, attribute)
		if err != nil {
			return "", err
		}
	}

	return resp, nil
}

func (r *PositionWithTx) UpdatePositionTx(ctx context.Context, req CreatePositionRequest) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	queries := New(tx)

	err = queries.UpdatePosition(ctx, UpdatePositionParams{
		Name:         req.Name,
		ProfessionID: req.ProfessionID,
		ID:           req.ID,
	})

	if err != nil {
		return err
	}

	err = queries.DeletePositionAttribute(ctx, req.ID)
	if err != nil {
		return err
	}

	for _, attribute := range req.Attributes {
		attribute.ID = uuid.NewRandom().String()
		attribute.PositionID = req.ID
		_, err = queries.CreatePositionAttribute(ctx, attribute)
		if err != nil {
			return err
		}
	}

	return nil
}
