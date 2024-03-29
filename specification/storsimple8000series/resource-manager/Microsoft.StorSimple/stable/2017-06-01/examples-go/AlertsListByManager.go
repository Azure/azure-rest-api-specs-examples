package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/AlertsListByManager.json
func ExampleAlertsClient_NewListByManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsClient().NewListByManagerPager("ResourceGroupForSDKTest", "ManagerForSDKTest1", &armstorsimple8000series.AlertsClientListByManagerOptions{Filter: to.Ptr("status%20eq%20'Active'%20and%20appearedOnTime%20ge%20'2017-06-09T18:30:00Z'%20and%20appearedOnTime%20le%20'2017-06-19T18:30:00Z'%20and%20sourceType%20eq%20'Device'%20and%20sourceName%20eq%20'Device05ForSDKTest'")})
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
		// page.AlertList = armstorsimple8000series.AlertList{
		// 	Value: []*armstorsimple8000series.Alert{
		// 		{
		// 			Name: to.Ptr("308b5bd2-824b-436f-840e-44bde075ef33"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/alerts"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/alerts/308b5bd2-824b-436f-840e-44bde075ef33"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.AlertProperties{
		// 				AlertType: to.Ptr("Security"),
		// 				AppearedAtSourceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-14T11:35:41.226Z"); return t}()),
		// 				AppearedAtTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-14T11:35:41.226Z"); return t}()),
		// 				ClearedAtSourceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				ClearedAtTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DetailedInformation: map[string]*string{
		// 				},
		// 				ErrorDetails: &armstorsimple8000series.AlertErrorDetails{
		// 					Occurences: to.Ptr[int32](1),
		// 				},
		// 				Recommendation: to.Ptr("Your device might be under attack or an authorized user is attempting to connect with an incorrect password. Check that your SnapShot manager instances are configured with the correct password. Once you have taken appropriate action, please clear this alert from the alerts page."),
		// 				Scope: to.Ptr(armstorsimple8000series.AlertScopeDevice),
		// 				Severity: to.Ptr(armstorsimple8000series.AlertSeverityWarning),
		// 				Source: &armstorsimple8000series.AlertSource{
		// 					Name: to.Ptr("Device05ForSDKTest"),
		// 					AlertSourceType: to.Ptr(armstorsimple8000series.AlertSourceTypeDevice),
		// 					TimeZone: to.Ptr("Pacific Standard Time"),
		// 				},
		// 				Status: to.Ptr(armstorsimple8000series.AlertStatusActive),
		// 				Title: to.Ptr("4 login attempts failed for SnapShotManager"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("3c3df90a-cbbc-4cc4-bd63-b54ca05997da"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/alerts"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/alerts/3c3df90a-cbbc-4cc4-bd63-b54ca05997da"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.AlertProperties{
		// 				AlertType: to.Ptr("Cloud Connectivity"),
		// 				AppearedAtSourceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-14T11:30:43.656Z"); return t}()),
		// 				AppearedAtTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-14T11:30:43.656Z"); return t}()),
		// 				ClearedAtSourceTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				ClearedAtTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DetailedInformation: map[string]*string{
		// 					"httP code": to.Ptr("400"),
		// 					"rpC code": to.Ptr("109"),
		// 					"status code": to.Ptr("12"),
		// 				},
		// 				ErrorDetails: &armstorsimple8000series.AlertErrorDetails{
		// 					Occurences: to.Ptr[int32](1),
		// 				},
		// 				Recommendation: to.Ptr("Check your network settings."),
		// 				Scope: to.Ptr(armstorsimple8000series.AlertScopeDevice),
		// 				Severity: to.Ptr(armstorsimple8000series.AlertSeverityWarning),
		// 				Source: &armstorsimple8000series.AlertSource{
		// 					Name: to.Ptr("Device05ForSDKTest"),
		// 					AlertSourceType: to.Ptr(armstorsimple8000series.AlertSourceTypeDevice),
		// 					TimeZone: to.Ptr("Pacific Standard Time"),
		// 				},
		// 				Status: to.Ptr(armstorsimple8000series.AlertStatusActive),
		// 				Title: to.Ptr("Connectivity to Cloud1 cannot be established"),
		// 			},
		// 	}},
		// }
	}
}
