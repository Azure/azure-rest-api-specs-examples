package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-08-01-preview/examples/DistributedAvailabilityGroupsListByInstance.json
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
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag"),
		// 			Properties: &armsql.DistributedAvailabilityGroupProperties{
		// 				Databases: []*armsql.DistributedAvailabilityGroupDatabase{
		// 					{
		// 						ConnectedState: to.Ptr(armsql.ReplicaConnectedStateCONNECTED),
		// 						DatabaseName: to.Ptr("testdb"),
		// 						InstanceRedoReplicationLagSeconds: to.Ptr[int32](1),
		// 						InstanceReplicaID: to.Ptr("4713ed91-1e8c-497d-9bd4-d8a3935ae49a"),
		// 						InstanceSendReplicationLagSeconds: to.Ptr[int32](1),
		// 						LastBackupLsn: to.Ptr("71000009405700001"),
		// 						LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						LastCommitLsn: to.Ptr("71000009407900004"),
		// 						LastCommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						LastHardenedLsn: to.Ptr("71000009408100001"),
		// 						LastHardenedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						LastReceivedLsn: to.Ptr("71000009407900001"),
		// 						LastReceivedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						PartnerReplicaID: to.Ptr("8ffa9723-a1ec-4323-b929-c4aedee3894b"),
		// 						ReplicaState: to.Ptr("Catchup"),
		// 						SynchronizationHealth: to.Ptr(armsql.ReplicaSynchronizationHealthHEALTHY),
		// 				}},
		// 				DistributedAvailabilityGroupID: to.Ptr("c856cff5-a6fe-418e-8894-17799cc20f5d"),
		// 				DistributedAvailabilityGroupName: to.Ptr("dag"),
		// 				FailoverMode: to.Ptr(armsql.FailoverModeTypeNone),
		// 				InstanceAvailabilityGroupName: to.Ptr("testcl"),
		// 				InstanceLinkRole: to.Ptr(armsql.LinkRolePrimary),
		// 				PartnerAvailabilityGroupName: to.Ptr("BoxLocalAg1"),
		// 				PartnerEndpoint: to.Ptr("TCP://SERVER:7022"),
		// 				PartnerLinkRole: to.Ptr(armsql.LinkRoleSecondary),
		// 				ReplicationMode: to.Ptr(armsql.ReplicationModeTypeAsync),
		// 				SeedingMode: to.Ptr(armsql.SeedingModeTypeAutomatic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dag2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/distributedAvailabilityGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag2"),
		// 			Properties: &armsql.DistributedAvailabilityGroupProperties{
		// 				Databases: []*armsql.DistributedAvailabilityGroupDatabase{
		// 					{
		// 						ConnectedState: to.Ptr(armsql.ReplicaConnectedStateCONNECTED),
		// 						DatabaseName: to.Ptr("testdb2"),
		// 						InstanceReplicaID: to.Ptr("81608df4-0840-4219-b3ae-9cc46be9dae9"),
		// 						LastBackupLsn: to.Ptr("71000009405700001"),
		// 						LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						LastCommitLsn: to.Ptr("71000009407900004"),
		// 						LastCommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						LastHardenedLsn: to.Ptr("71000009408100001"),
		// 						LastHardenedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						LastReceivedLsn: to.Ptr("71000009407900001"),
		// 						LastReceivedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-01T00:00:00.000Z"); return t}()),
		// 						PartnerReplicaID: to.Ptr("57bc3421-e77a-4fdc-8a62-af105c4b1e38"),
		// 						ReplicaState: to.Ptr("Seeding"),
		// 						SeedingProgress: to.Ptr("80%"),
		// 						SynchronizationHealth: to.Ptr(armsql.ReplicaSynchronizationHealthHEALTHY),
		// 				}},
		// 				DistributedAvailabilityGroupID: to.Ptr("8a52d869-c17e-4546-ab03-e038569e672a"),
		// 				DistributedAvailabilityGroupName: to.Ptr("dag2"),
		// 				InstanceAvailabilityGroupName: to.Ptr("testcl2"),
		// 				InstanceLinkRole: to.Ptr(armsql.LinkRoleSecondary),
		// 				PartnerAvailabilityGroupName: to.Ptr("BoxLocalAg2"),
		// 				PartnerEndpoint: to.Ptr("TCP://SERVER:7022"),
		// 				PartnerLinkRole: to.Ptr(armsql.LinkRolePrimary),
		// 				ReplicationMode: to.Ptr(armsql.ReplicationModeTypeAsync),
		// 			},
		// 	}},
		// }
	}
}
