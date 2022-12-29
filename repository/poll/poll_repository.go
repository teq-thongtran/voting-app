package poll

import (
	"context"

	"myapp/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, data *model.Poll) error
	Update(ctx context.Context, data *model.Poll) error
	GetByID(ctx context.Context, id int64) (*model.Poll, error)
	Delete(ctx context.Context, data *model.Poll, unscoped bool) error
	GetList(
		ctx context.Context,
		page int,
		limit int,
		conditions interface{},
		order []string,
	) ([]model.Poll, int64, error)
	GetListPollIds(ctx context.Context) ([]string, error)
}

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Create(ctx context.Context, data *model.Poll) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) Update(ctx context.Context, data *model.Poll) error {
	return p.getDB(ctx).Save(data).Error
}

func (p *pgRepository) GetByID(ctx context.Context, id int64) (*model.Poll, error) {
	var poll model.Poll

	err := p.getDB(ctx).Debug().
		Where("id = ?", id).
		First(&poll).
		Error

	if err != nil {
		return nil, err
	}

	return &poll, nil
}

func (p *pgRepository) Delete(ctx context.Context, data *model.Poll, unscoped bool) error {
	db := p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	return db.Select("PollOptions", "UserPolls").Debug().Delete(data).Error
}

func (p *pgRepository) GetList(
	ctx context.Context,
	page int,
	limit int,
	conditions interface{},
	order []string,
) ([]model.Poll, int64, error) {
	var (
		db     = p.getDB(ctx).Model(&model.Poll{}).Debug()
		data   = make([]model.Poll, 0)
		total  int64
		offset int
		ids    []string
	)

	if conditions != nil {
		db = db.Where(conditions)
	}

	ids, err := p.GetListPollIds(ctx)

	if err != nil {
		return nil, 0, err
	}

	db = db.Where("id IN ?", ids).Or("poll_policy = ? ", "public").Or("user_id = ?", ctx.Value("user_id"))

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

	err = db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}

func (p *pgRepository) GetListPollIds(ctx context.Context) ([]string, error) {
	var (
		db  = p.getDB(ctx).Model(&model.UserPoll{}).Debug()
		ids = make([]string, 0)
	)

	err := db.Where("user_id = ?", ctx.Value("user_id")).Pluck("poll_id", &ids).Error

	if err != nil {
		return nil, err
	}

	return ids, nil
}
