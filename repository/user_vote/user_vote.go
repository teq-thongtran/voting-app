package usvo

import (
	"context"

	"myapp/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, data *model.UserVote) error
	GetByID(ctx context.Context, id int64) (*model.UserVote, error)
	Delete(ctx context.Context, data *model.UserVote, unscoped bool) error
	GetList(
		ctx context.Context,
		page int,
		limit int,
		conditions interface{},
		pollID int64,
		order []string,
	) ([]model.UserVote, int64, error)
}

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Create(ctx context.Context, data *model.UserVote) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) GetByID(ctx context.Context, id int64) (*model.UserVote, error) {
	var UserVote model.UserVote

	err := p.getDB(ctx).Debug().
		Where("id = ?", id).
		First(&UserVote).
		Error

	if err != nil {
		return nil, err
	}

	return &UserVote, nil
}

func (p *pgRepository) Delete(ctx context.Context, data *model.UserVote, unscoped bool) error {
	db := p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	return db.Delete(data).Error
}

func (p *pgRepository) GetList(
	ctx context.Context,
	page int,
	limit int,
	conditions interface{},
	pollID int64,
	order []string,
) ([]model.UserVote, int64, error) {
	var (
		db     = p.getDB(ctx).Model(&model.UserVote{}).Debug()
		data   = make([]model.UserVote, 0)
		total  int64
		offset int
	)

	db = db.Joins("JOIN poll_options PollOption ON PollOption.id = user_votes.poll_option_id AND PollOption.poll_id = ?", pollID)

	if conditions != nil {
		db = db.Where(conditions)
	}

	for i := range order {
		db = db.Order(order[i])
	}

	if page != 1 {
		offset = limit * (page - 1)
	}

	if limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}
