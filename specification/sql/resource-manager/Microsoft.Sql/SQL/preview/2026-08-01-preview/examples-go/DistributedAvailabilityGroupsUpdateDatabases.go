package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/DistributedAvailabilityGroupsUpdateDatabases.json
func ExampleDistributedAvailabilityGroupsClient_BeginUpdate_updateTheDatabasesOfADistributedAvailabilityGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDistributedAvailabilityGroupsClient().BeginUpdate(ctx, "testrg", "testcl", "dag", armsql.DistributedAvailabilityGroup{
		Properties: &armsql.DistributedAvailabilityGroupProperties{
			Databases: []*armsql.DistributedAvailabilityGroupDatabase{
				{
					DatabaseName: to.Ptr("testdb1"),
				},
				{
					DatabaseName: to.Ptr("testdb2"),
				},
				{
					DatabaseName: to.Ptr("testdb3"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.DistributedAvailabilityGroupsClientUpdateResponse{
	// 	DistributedAvailabilityGroup: armsql.DistributedAvailabilityGroup{
	// 		Properties: &armsql.DistributedAvailabilityGroupProperties{
	// 			DistributedAvailabilityGroupName: to.Ptr("dag"),
	// 			DistributedAvailabilityGroupID: to.Ptr("c856cff5-a6fe-418e-8894-17799cc20f5d"),
	// 			ReplicationMode: to.Ptr(armsql.ReplicationModeTypeAsync),
	// 			PartnerLinkRole: to.Ptr(armsql.LinkRoleSecondary),
	// 			PartnerAvailabilityGroupName: to.Ptr("BoxLocalAg1"),
	// 			PartnerEndpoint: to.Ptr("TCP://SERVER:7022"),
	// 			InstanceLinkRole: to.Ptr(armsql.LinkRolePrimary),
	// 			InstanceAvailabilityGroupName: to.Ptr("testcl"),
	// 			FailoverMode: to.Ptr(armsql.FailoverModeTypeNone),
	// 			SeedingMode: to.Ptr(armsql.SeedingModeTypeAutomatic),
	// 			LinkMode: to.Ptr(armsql.LinkModeTypeMultiDatabase),
	// 			Databases: []*armsql.DistributedAvailabilityGroupDatabase{
	// 				{
	// 					DatabaseName: to.Ptr("testdb1"),
	// 					InstanceReplicaID: to.Ptr("4713ed91-1e8c-497d-9bd4-d8a3935ae49a"),
	// 					PartnerReplicaID: to.Ptr("8ffa9723-a1ec-4323-b929-c4aedee3894b"),
	// 					ReplicaState: to.Ptr("Connected"),
	// 					SynchronizationHealth: to.Ptr(armsql.ReplicaSynchronizationHealthHEALTHY),
	// 					ConnectedState: to.Ptr(armsql.ReplicaConnectedStateCONNECTED),
	// 					LastReceivedLsn: to.Ptr("71000009407900001"),
	// 					LastReceivedTime: to.Ptr(time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC)),
	// 					LastCommitLsn: to.Ptr("71000009407900004"),
	// 					LastCommitTime: to.Ptr(time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC)),
	// 					LastHardenedLsn: to.Ptr("71000009408100001"),
	// 					LastHardenedTime: to.Ptr(time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC)),
	// 					LastBackupLsn: to.Ptr("71000009405700001"),
	// 					LastBackupTime: to.Ptr(time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC)),
	// 					InstanceSendReplicationLagSeconds: to.Ptr[int32](1),
	// 					InstanceRedoReplicationLagSeconds: to.Ptr[int32](1),
	// 				},
	// 				{
	// 					DatabaseName: to.Ptr("testdb2"),
	// 					InstanceReplicaID: to.Ptr("5824fe02-2f9d-508e-0ce5-e9b4e4af5905"),
	// 					PartnerReplicaID: to.Ptr("9006a834-b2fd-5434-ca3a-d5bfeff5985c"),
	// 					ReplicaState: to.Ptr("Seeding"),
	// 					SeedingProgress: to.Ptr("30%"),
	// 					SynchronizationHealth: to.Ptr(armsql.ReplicaSynchronizationHealthPARTIALLYHEALTHY),
	// 					ConnectedState: to.Ptr(armsql.ReplicaConnectedStateCONNECTED),
	// 					LastReceivedLsn: to.Ptr("72000010508000002"),
	// 					LastReceivedTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					LastCommitLsn: to.Ptr("72000010508000005"),
	// 					LastCommitTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					LastHardenedLsn: to.Ptr("72000010509200002"),
	// 					LastHardenedTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					LastBackupLsn: to.Ptr("72000010506800002"),
	// 					LastBackupTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					InstanceSendReplicationLagSeconds: to.Ptr[int32](0),
	// 					InstanceRedoReplicationLagSeconds: to.Ptr[int32](0),
	// 				},
	// 				{
	// 					DatabaseName: to.Ptr("testdb3"),
	// 					InstanceReplicaID: to.Ptr("9477cc47-de5f-43a9-8d08-4b8547b00f5f"),
	// 					PartnerReplicaID: to.Ptr("f2b6db36-ab17-4206-9bc1-914d1e3504c4"),
	// 					ReplicaState: to.Ptr("Seeding"),
	// 					SeedingProgress: to.Ptr("10%"),
	// 					SynchronizationHealth: to.Ptr(armsql.ReplicaSynchronizationHealthPARTIALLYHEALTHY),
	// 					ConnectedState: to.Ptr(armsql.ReplicaConnectedStateCONNECTED),
	// 					LastReceivedLsn: to.Ptr("72000010508000002"),
	// 					LastReceivedTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					LastCommitLsn: to.Ptr("72000010508000005"),
	// 					LastCommitTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					LastHardenedLsn: to.Ptr("72000010509200002"),
	// 					LastHardenedTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					LastBackupLsn: to.Ptr("72000010506800002"),
	// 					LastBackupTime: to.Ptr(time.Date(2026, time.August, 1, 0, 5, 0, 0, time.UTC)),
	// 					InstanceSendReplicationLagSeconds: to.Ptr[int32](0),
	// 					InstanceRedoReplicationLagSeconds: to.Ptr[int32](0),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag"),
	// 		Name: to.Ptr("dag"),
	// 		Type: to.Ptr("Microsoft.Sql/managedInstances/distributedAvailabilityGroups"),
	// 	},
	// }
}
