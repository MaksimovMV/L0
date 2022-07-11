package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

type Order struct {
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          delivery  `json:"delivery"`
	Payment           payment   `json:"payment"`
	Items             []item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (it item) Validate() error {
	return validation.ValidateStruct(&it,
		validation.Field(&it.ChrtId, validation.Required),
		validation.Field(&it.TrackNumber, validation.Required, is.UpperCase),
		validation.Field(&it.Price, validation.Required),
		validation.Field(&it.Rid, validation.Required, is.Alphanumeric),
		validation.Field(&it.Name, validation.Required),
		validation.Field(&it.Sale, validation.Required),
		validation.Field(&it.Size, validation.Required),
		validation.Field(&it.TotalPrice, validation.Required),
		validation.Field(&it.NmId, validation.Required),
		validation.Field(&it.Brand, validation.Required),
		validation.Field(&it.Status, validation.Required),
	)
}

func (d delivery) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Phone, validation.Required),
		validation.Field(&d.Zip, validation.Required, is.Digit),
		validation.Field(&d.City, validation.Required),
		validation.Field(&d.Address, validation.Required),
		validation.Field(&d.Email, validation.Required, is.Email),
	)
}

func (p payment) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Transaction, validation.Required, validation.Length(19, 22), is.Alphanumeric),
		validation.Field(&p.Currency, validation.Required, is.CurrencyCode),
		validation.Field(&p.Provider, validation.Required),
		validation.Field(&p.Amount, validation.Required),
		validation.Field(&p.PaymentDt, validation.Required),
		validation.Field(&p.Bank, validation.Required),
		validation.Field(&p.DeliveryCost, validation.Required),
		validation.Field(&p.GoodsTotal, validation.Required),
	)
}

func (o Order) Validate() error {

	return validation.ValidateStruct(&o,
		validation.Field(&o.OrderUid, validation.Required, validation.Length(19, 22), is.Alphanumeric),
		validation.Field(&o.TrackNumber, validation.Required, is.UpperCase),
		validation.Field(&o.Entry, validation.Required, is.UpperCase),
		validation.Field(&o.Delivery),
		validation.Field(&o.Payment),
		validation.Field(&o.Items),
		validation.Field(&o.Locale, validation.Required),
		validation.Field(&o.CustomerId, validation.Required),
		validation.Field(&o.DeliveryService, validation.Required),
		validation.Field(&o.Shardkey, validation.Required),
		validation.Field(&o.SmId, validation.Required),
		validation.Field(&o.DateCreated, validation.Required),
		validation.Field(&o.OofShard, validation.Required),
	)
}
