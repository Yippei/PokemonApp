package model

import (
	"time"
)

type (
	RULE struct {
		RuleName    string    `json:"rulename"`
		RuleContent string    `json:"rule_content"`
		CreatedAt   time.Time `json:"creayed_at"`
		UpdateAt    time.Time `json:"update_at"`
	}

	CreatedRULEequest struct {
		RuleName    string `json:"rulename"`
		RuleContent string `json:"rule_content"`
	}

	CreatRULEResponse struct {
		RULE RULE `json:"rule"`
	}
)
