package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/GoldenGateConnections_ListBySubscription_MaximumSet_Gen.json
func ExampleGoldenGateConnectionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoldenGateConnectionsClient().NewListBySubscriptionPager(nil)
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
		// page = armoracledatabase.GoldenGateConnectionsClientListBySubscriptionResponse{
		// 	GoldenGateConnectionListResult: armoracledatabase.GoldenGateConnectionListResult{
		// 		Value: []*armoracledatabase.GoldenGateConnection{
		// 			{
		// 				Properties: &armoracledatabase.ConnectionBaseProperties{
		// 					ConnectionType: to.Ptr(armoracledatabase.ConnectionType("ConnectionBaseProperties")),
		// 					DisplayName: to.Ptr("resource1"),
		// 					ResourceAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
		// 					NetworkAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/networkAnchors/networkanchor1"),
		// 					CompartmentID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					RoutingMethod: to.Ptr(armoracledatabase.RoutingMethodSharedServiceEndpoint),
		// 					VaultID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					KeyID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					DoesUseSecretIDs: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					LifecycleState: to.Ptr(armoracledatabase.ConnectionLifecycleStateCreating),
		// 					LifecycleDetails: to.Ptr("ybufn"),
		// 					TimeCreated: to.Ptr("oqq"),
		// 					TimeUpdated: to.Ptr("2026-06-01T00:00:00Z"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("example"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key7255": to.Ptr("example"),
		// 				},
		// 				Location: to.Ptr("rtnrfzgdo"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
		// 				Name: to.Ptr("resource1"),
		// 				Type: to.Ptr("pl"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ns"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("example"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
