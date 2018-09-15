package menu

type Menu struct {
	// gorm.Model
	Name string `json:"name,omitempty" gorm:"column:name"`
	Url  string `json:"url,omitempty" gorm:"column:url"`
}

func (Menu) TableName() string {
	return "krcmf_menu"
}
