package entities

import (
	"context"
	"strings"
	"time"

	"google.golang.org/appengine/datastore"
)

const ENTITY_CATEGORY = "category"
const ENTITY_CATEGORY_PRODUCT = "category_product"

type Category struct {
	Id              int64 `datastore:"-" json:"id"`
	Name            string
	Path            string
	Description     string
	MetaDescription string
	Created         time.Time
	Thumbnail       string
	Featured        bool
}

func (c *Category) SetMissingDefaults() {
	if c.Thumbnail == "" {
		c.Thumbnail = "/assets/images/stock.jpeg"
	}

	if c.Path == "" {
		c.Path = c.Name
	}
}

func NewCategory(name string) *Category {
	return &Category{
		Name:      name,
		Created:   time.Now(),
		Thumbnail: "/assets/images/stock.jpeg",
	}
}

func ListCategories(ctx context.Context) ([]*Category, error) {
	categories := make([]*Category, 0)
	keys, err := datastore.NewQuery(ENTITY_CATEGORY).GetAll(ctx, &categories)
	if err != nil {
		return nil, err
	}

	for index, key := range keys {
		var category = categories[index]
		category.Id = key.IntID()
		category.SetMissingDefaults()
	}

	return categories, nil
}

func CreateCategory(ctx context.Context, name string) (*Category, error) {
	c := NewCategory(name)
	c.Path = name
	c.Path = strings.TrimSpace(c.Path)
	c.Path = strings.ToLower(c.Path)
	c.Path = strings.Replace(c.Path, " ", "-", -1)
	c.Path = strings.Replace(c.Path, "'", "", -1)

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, ENTITY_CATEGORY, nil), c)
	if err != nil {
		return nil, err
	}

	c.Id = key.IntID()
	return c, nil
}
