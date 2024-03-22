package models

import (
	"time"
)

type Order struct {
	Order_ID     uint      `gorm:"primaryKey;column:order_id;uniqueIndex:idx_order_id" json:"OrderID"`
	CustomerName string    `gorm:"type:VARCHAR(50);not null" json:"customerName"`
	Items        []Item    `json:"Items" gorm:"foreignKey:order_id;constraint:OnDelete:CASCADE"`
	OrderedAt    time.Time `gorm:"not null;type:timestamp;autoCreateTime" json:"orderedAt"`
}
