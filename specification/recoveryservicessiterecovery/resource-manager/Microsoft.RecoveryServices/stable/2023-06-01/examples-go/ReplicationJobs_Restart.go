package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationJobs_Restart.json
func ExampleReplicationJobsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationJobsClient().BeginRestart(ctx, "vault1", "resourceGroupPS1", "0664564c-353e-401a-ab0c-722257c10e25", nil)
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
	// 	Name: to.Ptr("42c7d13b-790c-4609-8e0b-0936f1c5e5fb"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/42c7d13b-790c-4609-8e0b-0936f1c5e5fb"),
	// 	Properties: &armrecoveryservicessiterecovery.JobProperties{
	// 		ActivityID: to.Ptr("2443a5b4-e675-499f-8983-4126ea0e232c ActivityId: 2a776896-5e56-470b-af55-3c981283c4bc"),
	// 		AllowedActions: []*string{
	// 		},
	// 		CustomDetails: &armrecoveryservicessiterecovery.AsrJobDetails{
	// 			AffectedObjectDetails: map[string]*string{
	// 			},
	// 			InstanceType: to.Ptr("AsrJobDetails"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-03T10:45:14Z"); return t}()),
	// 		Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 		},
	// 		FriendlyName: to.Ptr("Restart job"),
	// 		ScenarioName: to.Ptr("RestartJob"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-03T10:45:12.1320757Z"); return t}()),
	// 		State: to.Ptr("Succeeded"),
	// 		StateDescription: to.Ptr("Completed"),
	// 		TargetInstanceType: to.Ptr("ProtectionEntity"),
	// 		TargetObjectID: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 		TargetObjectName: to.Ptr("vm1"),
	// 		Tasks: []*armrecoveryservicessiterecovery.ASRTask{
	// 			{
	// 				Name: to.Ptr("RemediateTask"),
	// 				AllowedActions: []*string{
	// 				},
	// 				CustomDetails: &armrecoveryservicessiterecovery.ManualActionTaskDetails{
	// 					InstanceType: to.Ptr("ManualActionTaskDetails"),
	// 				},
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-03T10:45:13.6458467Z"); return t}()),
	// 				Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 				},
	// 				FriendlyName: to.Ptr("Restarting job"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-03T10:45:13.5677237Z"); return t}()),
	// 				State: to.Ptr("Succeeded"),
	// 				StateDescription: to.Ptr("Completed"),
	// 				TaskID: to.Ptr("RemediateWfTask"),
	// 				TaskType: to.Ptr("TaskDetails"),
	// 		}},
	// 	},
	// }
}
