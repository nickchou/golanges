package entity

type EsResponse struct {
	Status int32 `json:"status"` //状态号=0成功
	Error  Error `json:"error"`  //错误信息

	Index       string `json:"_index"`        //索引名
	Type        string `json:"_type"`         //Type名
	Id          string `json:"_id"`           //es编号
	Version     int32  `json:"_version"`      //版本号
	Result      int32  `json:"result"`        //created=创建 update=更新
	SeqNo       string `json:"_seq_no"`       //?
	PrimaryTerm string `json:"_primary_term"` //分词数
	Shard       Shard  `json:"_shards"`       //词条明细信息
}
type Error struct {
	Type   string `json:"type"`   //错误类型
	Reason string `json:"reason"` //错误原因
}
type Shard struct {
	Total      int32 `json:"total"`      //总数量
	Successful int32 `json:"successful"` //成功数量
	Failed     int32 `json:"failed"`     //失败数量
}
