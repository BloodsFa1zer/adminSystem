package model

import (
	"time"
)

type Tag struct {
	ID        uint       `gorm:"primaryKey"`
	PublicID  string     `json:"public_id" gorm:"type:uuid;not null"`
	Title     string     `json:"title" gorm:"unique;not null" validate:"required"`
	Products  []*Product `json:"products" gorm:"many2many:tag_product; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}

type Product struct {
	ID         uint      `gorm:"primaryKey"`
	PublicID   string    `json:"public_id" gorm:"type:uuid;not null"`
	Title      string    `json:"title"`
	Subtitle   string    `json:"subtitle"`
	Content    string    `json:"content"`
	Price      float64   `json:"price"`
	IsInStock  bool      `json:"isInStock" gorm:"type:boolean"`
	Tags       []*Tag    `json:"tags" gorm:"many2many:tag_product; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Images     []*Image  `json:"images" gorm:"many2many:image_product; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UniverseID uint      `json:"-"`
	Universe   Universe  `gorm:"foreignKey:UniverseID;references:ID;onUpdate:CASCADE,onDelete:SET NULL"`
	Currency   string    `json:"currency"`
	CreatedAt  time.Time `gorm:"index" json:"-"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"-"`
}

type TagProduct struct {
	ID        uint `gorm:"primaryKey"`
	TagID     uint
	ProductID uint
}

func (TagProduct) TableName() string {
	return "tag_product"
}

func (Product) TableName() string {
	return "products"
}

type Image struct {
	Filename  string     `json:"filename" gorm:"unique:idx_filename_case_sensitive; not null"`
	Path      string     `json:"path" gorm:"not null" `
	PublicID  string     `json:"public_id" gorm:"type:uuid;not null"`
	ID        uint       `json:"-" gorm:"primaryKey"`
	Products  []*Product `json:"image_products" gorm:"many2many:image_product; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time  `gorm:"index" json:"-"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"-"`
}

type ImageProduct struct {
	ID        uint `gorm:"primaryKey"`
	ImageID   uint
	ProductID uint
}

func (ImageProduct) TableName() string {
	return "image_product"
}

type ImageUpdate struct {
	Filename string `json:"filename" validate:"required" gorm:"unique:idx_filename_case_sensitive;not null"`
}

func (Image) TableName() string {
	return "images"
}

//type Tag struct {
//	ID        uint       `json:"-" gorm:"primaryKey"`
//	PublicID  string     `json:"public_id" gorm:"type:uuid;not null"`
//	Title     string     `json:"title" gorm:"unique;not null" validate:"required"`
//	Products  []*Product `json:"tag_products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
//	CreatedAt time.Time  `json:"-"`
//	UpdatedAt time.Time  `json:"-"`
//}

func (Tag) TableName() string {
	return "tags"
}

type TagInput struct {
	Title string `json:"title" gorm:"unique;not null" validate:"required"`
}

type Universe struct {
	ID        uint       `json:"-" gorm:"primaryKey"`
	PublicID  string     `json:"public_id" gorm:"type:uuid;not null"`
	Content   string     `json:"content"  gorm:"type:text; not null" validate:"required"`
	Title     string     `json:"title" gorm:"unique;not null" validate:"required"`
	ProductID []*Product `json:"product_id" gorm:"many2many:image_product;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}

func (Universe) TableName() string {
	return "universe"
}

type SortOptions struct {
	SortBy string `json:"sortBy"`
	Name   string `json:"name"`
	Offset int    `json:"offset"`
}
