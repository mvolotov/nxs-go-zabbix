package zabbix

// ApiGet gets zabbix api version
func (z *Context) ApiGet() (string, int, error) {

	var result string

	status, err := z.request("alert.get", nil, &result)
	if err != nil {
		return "", status, err
	}

	return result, status, nil
}
