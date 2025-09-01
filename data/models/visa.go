package models

type AppApplicationType struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(500);not null;" json:"name"`
	Order  int16  `gorm:"not null;default:99" json:"order"`
	Status uint8  `gorm:"not null;default:0" json:"status"`

	// Relations
	VisaTypes []AppVisaType `gorm:"foreignKey:ApplicationTypeID" json:"visa_types,omitempty"`
}

// TableName overrides the default table name
func (AppApplicationType) TableName() string {
	return "app_application_type"
}

// -------------------------------------------------------------------

type AppVisaType struct {
	ID                 uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	ApplicationTypeID  uint    `gorm:"not null;default:0" json:"application_type_id"`
	Name               string  `gorm:"type:varchar(250);not null" json:"name"`
	Validitiy          int16   `gorm:"not null;default:1" json:"validitiy"`
	Entry              *string `gorm:"type:varchar(200)" json:"entry,omitempty"`
	CurrencyCode       string  `gorm:"type:enum('USD','GBP');not null;default:'USD'" json:"currency_code"`
	Amount             float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	OrderList          uint8   `gorm:"not null;default:99" json:"order_list"`
	SupportingDoc      string  `gorm:"type:enum('Y','N');not null;default:'N'" json:"supporting_doc"`
	SupportingDocLabel string  `gorm:"type:varchar(250);default:''" json:"supporting_doc_label"`
	Status             uint8   `gorm:"not null;default:1" json:"status"`

	// Relations
	ApplicationType AppApplicationType `gorm:"foreignKey:ApplicationTypeID" json:"application_type,omitempty"`
}

func (AppVisaType) TableName() string {
	return "app_visa_type"
}
