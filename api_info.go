package zabbix

// ApiGet gets zabbix api version
func (z *Context) ApiGetVersion(ZbxAPIHost string) (string, int, error) {

	zbuff := z.host
	z.host = ZbxAPIHost

	var result string

	status, err := z.request("apiinfo.version", nil, &result)
	if err != nil {
		return "", status, err
	}

	z.host = zbuff

	return result, status, nil
}
