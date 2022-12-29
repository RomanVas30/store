package entities

type OrgUnit struct {
	Id          int           `json:"-" db:"id"`
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description"`
	Posts       []OrgUnitPost `json:"posts"`
	AddPosts    []OrgUnitPost `json:"add_posts"`
	DeletePosts []OrgUnitPost `json:"delete_posts"`
}
