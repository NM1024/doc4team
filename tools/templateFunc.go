package tools

// HttpMethodList http请求方式的列表
func HttpMethodList() []string {
	return []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTION", "PATCH"}
}

// Obj2String 模版函数中类型转换
func Obj2String(obj interface{}) string {
	if obj == nil {
		return ""
	}
	return obj.(string)
}
