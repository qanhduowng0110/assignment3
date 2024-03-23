package model

type RequestViewModel struct {
	Request_url string 							`json:"request_url`
	Request_method string 						`json:"request_method"`
	Request_headers map[string]interface{} 		`json:"request_headers"`
	Request_body map[string]interface{} 		`json:"request_body"`
	Response_status_code int 					`json:"response_status_code"`
	Response_body map[string]interface{} 		`json:"response_body"`
	Request_timestamp int64 					`json:"request_timestamp"`
	Created_at int64 							`json:"created_at"`
	Updated_at int64 							`json:"updated_at"`
}