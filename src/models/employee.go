package models

type Employee struct {
    ID              string `json:"id" bson:"_id,omitempty"`
    Name            string `json:"name" bson:"name"`
    NIK             string `json:"nik" bson:"nik"`
    LastEducation   string `json:"last_education" bson:"last_education"`
    JoinDate        string `json:"join_date" bson:"join_date"`
    EmploymentStatus string `json:"employment_status" bson:"employment_status"` // "KONTRAK", "TETAP"
}