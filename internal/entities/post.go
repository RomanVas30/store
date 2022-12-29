package entities

type StafferPost struct {
	OrgUnitName string  `json:"org_unit_name" binding:"required"`
	PostName    string  `json:"post_name"  binding:"required"`
	Rate        float64 `json:"rate" binding:"required"`
}

type OrgUnitPost struct {
	PostName string  `json:"post_name"  binding:"required"`
	Rate     float64 `json:"rate" binding:"required"`
}
