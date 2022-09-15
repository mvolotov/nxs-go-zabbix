package zabbix

// AlertObject struct is used to store zabbix alert contains information about whether certain action operations have been executed successfully.
//
// see: https://www.zabbix.com/documentation/6.0/en/manual/api/reference/alert/object#alert-object
type AlertObject struct {
	AlertID       string `json:"alertid,omitempty"`
	ActionID      string `json:"actionid,omitempty"`
	AlertType     int    `json:"alerttype,omitempty"`
	Clock         int    `json:"clock,omitempty"`
	Error         string `json:"error,omitempty"`
	EscStep       int    `json:"esc_step,omitempty"`
	EventID       string `json:"eventid,omitempty"`
	MediaTypeID   string `json:"mediatypeid,omitempty"`
	Message       string `json:"message,omitempty"`
	Retries       int    `json:"retries,omitempty"`
	SendTo        string `json:"sendto,omitempty"`
	Status        int    `json:"status,omitempty"`
	Subject       string `json:"subject,omitempty"`
	UserID        string `json:"userid,omitempty"`
	PEventID      string `json:"p_eventid,omitempty"`
	AcknowledgeID string `json:"acknowledgeid.omitempty"`
}

// AuditlogGetParams struct is used for auditlog get requests
//
// see: https://www.zabbix.com/documentation/6.0/ru/manual/api/reference/auditlog/get#parameters
type AlertGetParams struct {
	GetParameters

	AlertIDs     []string `json:"alertids,omitempty"`
	ActionIDs    []string `json:"actionids,omitempty"`
	EventIDs     []string `json:"eventids,omitempty"`
	GroupIDs     []string `json:"groupids,omitempty"`
	HostIDs      []string `json:"hostids,omitempty"`
	MediaTypeIDs []string `json:"mediatypeids,omitempty"`
	ObjectIDs    []string `json:"objectids,omitempty"`
	UserIDs      []string `json:"userids,omitempty"`
	EventObject  int      `json:"eventobject,omitempty"`
	EventSource  int      `json:"eventsource,omitempty"`
	TimeFrom     int      `json:"time_from,omitempty"`
	TimeTill     int      `json:"time_till,omitempty"`

	SelectHosts      SelectQuery `json:"selectHosts,omitempty"`
	SelectMediaTypes SelectQuery `json:"selectMediatypes,omitempty"`
	SelectUsers      SelectQuery `json:"selectUsers,omitempty"`
}

// AuditlogGet gets auditlog according to the given parameters
func (z *Context) AlertGet(params AlertGetParams) ([]AlertObject, int, error) {

	var result []AlertObject

	status, err := z.request("alert.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
