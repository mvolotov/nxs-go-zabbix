package zabbix

// AuditlogObject struct is used to store zabbix audit logs results.
//
// see: https://www.zabbix.com/documentation/6.0/ru/manual/api/reference/auditlog/object
type AuditlogObject struct {
	AuditID  string `json:"auditid,omitempty"`
	UserID   string `json:"userid,omitempty"`
	Username string `json:"username,omitempty"`
	Clock    int    `json:"clock,omitempty"`
	IP       string `json:"ip,omitempty"`
	Action   int    `json:"action,omitempty"`

	ResourceType int    `json:"resourcetype,omitempty"`
	ResourceID   string `json:"resourceid,omitempty"`
	ResourceName string `json:"resourcename,omitempty"`

	RecordsetID string `json:"recordsetid,omitempty"`
	Details     string `json:"details,omitempty"`
}

// AuditlogGetParams struct is used for auditlog get requests
//
// see: https://www.zabbix.com/documentation/6.0/ru/manual/api/reference/auditlog/get#parameters
type AuditlogGetParams struct {
	GetParameters

	AuditIDs []string `json:"auditids,omitempty"`
	UserIDs  []string `json:"userids,omitempty"`
	TimeFrom int      `json:"time_from,omitempty"`
	TimeTill int      `json:"time_till,omitempty"`
}

// AuditlogGet gets auditlog according to the given parameters
func (z *Context) AuditlogGet(params AuditlogGetParams) ([]AuditlogObject, int, error) {

	var result []AuditlogObject

	status, err := z.request("auditlog.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
