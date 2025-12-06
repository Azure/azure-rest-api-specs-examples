package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/RansomwareReports_Get.json
func ExampleRansomwareReportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRansomwareReportsClient().Get(ctx, "myRG", "account1", "pool1", "volume1", "ransomwareReport1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.RansomwareReportsClientGetResponse{
	// 	RansomwareReport: &armnetapp.RansomwareReport{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/ransomwarereports/ransomwareReport1"),
	// 		Name: to.Ptr("account1/pool1/volume1/ransomwareReport1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/ransomwarereports"),
	// 		Properties: &armnetapp.RansomwareReportProperties{
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-15T13:23:33Z"); return t}()),
	// 			Severity: to.Ptr(armnetapp.RansomwareReportSeverityModerate),
	// 			State: to.Ptr(armnetapp.RansomwareReportStateActive),
	// 			ClearedCount: to.Ptr[int32](0),
	// 			ReportedCount: to.Ptr[int32](1),
	// 			Suspects: []*armnetapp.RansomwareSuspects{
	// 				{
	// 					Extension: to.Ptr(".threat"),
	// 					Resolution: to.Ptr(armnetapp.RansomwareSuspectResolutionPotentialThreat),
	// 					FileCount: to.Ptr[int32](1),
	// 					SuspectFiles: []*armnetapp.SuspectFile{
	// 						{
	// 							SuspectFileName: to.Ptr("file1.threat"),
	// 							FileTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-15T13:23:33Z"); return t}()),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
