package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
)

type supplierRepository struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) interfaces.SupplierRepository {
	return &supplierRepository{
		DB: db,
	}
}

func (r *supplierRepository) Get(ctx context.Context, params *model.GetSupplierRequest) (data []model.Supplier, total int64, err error) {
	var (
		offset = (params.Page - 1) * params.Limit
	)

	db := r.DB
	if params.Search != "" {
		db = db.Where("name ILIKE ?", "%"+params.Search+"%")
	}

	err = db.Order(fmt.Sprintf("%s %s", params.Sort, params.Order)).Offset(offset).Limit(params.Limit).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&model.Supplier{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *supplierRepository) GetById(ctx context.Context, id uuid.UUID) (data *model.Supplier, err error) {
	err = r.DB.First(&data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewCustomError(http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return nil, response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}
	return
}

func (r *supplierRepository) Create(ctx context.Context, payload *model.Supplier) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		payload.CreatedBy = ""
	} else {
		payload.CreatedBy = userID
	}
	if err := r.DB.Create(payload).Error; err != nil {
		return err
	}
	return nil
}

func (r *supplierRepository) Update(ctx context.Context, payload *model.Supplier, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		payload.UpdatedBy = ""
	} else {
		payload.UpdatedBy = userID
	}

	var existingItem model.Supplier
	if err := r.DB.Where("id = ?", id).First(&existingItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewCustomError(http.StatusNotFound, "Supplier tidak ditemukan", err)
		}
		return response.NewCustomError(http.StatusInternalServerError, "Gagal mengambil data item", err)
	}

	err = r.DB.Model(&model.Supplier{}).Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *supplierRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		userID = ""
	}

	// Set deleted_by field to the user who deleted the record
	err = r.DB.Model(&model.Supplier{}).
		Where("id = ?", id).
		Update("deleted_by", userID).Error
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	if err := r.DB.Delete(&model.Supplier{}, id).Error; err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Gagal menghapus data", err)
	}

	return nil
}
