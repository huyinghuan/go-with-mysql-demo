package model

import (
	"maid-data-log/model/schema"
)

type LogBean struct{}

func (logBean *LogBean) Add(logs *schema.Log) (int64, error) {
	engine := GetDBConect()
	return engine.Insert(logs)
}

func (logBean *LogBean) FindAll() ([]schema.Log, error) {
	var logList []schema.Log
	engine := GetDBConect()
	if err := engine.Find(&logList); err != nil {
		return nil, err
	}
	return logList, nil
}
