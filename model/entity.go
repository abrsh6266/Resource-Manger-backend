package model

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Department struct {
	Name     	 string `json:"name"`
	Id       	 int `json:"id"`
	Location 	 string `json:"location"`
	UpdateColumn string `json:"column"`
	UpdateValue  string `json:"value"`
	Custodians    []Custodian `json:"Custodians"`
}
type Custodian struct {
	Name     	 string `json:"name"`
	Id       	 int `json:"id"`
	Department 	 string `json:"department"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	UpdateColumn string `json:"column"`
	UpdateValue  string `json:"value"`
}
type Category struct {
	Name     	 string `json:"name"`
	Id       	 int `json:"id"`
	Total	     int `json:"total"`
	UpdateValue  string `json:"value"`
	Materials    []Material `json:"Materials"`
}
type MaterialUserRelation struct {
	Name     	 string `json:"userName"`
	Id       	 int `json:"id"`
	Amount     	 int `json:"amount"`
	SerialNumber string `json:"materialsSerialNumber"`
	Material     Material `json:"Material"`
	Custodian    Custodian `json:"Custodian"`
}
type Material struct {
	Name     	 string `json:"name"`
	Id       	 int `json:"id"`
	SerialNumber string `json:"serialNumber"`
	Owner        string `json:"owner"`
	Model 	     string `json:"model"`
	Processor 	 string `json:"processor"`
	DiskType   	 string `json:"diskType"`
	ScanType   	 string `json:"scanType"`
	NetworkType  string `json:"networkType"`
	Type         string `json:"type"`
	Total     	 int    `json:"total"`
	Taken     	 int    `json:"taken"`
	UpdateColumn string `json:"column"`
	UpdateValue  string `json:"value"`
	MaterialUserRelations []MaterialUserRelation `json:"MaterialCustodianRels"`
}