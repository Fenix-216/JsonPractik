package pl

//Основная структура
type Test struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first-name"`
	Features      []string  `json:"features"`
	Bio           []Bio     `json:"bio"`
	Health        []string  `json:"health"`
	Skills        []Skills  `json:"skills"`
	Communication []Communi `json:"communication"`
	Fire          []Fire    `json:""`
}

//Вложенные
type Bio struct {
	AgeBio    int    `json:"age-bio"`
	AgeHistor int    `json:"age-histor"`
	Gender    string `json:"gender"`
}

type Skills struct {
	Shooting        int `json:"shooting"`
	Meele           int `json:"meele"`
	Communication   int `json:"coommunication"`
	AnimalHusbandry int `json:"animal-husbandry"`
	Med             int `json:"med"`
	Cooking         int `json:"cooking"`
	Constr          int `json:"constr"`
	Farming         int `json:"farming"`
	Mining          int `json:"mining"`
	Art             int `json:"art"`
	Craft           int `json:"craft"`
	Research        int `json:"research"`
}

type Communi struct {
	Person     string `json:"person"`
	NamePerson string `json:"name-person"`
	SurName    string `json:"sur-name"`
	Buff       int    `json:"buff"`
}

type Fire struct {
	HaveFire bool `json:"have-fire"` // можно ли сделать так?
	TwoFire  int  `json:"two-fire"`
}

//ghbft
