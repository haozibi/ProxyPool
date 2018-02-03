package proxy

type proxyJson struct {
	IP               string `json:"ip"`           // 123.123.123.123:1024
	Class            int    `json:"class"`        // 0:http, 1:https
	ClassName        string `json:"class_name"`   // http or https
	HealthEvaluation int    `json:"health_value"` // 健康度 = 可用次数/ 总共测试次数
	TestCount        int    `json:"test_count"`   // 总共测试次数
	CreadtedAt       string `json:"creadted_at"`  // 收录时间
	UpdatedAt        string `json:"updated_at"`   // 更新时间
}
