package models

type EmployeeResponse struct {
    Name             string     `json:"name"`
    JoinDate         string		`json:"join_date"`
    EmploymentStatus string     `json:"employment_status"`
}