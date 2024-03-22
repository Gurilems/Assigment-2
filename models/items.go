package models

type Item struct {
	Item_ID     uint   `gorm:"primaryKey;column:item_id" json:"lineItemId"`
	ItemCode    string `gorm:"not null;type:VARCHAR(20)" json:"itemCode"`
	Description string `gorm:"type:TEXT" json:"description" `
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderID     uint   `gorm:"not null" json:"OrderID"`
}
