package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/Connectors_ListBySubscription.json
func ExampleConnectorsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("74f5e23f-d4d9-410a-bb4d-8f0608adb10d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorsClient().NewListBySubscriptionPager(nil)
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
		// page = armimpactreporting.ConnectorsClientListBySubscriptionResponse{
		// 	ConnectorListResult: armimpactreporting.ConnectorListResult{
		// 		Value: []*armimpactreporting.Connector{
		// 			{
		// 				Properties: &armimpactreporting.ConnectorProperties{
		// 					ProvisioningState: to.Ptr(armimpactreporting.ProvisioningStateSucceeded),
		// 					ConnectorID: to.Ptr("430a444e-6a84-4a6f-8c50-124843ca7cd4"),
		// 					TenantID: to.Ptr("23a8d1da-a7e9-4443-9797-4cd3e3aeb8f8"),
		// 					ConnectorType: to.Ptr(armimpactreporting.PlatformAzureMonitor),
		// 					LastRunTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-19T06:23:56.238Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/74f5e23f-d4d9-410a-bb4d-8f0608adb10d/providers/Microsoft.Impact/connectors/testconnector1"),
		// 				Name: to.Ptr("testconnector1"),
		// 				Type: to.Ptr("microsoft.impact/connectors"),
		// 				SystemData: &armimpactreporting.SystemData{
		// 					CreatedBy: to.Ptr("testuser@hotmail.com"),
		// 					CreatedByType: to.Ptr(armimpactreporting.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-07T06:19:01.6431721Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("testuser@hotmail.com"),
		// 					LastModifiedByType: to.Ptr(armimpactreporting.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-15T08:29:20.8549373Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
