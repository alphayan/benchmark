package benchmark

func ano() {
	type param struct {
		Name     string `json:"name"`      //信息名称
		Openid   string `json:"openid"`    //openid
		InfoData string `json:"info_data"` //入库的json
		Start    string `json:"start"`     //开始时间
		End      string `json:"end"`       //结束时间
		Type     string `json:"type"`      //个人或者商务
		Image    string `json:"image"`     //图片
		Num      int    `json:"num"`       //次数
	}
	var p param
	_ = p
}

type param struct {
	Name     string `json:"name"`      //信息名称
	Openid   string `json:"openid"`    //openid
	InfoData string `json:"info_data"` //入库的json
	Start    string `json:"start"`     //开始时间
	End      string `json:"end"`       //结束时间
	Type     string `json:"type"`      //个人或者商务
	Image    string `json:"image"`     //图片
	Num      int    `json:"num"`       //次数
}

func unano() {
	var p param
	_ = p
}
