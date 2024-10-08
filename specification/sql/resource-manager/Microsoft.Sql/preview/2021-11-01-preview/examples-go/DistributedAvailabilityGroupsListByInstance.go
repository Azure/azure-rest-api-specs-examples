package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DistributedAvailabilityGroupsListByInstance.json
func ExampleDistributedAvailabilityGroupsClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDistributedAvailabilityGroupsClient().NewListByInstancePager("testrg", "testcl", nil)
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
		// page.DistributedAvailabilityGroupsListResult = armsql.DistributedAvailabilityGroupsListResult{
		// 	Value: []*armsql.DistributedAvailabilityGroup{
		// 		{
		// 			Name: to.Ptr("dag"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/distributedAvailabilityGroups"),
		// 			ID: to.Ptr("/subscriptions/f2669dff-5f08-45dd-b857-b2a60b72cdc9/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag"),
		// 			Properties: &armsql.DistributedAvailabilityGroupProperties{
		// 				DistributedAvailabilityGroupID: to.Ptr("6bc05a51-aa36-a196-09bd-481d7a0973c0"),
		// 				LastHardenedLsn: to.Ptr("39000000030400001"),
		// 				LinkState: to.Ptr("Catchup"),
		// 				ReplicationMode: to.Ptr(armsql.ReplicationModeAsync),
		// 				SourceEndpoint: to.Ptr("TCP://SERVER:7022"),
		// 				SourceReplicaID: to.Ptr("543dd519-7585-faff-6ad2-11fb826d4f4d"),
		// 				TargetDatabase: to.Ptr("testdb"),
		// 				TargetReplicaID: to.Ptr("7e218aba-0a53-6231-be09-895d99f96bf2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dag2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/distributedAvailabilityGroups"),
		// 			ID: to.Ptr("/subscriptions/f2669dff-5f08-45dd-b857-b2a60b72cdc9/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag2"),
		// 			Properties: &armsql.DistributedAvailabilityGroupProperties{
		// 				DistributedAvailabilityGroupID: to.Ptr("7ec05a51-aa36-a196-09bd-481d7a0973c0"),
		// 				LastHardenedLsn: to.Ptr("39000000030400001"),
		// 				LinkState: to.Ptr("Catchup"),
		// 				ReplicationMode: to.Ptr(armsql.ReplicationModeAsync),
		// 				SourceEndpoint: to.Ptr("TCP://SERVER:7022"),
		// 				SourceReplicaID: to.Ptr("d423d519-7585-faff-6ad2-11fb826d4f4d"),
		// 				TargetDatabase: to.Ptr("testdb2"),
		// 				TargetReplicaID: to.Ptr("32578aba-0a53-6231-be09-895d99f96bf2"),
		// 			},
		// 	}},
		// }
	}
}
