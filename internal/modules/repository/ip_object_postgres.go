package repository

import (
	"fmt"
	"server/internal/common/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IpObjectPostgres struct {
	db *sqlx.DB
}

const ipObjectTable = "ip_objects"

func NewIpObjectPostgres(db *sqlx.DB) *IpObjectPostgres {
	return &IpObjectPostgres{
		db: db,
	}
}

func (r *IpObjectPostgres) CreateIpObject(ipObject *entity.IpObject) error {
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, title, description, jurisdiction, patent_type) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, title, description, jurisdiction, patent_type, created_at, updated_at", ipObjectTable)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return r.db.QueryRow(query, id, ipObject.UserId, ipObject.Title, ipObject.Description, ipObject.Jurisdiction, ipObject.PatentType).Scan(&ipObject.Id, &ipObject.UserId, &ipObject.Title, &ipObject.Description, &ipObject.Jurisdiction, &ipObject.PatentType, &ipObject.CreatedAt, &ipObject.UpdatedAt)
}

func (r *IpObjectPostgres) GetIpObjectsByUserId(userId string) ([]*entity.IpObject, error) {
	query := fmt.Sprintf("SELECT id, user_id, title, description, jurisdiction, patent_type, created_at, updated_at FROM %s WHERE user_id = $1", ipObjectTable)
	var ipObjects []*entity.IpObject
	if err := r.db.Select(&ipObjects, query, userId); err != nil {
		return nil, err
	}
	return ipObjects, nil
}

func (r *IpObjectPostgres) GetIpObjectsById(id string) (*entity.IpObject, error) {
	query := fmt.Sprintf("SELECT id, user_id, title, description, jurisdiction, patent_type, created_at, updated_at FROM %s WHERE id = $1", ipObjectTable)
	var ipObject entity.IpObject
	err := r.db.Get(&ipObject, query, id)

	return &ipObject, err
}
