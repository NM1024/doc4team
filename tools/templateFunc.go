package tools

func HttpMethodList() []string {
	return []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTION", "PATCH"}
}

func Obj2String(obj interface{}) string {
	if obj == nil {
		return ""
	}
	return obj.(string)
}
