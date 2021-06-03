package model

type Employee struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Skills     struct {
		Soft []string `json:"soft"`
		Hard []string `json:"hard"`
	} `json:"skills"`
	Salary struct {
		Main  int     `json:"main"`
		Bonus float64 `json:"bonus"`
	} `json:"salary"`
}
