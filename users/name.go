package users

import "fmt"

type FullName struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
}

func (f *FullName) String() string {
	return fmt.Sprintf("%s %s %s", f.FirstName, f.LastName, f.Patronymic)
}

type Position struct {
	JobTitle string `json:"job_title"`
}

func (p *Position) String() string {
	return fmt.Sprintf("%s", p.JobTitle)
}	