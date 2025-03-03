package main

import "encoding/json"

type Order struct {
	AutoPay    bool        `json:"autoPay"`
	Orders     []OrderItem `json:"orders"`
	CustomInfo CustomInfo  `json:"customerId"`
}

type CustomInfo struct {
	Phone    string   `json:"phone,omitempty"`
	Identity Identity `json:"identity,omitempty"`
	Type     int      `json:"type,omitempty"`
	Email    string   `json:"email,omitempty"`
}

type Identity struct {
	AccountId string `json:"accountId,omitempty"`
}

type OrderItem struct {
	Items       []Item `json:"items,omitempty"`
	CycleType   int    `json:"cycleType,omitempty"`
	CycleCnt    int    `json:"cycleCnt,omitempty"`
	InstanceCnt int    `json:"instanceCnt,omitempty"`
}

type Item struct {
	ItemConfig   map[string]interface{} `json:"itemConfig,omitempty"`
	Master       bool                   `json:"master,omitempty"`
	ResourceType string                 `json:"resourceType,omitempty"`
	ServiceTag   string                 `json:"serviceTag,omitempty"`
	ItemValue    int                    `json:"itemValue,omitempty"`
}

func (o *Order) Marshal() string {
	data, _ := json.Marshal(o)
	return string(data)
}
