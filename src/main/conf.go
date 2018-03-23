package main


type ELuckyMoneyStruct struct{
	LuckyBoyPhone string  `json:"phone"`
	E_Url string	`json:"e_url"`
	HungryManName string `json:"name"`
	registerSuccess bool
	luckyMoneySuccess bool

	lucky_number int 
	platform string 
	sn string
	EGonaLucky bool
	isUsedLucky bool
}

type HungryManStruct struct{
	WechatCookie string  `json:"WechatCookie"`
	Phone string	`json:"Phone"`
	openid string
	eleme_key string
	realHungry bool
	nickname string
}
const myEOrigin string = "https://h5.ele.me"

type registerStruct struct{
	Sign string `json:"sign"`
	Phone string `json:"phone"`
}

type luckydoorStruct struct{
	Device_id string `json:"device_id"`
	Group_sn string `json:"group_sn"`
	Hardware_id string `json:"hardware_id"`
	Method string `json:"method"`
	Phone string `json:"phone"`
	Platform int `json:"platform"`
	Sign string `json:"sign"`
	Track_id string `json:"track_id"`
	Unionid string `json:"unionid"`
	Weixin_avatar string `json:"weixin_avatar"`
	Weixin_username string `json:"weixin_username"`
}
type luckydoorReturnJsonStruct struct{
	Promotion_records []Promotion_recordsStruct `json:"promotion_records"`
	Account string `json:"account"`
	Ret_code int `json:"ret_code"`
}

type Promotion_recordsStruct struct{
	Amount float32 `json:"amount"`
	Created_at int `json:"created_at"`
	Is_lucky bool `json:"is_lucky"`
	Sns_avatar string `json:"sns_avatar"`
	Sns_username string `json:"sns_username"`
}
const alreadyCome int  = 2
const successEdoor4 int  = 4
const successEdoor3 int  = 3
const failureEdoor int  = 1

const fivemost int  = 5