package schema

//Log 虚拟机配置
type Log struct {
	ID             int64  `xorm:"unique" json:"id"`
	Belong         string `xorm:"belong" json:"belong"`
	URI            string `xorm:"uri" json:"uri"`
	UpdatedAt      int64  `xorm:"updated" json:"updated"`
	Data           string `xorm:"longtext 'data'" json:"data"`
	RequestOptions string `xorm:"text 'request_options'" json:"request_options"`
	Status         string `xorm:"status" json:"status"`
}
