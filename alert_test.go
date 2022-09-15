package zabbix

import (
	"testing"
	"time"
)

//

const (
	testAlertHostGroupID = "4"
)

func TestAlert(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Get
	testAlertGet(t, z)
}

func testAlertGet(t *testing.T, z Context) []AlertObject {

	AObjects, _, err := z.AlertGet(AlertGetParams{
		GroupIDs: []string{testAlertHostGroupID},
		TimeTill: int(time.Now().Unix()) + 3*60,
		GetParameters: GetParameters{
			Output: SelectExtendedOutput,
			Limit:  3,
		},
	})

	if err != nil {
		t.Error("Alert get error:", err)
	} else {
		if len(AObjects) == 0 {
			t.Error("Alert get error: unable to find log")
		} else {
			t.Logf("Alert get: success")
		}
	}

	return AObjects
}
