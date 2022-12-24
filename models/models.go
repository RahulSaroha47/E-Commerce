package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID      `json:"id" bson:"id"`
	FirstName       *string                 `json:"first_name"  validate:"required,min=2,max=30"`
	LastName        *string                 `json:"last_name"   validate:"required,min=2,max=30"`
	Password        *string                 `json:"password"    validate:"required,min=6"`
	Email           *string                 `json:"email"       validate:"email, required"`     
	Phone           *string                 `json:"phone"       validate:"required"` 
	Token           *string                 `json:"token"`
	RefreshToken    *string                 `json:"refresh_token"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	UserId          *string                 `json:"user_id"`
	UserCart        []ProductUser           `json:"user_cart" bson:"user_cart"`
	AddressDetails  []Address               `json:"address" bson:"address"`
	OrderStatus     []Order                 `json:"orders" bson:"orders"`
}

type Product struct {
	ProductID       primitive.ObjectID      `bson:"product_id"`
	ProductName     *string					`json:"product_name"`
	Price           *int					`json:"price"`
	Rating          *int					`json:"rating"`
	Image           *string					`json:"image"`
}

type ProductUser struct {
	ProductID      primitive.ObjectID       `bson:"product_id"`
	ProductName    *string                  `json:"product_name" bson:"producr_name"`
	Price          *int						`json:"price" bson:"price"`
	Rating         *int						`json:"rating" bson:"rating"`
	Image          *string					`json:"image" bson:"image"`
}

type Address struct {
	AddressId    primitive.ObjectID        `bson:"id"`
	House        *string					`json:"house" bson:"house"`
	Street       *string					`json:"street" bson:"street"`
	City         *string 					`json:"city" bson:"city"`
	Pincode      *string					`json:"pincode" bson:"pincode"`
}

type Order struct {
	OrderId        primitive.ObjectID    `bson:"id"`
	OrderCart      []ProductUser         `json:"order_list" bson:"order_list"`
	OrderedAt      time.Time             `json:"ordered_at" bson:"ordered_at"`
	Price          *int                  `json:"total_price" bson:"total_price"`
	Discount       *int                  `json:"discount" bson:"dicount"`
	PaymentMethod  Payment               `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool              `json:"digital" bson:"digital"`
	COD     bool              `json:"cod" bson:"cod"`
}

