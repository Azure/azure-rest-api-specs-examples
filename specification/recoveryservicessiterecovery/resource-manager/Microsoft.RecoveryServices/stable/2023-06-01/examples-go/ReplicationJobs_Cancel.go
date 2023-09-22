package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationJobs_Cancel.json
func ExampleReplicationJobsClient_BeginCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationJobsClient().BeginCancel(ctx, "vault1", "resourceGroupPS1", "2653c648-fc72-4316-86f3-fdf8eaa0066b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armrecoveryservicessiterecovery.Job{
	// 	Name: to.Ptr("2653c648-fc72-4316-86f3-fdf8eaa0066b"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/2653c648-fc72-4316-86f3-fdf8eaa0066b"),
	// 	Properties: &armrecoveryservicessiterecovery.JobProperties{
	// 		ActivityID: to.Ptr("bfbbf6dd-9cbb-4cbc-98a6-faecc8891579 ActivityId: 07cc35ca-b63f-4e42-83c9-81ae0191c322"),
	// 		AllowedActions: []*string{
	// 		},
	// 		CustomDetails: &armrecoveryservicessiterecovery.AsrJobDetails{
	// 			AffectedObjectDetails: map[string]*string{
	// 				"PrimaryCloudId": to.Ptr("cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 				"PrimaryCloudName": to.Ptr("cloud1"),
	// 				"PrimaryFabricProviderId": to.Ptr("HyperVSite"),
	// 				"PrimaryVmId": to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 				"PrimaryVmName": to.Ptr("vm1"),
	// 				"PrimaryVmmId": to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 				"PrimaryVmmName": to.Ptr("cloud1"),
	// 				"ProtectionProfileId": to.Ptr("af095a1e-1f1b-5365-87c9-99162ebcfaf0"),
	// 				"RecoveryCloudId": to.Ptr("d38048d4-b460-4791-8ece-108395ee8478"),
	// 				"RecoveryCloudName": to.Ptr("Microsoft Azure"),
	// 				"RecoveryFabricProviderId": to.Ptr("Azure"),
	// 				"RecoveryVmId": to.Ptr(""),
	// 				"RecoveryVmName": to.Ptr("vm1"),
	// 				"RecoveryVmmId": to.Ptr("21a9403c-6ec1-44f2-b744-b4e50b792387"),
	// 				"RecoveryVmmName": to.Ptr("Microsoft Azure"),
	// 			},
	// 			InstanceType: to.Ptr("AsrJobDetails"),
	// 		},
	// 		Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 		},
	// 		FriendlyName: to.Ptr("Planned failover"),
	// 		ScenarioName: to.Ptr("PlannedFailover"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-03T06:16:06.3235186Z"); return t}()),
	// 		State: to.Ptr("Cancelling"),
	// 		StateDescription: to.Ptr("Cancelling"),
	// 		TargetInstanceType: to.Ptr("ProtectionEntity"),
	// 		TargetObjectID: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 		TargetObjectName: to.Ptr("vm1"),
	// 		Tasks: []*armrecoveryservicessiterecovery.ASRTask{
	// 		},
	// 	},
	// }
}
