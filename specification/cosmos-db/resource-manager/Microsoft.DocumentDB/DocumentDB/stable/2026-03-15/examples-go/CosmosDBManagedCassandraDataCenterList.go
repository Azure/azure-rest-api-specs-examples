package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBManagedCassandraDataCenterList.json
func ExampleCassandraDataCentersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraDataCentersClient().NewListPager("cassandra-prod-rg", "cassandra-prod", nil)
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
		// page = armcosmos.CassandraDataCentersClientListResponse{
		// 	ListDataCenters: armcosmos.ListDataCenters{
		// 		Value: []*armcosmos.DataCenterResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/dataCenters/dc1"),
		// 				Name: to.Ptr("dc1"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/dataCenters"),
		// 				Properties: &armcosmos.DataCenterResourceProperties{
		// 					ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
		// 					DataCenterLocation: to.Ptr("West US 2"),
		// 					DelegatedSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1"),
		// 					NodeCount: to.Ptr[int32](9),
		// 					SeedNodes: []*armcosmos.SeedNode{
		// 						{
		// 							IPAddress: to.Ptr("192.168.12.2"),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("192.168.12.3"),
		// 						},
		// 						{
		// 							IPAddress: to.Ptr("192.168.12.4"),
		// 						},
		// 					},
		// 					Base64EncodedCassandraYamlFragment: to.Ptr("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
