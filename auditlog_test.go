package zabbix

import (
	"testing"
	"time"
)

var (
	testAuditlogUserID   = "1"
	testAuditlogTimeFrom = time.Now().Unix() - (3 * 60)
	testAuditlogTimeTill = time.Now().Unix() + (3 * 60)
)

func TestAuditlog(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Get
	testAuditlogGet(t, z)
}

func testAuditlogGet(t *testing.T, z Context) []AuditlogObject {

	AlObjects, _, err := z.AuditlogGet(AuditlogGetParams{
		UserIDs:  []string{testAuditlogUserID},
		TimeFrom: int(testAuditlogTimeFrom),
		TimeTill: int(testAuditlogTimeTill),
		GetParameters: GetParameters{
			Limit: 3,
		},
	})

	if err != nil {
		t.Error("Auditlog get error:", err)
	} else {
		if len(AlObjects) == 0 {
			t.Error("Auditlog get error: unable to find log")
		} else {
			t.Logf("Auditlog get: success")
		}
	}

	return AlObjects
}
