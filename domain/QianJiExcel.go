package domain

type QianJiExcel struct {
	Time     string `json:"time"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Money    string `json:"money"`
	Account1 string `json:"account1"`
	Account2 string `json:"account2"`
	Remark   string `json:"remark"`
	Urls     string `json:"urls"`
}
