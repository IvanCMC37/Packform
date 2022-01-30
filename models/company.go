package models

type company struct {
	ID   int    `json: company_id`
	Name string `json: company_name`
}

// type company struct {
// 	ID   int    `json: "COMPANY_ID" gorm: "primary_key"`
// 	Name string `json: "COMPANY_NAME"`
// }
