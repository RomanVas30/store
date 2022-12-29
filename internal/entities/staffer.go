package entities

type Staffer struct {
	Id             int           `json:"-" db:"id"`
	FIO            string        `json:"fio" db:"fio" binding:"required"`
	BirthYear      int           `json:"birth_year" db:"birth_year" binding:"required"`
	EmploymentDate string        `json:"-" db:"employment_date"`
	SNILS          string        `json:"snils" db:"snils"`
	Posts          []StafferPost `json:"posts"`
}

type ShortStaffer struct {
	FIO       string `json:"fio" db:"fio" binding:"required"`
	BirthYear int    `json:"birth_year" db:"birth_year" binding:"required"`
	SNILS     string `json:"snils" db:"snils"`
}

type ActionStaffer struct {
	FIO         string        `json:"fio"`
	SNILS       string        `json:"snils"`
	AddPosts    []StafferPost `json:"add_posts"`
	DeletePosts []StafferPost `json:"delete_posts"`
}

func NewStafferByUpdateStaffer(us *ActionStaffer) *Staffer {
	return &Staffer{
		FIO:   us.FIO,
		SNILS: us.SNILS,
	}
}
