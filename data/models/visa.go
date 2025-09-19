package models

type AppApplicationType struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
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
	ID                 uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	ApplicationTypeID  uint64  `gorm:"not null;default:0" json:"application_type_id"`
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

type AppApplication struct {
	ID                  uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	ApplicationID       *string `gorm:"type:varchar(50)" json:"application_id,omitempty"`
	ApplicationTypeID   *uint64 `gorm:"type:int unsigned" json:"application_type_id,omitempty"`
	ApplicationTypeName *string `gorm:"type:varchar(300)" json:"application_type_name,omitempty"`
	VisaTypeID          *uint64 `gorm:"type:int unsigned" json:"visa_type_id,omitempty"`
	VisaTypeName        *string `gorm:"type:varchar(300)" json:"visa_type_name,omitempty"`
	VisaValiditiy       uint16  `gorm:"type:smallint unsigned;default:0" json:"visa_validitiy"`
	VisaEntry           *string `gorm:"type:varchar(100)" json:"visa_entry,omitempty"`
	VisaCurrencyCode    *string `gorm:"type:varchar(10)" json:"visa_currency_code,omitempty"`
	VisaAmount          float64 `gorm:"type:decimal(10,2);default:0.00" json:"visa_amount"`
	Fname               string  `gorm:"type:varchar(100)" json:"fname"`
	Mname               string  `gorm:"type:varchar(100)" json:"mname"`
	Lname               string  `gorm:"type:varchar(100)" json:"lname,omitempty"`
	// PassportType           *string    `gorm:"type:varchar(100)" json:"passport_type,omitempty"`
	// Nationality            *uint16    `gorm:"type:smallint unsigned" json:"nationality,omitempty"`
	// NationalityCode        *string    `gorm:"type:char(2)" json:"nationality_code,omitempty"`
	// PortOfArrivalID        *uint      `gorm:"type:int unsigned" json:"port_of_arrival,omitempty"`
	// PortExit               *uint      `gorm:"type:int unsigned" json:"port_exit,omitempty"`
	// PassportNo             *string    `gorm:"type:varchar(100)" json:"passportno,omitempty"`
	// PassportIssuePlace     *string    `gorm:"type:varchar(100)" json:"passport_issue_place,omitempty"`
	// PassportIssueDate      *time.Time `gorm:"type:date" json:"passport_issue_date,omitempty"`
	// PassportExpiryDate     *time.Time `gorm:"type:date" json:"passport_expiry_date,omitempty"`
	// OtherICCountry         *uint16    `gorm:"type:smallint unsigned" json:"other_ic_country,omitempty"`
	// OtherICNo              *string    `gorm:"type:varchar(100)" json:"other_ic_no,omitempty"`
	// OtherICIssuePlace      *string    `gorm:"type:varchar(100)" json:"other_ic_issue_place,omitempty"`
	// OtherICIssueDate       *time.Time `gorm:"type:date" json:"other_ic_issue_date,omitempty"`
	// OtherICNationality     *uint16    `gorm:"type:smallint unsigned" json:"other_ic_nationality,omitempty"`
	// DOB                    *time.Time `gorm:"type:date" json:"dob,omitempty"`
	// DateOfArrival          *time.Time `gorm:"type:date" json:"date_of_arrival,omitempty"`
	// Email                  *string    `gorm:"type:varchar(100)" json:"email,omitempty"`
	// Phone                  *string    `gorm:"type:varchar(100)" json:"phone,omitempty"`
	// HavePreviousName       *string    `gorm:"type:char(1)" json:"have_previous_name,omitempty"`
	// PreviousSurname        *string    `gorm:"type:varchar(100)" json:"previous_surname,omitempty"`
	// PreviousName           *string    `gorm:"type:varchar(100)" json:"previous_name,omitempty"`
	// Gender                 string     `gorm:"type:char(1);default:'M'" json:"gender"`
	// BirthCity              *string    `gorm:"type:varchar(100)" json:"birth_city,omitempty"`
	// BirthCountry           *int       `gorm:"type:smallint" json:"birth_country,omitempty"`
	// NationalID             *string    `gorm:"type:varchar(50)" json:"national_id,omitempty"`
	// Religion               *string    `gorm:"type:varchar(50)" json:"religion,omitempty"`
	// VisibleIdentification  *string    `gorm:"type:varchar(100)" json:"visible_identification_marks,omitempty"`
	// Qualification          *string    `gorm:"type:varchar(100)" json:"qualification,omitempty"`
	// AcquireNationality     *string    `gorm:"type:varchar(50)" json:"acquire_nationality,omitempty"`
	// PrevNationality        *int       `gorm:"type:int" json:"prev_nationality,omitempty"`
	// HouseNo                *string    `gorm:"type:varchar(50)" json:"houseno,omitempty"`
	// PermanentHouseNo       *string    `gorm:"type:varchar(50)" json:"permanent_houseno,omitempty"`
	// City                   *string    `gorm:"type:varchar(50)" json:"city,omitempty"`
	// PermanentCity          *string    `gorm:"type:varchar(50)" json:"permanent_city,omitempty"`
	// District               *string    `gorm:"type:varchar(50)" json:"district,omitempty"`
	// PermanentDistrict      *string    `gorm:"type:varchar(50)" json:"permanent_district,omitempty"`
	// ZipCode                *string    `gorm:"type:varchar(10)" json:"zipcode,omitempty"`
	// AddressCountry         *uint16    `gorm:"type:smallint unsigned" json:"address_country,omitempty"`
	// Mobile                 *string    `gorm:"type:varchar(12)" json:"mobile,omitempty"`
	// FatherName             *string    `gorm:"type:varchar(100)" json:"father_name,omitempty"`
	// FatherNationality      *uint16    `gorm:"type:smallint unsigned" json:"father_nationality,omitempty"`
	// FatherPreNationality   *uint16    `gorm:"type:smallint unsigned" json:"father_pre_nationality,omitempty"`
	// FatherBirthPlace       *string    `gorm:"type:varchar(100)" json:"father_birth_place,omitempty"`
	// FatherBirthCountry     *uint16    `gorm:"type:smallint unsigned" json:"father_birth_country,omitempty"`
	// MotherName             *string    `gorm:"type:varchar(100)" json:"mother_name,omitempty"`
	// MotherNationality      *uint16    `gorm:"type:smallint unsigned" json:"mother_nationality,omitempty"`
	// MotherPreNationality   *uint16    `gorm:"type:smallint unsigned" json:"mother_pre_nationality,omitempty"`
	// MotherBirthPlace       *string    `gorm:"type:varchar(100)" json:"mother_birth_place,omitempty"`
	// MotherBirthCountry     *uint16    `gorm:"type:smallint unsigned" json:"mother_birth_country,omitempty"`
	// IsMarried              *string    `gorm:"type:enum('Y','N');default:'N'" json:"is_married,omitempty"`
	// SpouseName             *string    `gorm:"type:varchar(100)" json:"spouse_name,omitempty"`
	// SpouseNationality      *uint16    `gorm:"type:smallint unsigned" json:"spouse_nationality,omitempty"`
	// SpousePreNationality   *uint16    `gorm:"type:smallint unsigned" json:"spouse_pre_nationality,omitempty"`
	// SpouseBirthPlace       *string    `gorm:"type:varchar(100)" json:"spouse_birth_place,omitempty"`
	// SpouseBirthCountry     *uint16    `gorm:"type:smallint unsigned" json:"spouse_birth_country,omitempty"`
	// GrandParentPakistan    *string    `gorm:"type:enum('Y','N');default:'N'" json:"grand_parent_pakistan,omitempty"`
	// GrandParentPakDetail   *string    `gorm:"type:varchar(300)" json:"grand_parent_pak_detail,omitempty"`
	// Occupation             *string    `gorm:"type:varchar(100)" json:"occupation,omitempty"`
	// EmployerName           *string    `gorm:"type:varchar(100)" json:"employer_name,omitempty"`
	// Designation            *string    `gorm:"type:varchar(100)" json:"designation,omitempty"`
	// EmployerAddress        *string    `gorm:"type:varchar(200)" json:"employer_address,omitempty"`
	// EmployerPhone          *string    `gorm:"type:varchar(18)" json:"employer_phone,omitempty"`
	// PastOccupation         *string    `gorm:"type:varchar(100)" json:"past_occupation,omitempty"`
	// InMilitaryService      *string    `gorm:"type:enum('Y','N');default:'N'" json:"in_military_service,omitempty"`
	// MilitaryOrganisation   *string    `gorm:"type:varchar(100)" json:"military_organisation,omitempty"`
	// MilitaryDesignation    *string    `gorm:"type:varchar(100)" json:"military_designation,omitempty"`
	// MilitaryRank           *string    `gorm:"type:varchar(100)" json:"military_rank,omitempty"`
	// MilitaryPosting        *string    `gorm:"type:varchar(100)" json:"military_posting,omitempty"`
	// PlacesLikeToVisit      *string    `gorm:"type:varchar(300)" json:"places_like_to_visit,omitempty"`
	// VisitedIndiaBefore     *string    `gorm:"type:enum('Y','N');default:'N'" json:"visited_india_before,omitempty"`
	// VisitedAddress         string     `gorm:"type:text;not null" json:"visited_address"`
	// VisitedCities          *string    `gorm:"type:text" json:"visited_cities,omitempty"`
	// LastVisaNo             *string    `gorm:"type:varchar(50)" json:"last_visa_no,omitempty"`
	// LastVisaType           *string    `gorm:"type:varchar(100)" json:"last_visa_type,omitempty"`
	// LastVisaIssuePlace     *string    `gorm:"type:varchar(100)" json:"last_visa_issue_place,omitempty"`
	// LastVisaIssueDate      *time.Time `gorm:"type:date" json:"last_visa_issue_date,omitempty"`
	// ExtendVisaRefuse       *string    `gorm:"type:enum('Y','N');default:'N'" json:"extend_visa_refuse,omitempty"`
	// ExtendVisaRefuseDetail *string    `gorm:"type:text" json:"extend_visa_refuse_detail,omitempty"`
	// VisitedCountries       *string    `gorm:"type:text" json:"visited_countries,omitempty"`
	// RefName                *string    `gorm:"type:varchar(100)" json:"ref_name,omitempty"`
	// RefAddress             *string    `gorm:"type:text" json:"ref_address,omitempty"`
	// RefPhone               *string    `gorm:"type:varchar(20)" json:"ref_phone,omitempty"`
	// RefHomeName            *string    `gorm:"type:varchar(100)" json:"ref_home_name,omitempty"`
	// RefHomeAddress         *string    `gorm:"type:text" json:"ref_home_address,omitempty"`
	// RefHomePhone           *string    `gorm:"type:varchar(20)" json:"ref_home_phone,omitempty"`
	// Image                  *string    `gorm:"type:varchar(100)" json:"image,omitempty"`
	// Passport               *string    `gorm:"type:varchar(100)" json:"passport,omitempty"`
	// SupportingDoc          *string    `gorm:"type:varchar(100)" json:"supporting_doc,omitempty"`
	// Status                 *uint8     `gorm:"type:tinyint unsigned;default:0" json:"status,omitempty"`
	// PaymentStatus          *uint8     `gorm:"type:tinyint unsigned;default:0" json:"payment_status,omitempty"`
	// PaymentComment         *string    `gorm:"type:text" json:"payment_comment,omitempty"`
	// PaymentDate            *time.Time `gorm:"type:datetime" json:"payment_date,omitempty"`
	// CreatedAt              time.Time  `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt              *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

// TableName overrides the default pluralized name
func (AppApplication) TableName() string {
	return "app_application"
}
