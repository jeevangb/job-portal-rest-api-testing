package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
}
type JobLocation struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
type Technology struct {
	gorm.Model
	Tname string `json:"technologyName" gorm:"unique"`
}
type WorkMode struct {
	gorm.Model
	WMode string `json:"workMode" gorm:"unique"`
}
type Qualification struct {
	gorm.Model
	Qname string `json:"qualification" gorm:"unique"`
}
type Shift struct {
	gorm.Model
	Shifttype string `json:"shift" gorm:"unique"`
}
type JobType struct {
	gorm.Model
	JobType string `json:"jobType" gorm:"unique"`
}
type ResponseJobId struct {
	ID uint
}

type Jobs struct {
	gorm.Model
	Company         Company         `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint            `json:"cid"`
	Title           string          `json:"title"`
	MinNoticePeriod string          `json:"minnp"`
	MaxNoticePeriod string          `json:"maxnp"`
	Budget          string          `json:"budget"`
	JobLocation     []JobLocation   `gorm:"many2many:job_location;"`
	Technology      []Technology    `gorm:"many2many:technology;"`
	WorkMode        []WorkMode      `gorm:"many2many:work_mode;"`
	Jobdescription  string          `json:"job_description"`
	Qualification   []Qualification `gorm:"many2many:qualification;"`
	Shift           []Shift         `gorm:"many2many:shift;"`
	JobType         []JobType       `gorm:"many2many:job_type;"`
}

type Hr struct {
	Title          string `json:"title"`
	Minnp          string `json:"minnp"`
	Maxnp          string `json:"maxnp"`
	Budget         string `json:"budget"`
	JobLocation    []uint `json:"job_location"`
	Technology     []uint `json:"technology"`
	WorkMode       []uint `json:"workmode"`
	JobDescription string `json:"job_description"`
	Qualification  []uint `json:"qualification"`
	Shift          []uint `json:"shift"`
	JobType        []uint `json:"job_type"`
}
type RespondJobApplicant struct {
	Name string       `json:"name"`
	Jid  uint         `json:"jid"`
	Jobs JobApplicant `json:"job_appication"`
}
type JobApplicant struct {
	Jid            uint   `json:"cid"`
	Title          string `json:"title"`
	Salary         string `json:"salary"`
	Np             string `json:"np"`
	Budget         string `json:"budget"`
	JobLocation    []uint `json:"job_location"`
	Technology     []uint `json:"technology"`
	WorkMode       []uint `json:"workmode"`
	JobDescription string `json:"job_description"`
	Qualification  []uint `json:"qualification"`
	Shift          []uint `json:"shift"`
	JobType        []uint `json:"job_type"`
}
