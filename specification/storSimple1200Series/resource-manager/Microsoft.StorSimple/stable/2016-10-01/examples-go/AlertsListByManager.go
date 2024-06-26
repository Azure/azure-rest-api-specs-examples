package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/AlertsListByManager.json
func ExampleAlertsClient_NewListByManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsClient().NewListByManagerPager("ResourceGroupForSDKTest", "hManagerForSDKTest4", &armstorsimple1200series.AlertsClientListByManagerOptions{Filter: to.Ptr("status%20eq%20'Active'%20and%20severity%20eq%20'Critical'")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AlertList = armstorsimple1200series.AlertList{
		// 	Value: []*armstorsimple1200series.Alert{
		// 		{
		// 			Name: to.Ptr("a1f3a945-74dc-4387-bf17-a4f226374833"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/alerts"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hManagerForSDKTest4/devices/is2-hlmdhdgu1et/alerts/a1f3a945-74dc-4387-bf17-a4f226374833"),
		// 			Properties: &armstorsimple1200series.AlertProperties{
		// 				AlertType: to.Ptr("Cloud Connectivity"),
		// 				AppearedAtSourceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-02T04:08:12.914Z"); return t}()),
		// 				AppearedAtTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-02T04:08:12.914Z"); return t}()),
		// 				ClearedAtSourceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				ClearedAtTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DetailedInformation: map[string]*string{
		// 					"first observed": to.Ptr("1/1/0001 12:00:00 AM (UTC)"),
		// 					"last observed": to.Ptr("8/2/2018 4:08:12 AM (UTC)"),
		// 					"occurrences": to.Ptr("2"),
		// 				},
		// 				ErrorDetails: &armstorsimple1200series.AlertErrorDetails{
		// 					ErrorCode: to.Ptr(""),
		// 					ErrorMessage: to.Ptr(""),
		// 					Occurences: to.Ptr[int32](2),
		// 				},
		// 				Recommendation: to.Ptr("The StorSimple Device Manager service has not received a heartbeat from the virtual array. The virtual array may be offline. This could be due to one of the following reasons. <br/>\r\n	1. The virtual array may be turned off or paused on the hypervisor. In Hyper-V, your virtual array will be paused automatically when the volume that stores snapshots or virtual hard disks, runs out of available storage. The state of the virtual array is listed as <i>\"paused-critical\"</i> in Hyper-V Manager. Resolve this by creating additional space on the volume. If you still cannot connect, check the Hyper-V host or ESX server to ensure that the VM is healthy.<br/><br/>\r\n	2. The virtual array is not able to communicate with the StorSimple Device Manager service. In the local web UI of the virtual array, go to <b>Troubleshooting > Diagnostic tests</b> and click <b>Run diagnostic tests</b>. Resolve the reported issues. <br/> <br/>\r\n  If the issue persists, contact your network administrator."),
		// 				ResolutionReason: to.Ptr(""),
		// 				Scope: to.Ptr(armstorsimple1200series.AlertScopeDevice),
		// 				Severity: to.Ptr(armstorsimple1200series.AlertSeverityCritical),
		// 				Source: &armstorsimple1200series.AlertSource{
		// 					Name: to.Ptr("is2-hlmdhdgu1et"),
		// 					AlertSourceType: to.Ptr(armstorsimple1200series.AlertSourceTypeDevice),
		// 					TimeZone: to.Ptr("UTC"),
		// 				},
		// 				Status: to.Ptr(armstorsimple1200series.AlertStatusActive),
		// 				Title: to.Ptr("No device heartbeat received for 5 minutes."),
		// 			},
		// 	}},
		// }
	}
}
