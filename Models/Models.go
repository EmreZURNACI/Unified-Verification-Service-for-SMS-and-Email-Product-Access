package models

type DatabaseModel struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Dbname   string `json:"Dbname"`
}
type VerificationModel struct {
	AccountSid         string `json:"accountSid"`
	AuthToken          string `json:"authToken"`
	VerificationString string `json:"verificationString"`
}
type Smtp struct {
	Address  string `json:"Address"`
	Password string `json:"Password"`
}
type VerifyCode struct {
	Code string `json:"code"`
}
type ConfigFile struct {
	Dbmodel    DatabaseModel     `json:"Database"`
	Verifmodel VerificationModel `json:"Verification"`
	Smtp       Smtp              `json:"Smtp"`
}
type DatabaseMessage struct {
	Statu   string `json:"statu"`
	Message string `json:"message"`
}
type BodyAuthRes struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	Verifytype int    `json:"verifytype"`
	Tel        string `json:"tel"`
}
type BodyProductRes struct {
	Search  string `json:"search"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	Sorting int    `json:"sorting"`
}
type Product struct {
	Id             int    `json:"id"`
	Marka          string `json:"marka"`
	Model          string `json:"model"`
	IsletimSistemi string `json:"isletimsistemi"`
}

// type DataModel struct {
// 	Veri []Product `json:"data"`
// }

type StandartResponseModel struct {
	Status     bool      `json:"status"`
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Data       []Product `json:"data"`
}

//gerekli fieldları doldurt mesaj gonder
/*

{
		"status": "success",
		"status_code": 200,
		"message": "İşlem başarılı",
		"data": {
		  "veri": "değer"
		}
	  }

*/
