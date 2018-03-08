package entity

/*
用户实体信息
*/
type User struct {
	Id       int64  `json:"id" xml:"IID"`     //编号
	Name     string `json:"name" xml:"NName"` //姓名
	Ctfid    string `json:"ctfid"`            //？
	Gender   string `json:"gender"`           //F：女性 M：男性
	Birthday string `json:"birthday"`         //生日格式yyyyMMdd
	Age      int32  `json:"age"`              //年龄，根据生日去算
	City     string `json:"city"`             //模拟数据
	Address  string `json:"address"`          //地址
	Zip      string `json:"zip"`              //邮政编码
	Mobile   string `json:"mobile"`           //手机号码
	Tel      string `json:"tel"`              //电话号码
	Fax      string `json:"fax"`              //传真
	Email    string `json:"email"`            //邮箱
}
